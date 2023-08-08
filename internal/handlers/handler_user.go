package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/zvash/go-jwt-auth/internal/config"
	"github.com/zvash/go-jwt-auth/internal/database"
	"github.com/zvash/go-jwt-auth/internal/resources"
	"github.com/zvash/go-jwt-auth/internal/responsemaker"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

var app *config.AppConfig

func SetAppConfig(a *config.AppConfig) {
	app = a
}

func Register(w http.ResponseWriter, r *http.Request) {
	type userPayload struct {
		Name                 string `json:"name"`
		Email                string `json:"email"`
		Password             string `json:"password"`
		PasswordConfirmation string `json:"password_confirmation"`
	}

	payload := userPayload{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		responsemaker.ResponseWithError(w, 400, fmt.Sprintf("Error parsing responsemaker: %v", err))
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		responsemaker.ResponseWithError(w, 500, "could not hash the password")
		return
	}
	dbUser, err := app.DB.CreateUser(r.Context(), database.CreateUserParams{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		responsemaker.ResponseWithError(w, 500, "could not create the new user")
		return
	}
	user := resources.User{}
	user.FillWithDbUser(dbUser)
	responsemaker.RespondWithJSON(w, 201, user)
}
