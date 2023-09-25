package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/schattenbrot/auth/internal/models"
	"github.com/schattenbrot/auth/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *Repository) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")

	_, err := m.DB.UpdateUserById(userId, &models.User{
		Inactive: utils.BoolPointer(true),
	})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			utils.SendError(w, m.App.Logger, err, http.StatusNotFound)
			return
		}
		utils.SendError(w, m.App.Logger, err, http.StatusInternalServerError)
		return
	}

	type jsonResponse struct {
		Message string `json:"message"`
	}
	response := &jsonResponse{
		Message: "user deleted",
	}

	utils.Send(w, m.App.Logger, http.StatusOK, response)
}
