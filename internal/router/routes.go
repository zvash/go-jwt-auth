package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/zvash/go-jwt-auth/internal/config"
	"net/http"
)

var app *config.AppConfig

func SetAppConfig(a *config.AppConfig) {
	app = a
}

func SetupRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	v1Router := chi.NewRouter()

	v1Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/responsemaker")
		w.WriteHeader(200)
	})

	v1Router.Mount("/users", setupUserRoutes())

	router.Mount("/v1", v1Router)
	authRouter, _ := setupAuthRoutes(app)
	router.Mount("/oauth", authRouter)
	return router
}
