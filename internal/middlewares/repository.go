package middlewares

import (
	"github.com/schattenbrot/auth/internal/config"
	"github.com/schattenbrot/auth/internal/database"
)

// Repository represents the controller repository
type Repository struct {
	App *config.AppConfig
	DB  database.DatabaseRepo
}

// Repo is the repository used by the controllers
var Repo *Repository

func NewRepo(a *config.AppConfig, db database.DatabaseRepo) {
	Repo = &Repository{
		App: a,
		DB:  db,
	}
}
