package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/schattenbrot/auth/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *Repository) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	user, err := m.DB.GetUserById(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			utils.SendError(w, m.App.Logger, err, http.StatusNotFound)
			return
		}
		utils.SendError(w, m.App.Logger, err, http.StatusInternalServerError)
		return
	}

	utils.Send(w, m.App.Logger, http.StatusOK, user)
}
