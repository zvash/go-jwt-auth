package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zvash/go-jwt-auth/internal/service"
)

func SetupRoutes(app *service.App) {
	app.Server.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
}
