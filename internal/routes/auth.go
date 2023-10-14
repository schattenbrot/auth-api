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
	router.Get("/activate-email", controllers.Repo.ActivateEmail)
	router.With(middlewares.Repo.ValidateResetPasswordRequest).Post("/reset-password/request", controllers.Repo.ResetPasswordRequest)
	router.Post("/reset-password", controllers.Repo.ResetPasswordByToken)
	router.With(middlewares.Repo.IsAuth, middlewares.Repo.IsAdmin, middlewares.Repo.ValidateResetPasswordRequest).Post("/reset-password/revoke", controllers.Repo.ResetPasswordRevoke)
}
