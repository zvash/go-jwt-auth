package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/zvash/go-jwt-auth/internal/config"
	"github.com/zvash/go-jwt-auth/internal/handlers"
	"github.com/zvash/go-jwt-auth/internal/middlewares"
	"github.com/zvash/go-jwt-auth/internal/router"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	err := setupAndListen()
	if err != nil {
		log.Panic(err)
	}
}

func setupAndListen() error {
	app := config.AppConfig{}
	err := godotenv.Load()
	if err != nil {
		return err
	}
	portString := os.Getenv("PORT")
	if portString == "" {
		portString = "8000"
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return err
	}
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		return err
	}
	app.SetDB(conn)

	middlewares.SetAppConfig(&app)
	router.SetAppConfig(&app)
	handlers.SetAppConfig(&app)

	srv := &http.Server{
		Handler: router.SetupRoutes(),
		Addr:    fmt.Sprintf(":%s", portString),
	}
	log.Printf("Server starting on port: %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
