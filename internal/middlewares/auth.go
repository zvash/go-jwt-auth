package middlewares

import (
	"fmt"
	"github.com/zvash/go-jwt-auth/internal/config"
	"github.com/zvash/go-jwt-auth/internal/database"
	"github.com/zvash/go-jwt-auth/internal/responsemaker"
	"net/http"
	"strconv"
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
		info, err := app.AuthServer.ValidationBearerToken(r)
		if err != nil {
			responsemaker.ResponseWithError(w, 403, fmt.Sprintf("Unathorized"))
			return
		}
		userId, err := strconv.Atoi(info.GetUserID())
		if err != nil {
			fmt.Printf("Invalid user id: %v", err)
			responsemaker.ResponseWithError(w, 403, fmt.Sprintf("Invalid UserId"))
			return
		}
		dbUser, err := app.DB.GetUserById(r.Context(), int32(userId))
		if err != nil {
			responsemaker.ResponseWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}
		handler(w, r, dbUser)
	}
}
