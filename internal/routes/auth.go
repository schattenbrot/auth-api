package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/schattenbrot/auth/internal/controllers"
	"github.com/schattenbrot/auth/internal/middlewares"
)

func authRoutes(router chi.Router) {
	router.Get("/api-status", controllers.Repo.ApiStatus)
	router.With(middlewares.Repo.ValidateAuthUser, middlewares.Repo.EmailExists).Post("/sign-up", controllers.Repo.SignUp)
	router.Post("/sign-in", controllers.Repo.SignIn)
	router.Get("/sign-out", controllers.Repo.SignOut)
}
