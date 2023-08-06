package main

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/zvash/go-jwt-auth/internal/service"
	"github.com/zvash/go-jwt-auth/routes"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	_, err := setupAndListen()
	if err != nil {
		log.Panic(err)
	}
}

func setupAndListen() (service.App, error) {
	app := service.App{}
	err := godotenv.Load()
	if err != nil {
		return app, err
	}
	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "8000"
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return app, err
	}
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		return app, err
	}
	app.SetDB(conn)
	app.SetServer(fiber.New())
	routes.SetupRoutes(&app)

	err = app.Server.Listen(fmt.Sprintf(":%s", portString))
	if err != nil {
		return app, err
	}

	return app, nil
}
