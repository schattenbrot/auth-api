package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/schattenbrot/auth/internal/models"
	"github.com/schattenbrot/auth/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *Repository) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.SendError(w, m.App.Logger, err)
		return
	}

	updatedUser, err := m.DB.UpdateUserById(userId, &models.User{
		Username: user.Username,
		Email:    user.Email,
		Roles:    user.Roles,
		Inactive: user.Inactive,
	})
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			err = errors.New("email already exists")
			utils.SendError(w, m.App.Logger, err)
			return
		}
		utils.SendError(w, m.App.Logger, err, http.StatusInternalServerError)
		return
	}

	utils.Send(w, m.App.Logger, http.StatusOK, updatedUser)
}
