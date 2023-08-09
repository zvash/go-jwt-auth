package router

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/jackc/pgx/v4"
	pg "github.com/vgarvardt/go-oauth2-pg/v4"
	"github.com/vgarvardt/go-pg-adapter/pgx4adapter"
	"github.com/zvash/go-jwt-auth/internal/config"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type Oauth2ClientInfo struct {
	ID     string
	Secret string
	Domain string
	Public bool
	UserID string
}

//
//func (oc *Oauth2ClientInfo) GetID() string {
//	return oc.ID
//}
//
//func (oc *Oauth2ClientInfo) GetSecret() string {
//	return oc.Secret
//}
//
//func (oc *Oauth2ClientInfo) GetDomain() string {
//	return oc.Domain
//}
//
//func (oc *Oauth2ClientInfo) IsPublic() bool {
//	return oc.Public
//}
//
//func (oc *Oauth2ClientInfo) GetUserID() string {
//	return oc.UserID
//}

func setupAuthRoutes(app *config.AppConfig) (*chi.Mux, error) {

	authRouter := chi.NewRouter()
	// Middlewares
	authRouter.Use(middleware.Logger)
	authRouter.Use(middleware.Recoverer)

	//manager := manage.NewDefaultManager()
	//manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	//
	//// token memory store
	//manager.MustTokenStorage(store.NewMemoryTokenStore())
	//clientStore := store.NewClientStore()
	//
	//clients, err := app.DB.GetAllClients(context.Background())
	//if err != nil {
	//	return authRouter, err
	//}
	//for _, oauthClient := range clients {
	//	clientId := oauthClient.ClientID.String()
	//	err = clientStore.Set(clientId, &models.Client{
	//		ID:     clientId,
	//		Secret: oauthClient.Secret,
	//		Domain: oauthClient.Domain,
	//	})
	//	if err != nil {
	//		return authRouter, err
	//	}
	//}
	//manager.MapClientStorage(clientStore)
	//
	//srv := server.NewServer(server.NewConfig(), manager)

	srv := buildOauthServer()
	app.AuthServer = srv

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
		userID = strconv.Itoa(int(user.ID))
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

func buildOauthServer() *server.Server {
	pgxConn, _ := pgx.Connect(context.TODO(), os.Getenv("DB_URL"))
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	adapter := pgx4adapter.NewConn(pgxConn)
	tokenStore, _ := pg.NewTokenStore(adapter, pg.WithTokenStoreGCInterval(time.Minute))
	defer tokenStore.Close()
	clientStore, _ := pg.NewClientStore(adapter)

	//oauth2ClientInfo := Oauth2ClientInfo{
	//	ID:     "1",
	//	Secret: "H8paZSe6k249cZeSWgIPrMSo2F4P8H4N",
	//	Domain: "http://localhost",
	//	Public: true,
	//	UserID: "1",
	//}
	//
	//err := clientStore.Create(&oauth2ClientInfo)
	//if err != nil {
	//	fmt.Printf("Couldn't create client: %v", err)
	//	return nil
	//}

	manager.MapTokenStorage(tokenStore)
	manager.MapClientStorage(clientStore)

	srv := server.NewServer(server.NewConfig(), manager)
	return srv
}
