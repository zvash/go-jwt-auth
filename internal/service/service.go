package service

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/zvash/go-jwt-auth/internal/database"
)

type App struct {
	DB     *database.Queries
	Server *fiber.App
}

func (a *App) SetDB(connection *sql.DB) {
	a.DB = database.New(connection)
}

func (a *App) SetServer(server *fiber.App) {
	a.Server = server
}
