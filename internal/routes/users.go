package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/schattenbrot/auth/internal/controllers"
	"github.com/schattenbrot/auth/internal/middlewares"
)

func userRoutes(router chi.Router) {
	router.With(middlewares.Repo.IsAuth, middlewares.Repo.IsAdmin).Get("/", controllers.Repo.GetUsers)

	router.With(middlewares.Repo.IsAuth).Route("/me", func(router chi.Router) {
		router.Get("/", controllers.Repo.GetMe)
		router.With(middlewares.Repo.ValidateUpdateMeUsername).Patch("/username", controllers.Repo.UpdateMeUsername)
		router.With(middlewares.Repo.ValidateUpdateMeEmail).Patch("/email", controllers.Repo.UpdateMeEmail)
		router.With(middlewares.Repo.ValidateUpdateMePassword).Patch("/password", controllers.Repo.UpdateMePassword)
	})

	router.With(middlewares.Repo.IsAuth, middlewares.Repo.RequiredId).Route("/{id}", func(router chi.Router) {
		router.Get("/", controllers.Repo.GetUser)
		router.With(middlewares.Repo.ValidateUpdateUser, middlewares.Repo.IsAdmin, middlewares.Repo.RolesExist).Put("/", controllers.Repo.UpdateUser)
		router.With(middlewares.Repo.IsAdmin).Delete("/", controllers.Repo.DeleteUser)
		router.With(middlewares.Repo.IsAdmin).Get("/reactivate", controllers.Repo.ReactivateUser)
	})
}
