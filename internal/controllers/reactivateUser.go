package controllers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/schattenbrot/auth/internal/models"
	"github.com/schattenbrot/auth/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *Repository) ReactivateUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")

	userToReactivate, err := m.DB.GetInactiveUserById(userId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			utils.SendError(w, m.App.Logger, err, http.StatusNotFound)
			return
		}
		utils.SendError(w, m.App.Logger, err, http.StatusInternalServerError)
		return
	}

	_, err = m.DB.GetUserByEmail(userToReactivate.Email)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			utils.SendError(w, m.App.Logger, err, http.StatusInternalServerError)
			return
		}
	}
	if err == nil {
		err = errors.New("user already exists")
		utils.SendError(w, m.App.Logger, err)
		return
	}

	_, err = m.DB.UpdateUserById(userId, &models.User{
		Inactive: utils.BoolPointer(false),
	})
	if err != nil {
		utils.SendError(w, m.App.Logger, err, http.StatusInternalServerError)
		return
	}

	type jsonResponse struct {
		Message string `json:"message"`
	}
	response := &jsonResponse{
		Message: "user reactivated",
	}

	utils.Send(w, m.App.Logger, http.StatusOK, response)
}
