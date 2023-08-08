package config

import (
	"database/sql"
	"github.com/zvash/go-jwt-auth/internal/database"
)

type AppConfig struct {
	DB *database.Queries
}

func (app *AppConfig) SetDB(connection *sql.DB) {
	app.DB = database.New(connection)
}
