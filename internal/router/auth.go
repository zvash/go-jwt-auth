package router

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/zvash/go-jwt-auth/internal/config"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func setupAuthRoutes(app *config.AppConfig) (*chi.Mux, error) {

	authRouter := chi.NewRouter()
	// Middlewares
	authRouter.Use(middleware.Logger)
	authRouter.Use(middleware.Recoverer)

	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())
	clientStore := store.NewClientStore()

	clients, err := app.DB.GetAllClients(context.Background())
	if err != nil {
		return authRouter, err
	}
	for _, oauthClient := range clients {
		clientId := oauthClient.ClientID.String()
		err = clientStore.Set(clientId, &models.Client{
			ID:     clientId,
			Secret: oauthClient.Secret,
			Domain: oauthClient.Domain,
		})
		if err != nil {
			return authRouter, err
		}
	}
	manager.MapClientStorage(clientStore)

	srv := server.NewServer(server.NewConfig(), manager)
	srv.SetClientInfoHandler(server.ClientFormHandler)
	srv.SetAllowedGrantType(oauth2.PasswordCredentials, oauth2.ClientCredentials)

	srv.SetPasswordAuthorizationHandler(func(ctx context.Context, clientID, username, password string) (userID string, err error) {
		user, err := app.DB.GetUserByUsername(ctx, username)
		if err != nil {
			fmt.Printf("couldn't get the user: %v", err)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			fmt.Printf("password mismatch: %v", err)
			return
		}
		userID = username
		return
	})

	authRouter.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Println("error ", err)
		}
		err = srv.HandleTokenRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	return authRouter, nil
}
