package controllers

import (
	"net/http"

	"github.com/schattenbrot/auth/internal/models"
	"github.com/schattenbrot/auth/internal/utils"
)

func (m *Repository) GetMe(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(m.App.UserContextKey).(*models.User)

	utils.Send(w, m.App.Logger, http.StatusOK, user)
}
