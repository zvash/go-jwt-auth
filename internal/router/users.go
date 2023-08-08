package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/zvash/go-jwt-auth/internal/handlers"
)

func setupUserRoutes() *chi.Mux {
	userRouter := chi.NewRouter()
	// Middlewares
	userRouter.Use(middleware.Logger)
	userRouter.Use(middleware.Recoverer)

	userRouter.Post("/register", handlers.Register)

	return userRouter
}
