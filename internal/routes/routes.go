package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/schattenbrot/auth/internal/config"
)

func Routes(app config.AppConfig) *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:     app.Config.Cors,
		AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposedHeaders:     []string{"Content-Length"},
		AllowCredentials:   true,
		MaxAge:             300,
		OptionsPassthrough: true,
		Debug:              false,
	}))

	router.Route("/auth", authRoutes)

	return router
}
