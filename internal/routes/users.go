package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/schattenbrot/auth/internal/controllers"
	"github.com/schattenbrot/auth/internal/middlewares"
)

func userRoutes(router chi.Router) {
	router.With(middlewares.Repo.IsAuth, middlewares.Repo.IsAdmin).Get("/", controllers.Repo.GetUsers)
	router.With(middlewares.Repo.IsAuth).Get("/me", controllers.Repo.GetMe)
	router.With(middlewares.Repo.IsAuth, middlewares.Repo.ValidateUpdateMeUsername).Patch("/me/username", controllers.Repo.UpdateMeUsername)
	router.With(middlewares.Repo.IsAuth, middlewares.Repo.ValidateUpdateMeEmail).Patch("/me/email", controllers.Repo.UpdateMeEmail)
}
