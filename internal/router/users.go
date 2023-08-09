package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/zvash/go-jwt-auth/internal/handlers"
	"github.com/zvash/go-jwt-auth/internal/middlewares"
)

func setupUserRoutes() *chi.Mux {
	userRouter := chi.NewRouter()
	// Middlewares
	userRouter.Use(middleware.Logger)
	userRouter.Use(middleware.Recoverer)

	userRouter.Post("/register", handlers.Register)
	userRouter.Get("/authenticated", middlewares.Auth(handlers.Authenticated))

	return userRouter
}
