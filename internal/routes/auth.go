package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/schattenbrot/auth/internal/controllers"
)

func authRoutes(router chi.Router) {
	router.Get("/api-status", controllers.Repo.ApiStatus)
	router.Post("/sign-up", controllers.Repo.SignUp)
	router.Post("/sign-in", controllers.Repo.SignIn)
	router.Get("/sign-out", controllers.Repo.SignOut)
}
