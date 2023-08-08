package middlewares

import (
	"fmt"
	"github.com/zvash/go-jwt-auth/internal/config"
	"github.com/zvash/go-jwt-auth/internal/database"
	"github.com/zvash/go-jwt-auth/internal/responsemaker"
	"net/http"
)

var app *config.AppConfig

func SetAppConfig(a *config.AppConfig) {
	app = a
}

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func Auth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//apiKey, err := auth.GetAPIKey(r.Header)
		//if err != nil {
		//	responseWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		//	return
		//}
		dbUser, err := app.DB.GetUserByUsername(r.Context(), "test")
		if err != nil {
			responsemaker.ResponseWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}
		handler(w, r, dbUser)
	}
}
