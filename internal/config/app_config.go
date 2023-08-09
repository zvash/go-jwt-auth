package config

import (
	"database/sql"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/zvash/go-jwt-auth/internal/database"
)

type AppConfig struct {
	DB         *database.Queries
	AuthServer *server.Server
}

func (app *AppConfig) SetDB(connection *sql.DB) {
	app.DB = database.New(connection)
}

func (app *AppConfig) SetAuthServer(server *server.Server) {
	app.AuthServer = server
}
