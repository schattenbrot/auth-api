package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/schattenbrot/auth/internal/models"
	"github.com/schattenbrot/auth/internal/utils"
)

func (m *Repository) UpdateMeUsername(w http.ResponseWriter, r *http.Request) {
	currentUser := r.Context().Value(m.App.UserContextKey).(*models.User)

	var user models.UpdateMeUsernameUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.SendError(w, m.App.Logger, err)
		return
	}

	updatedUser, err := m.DB.UpdateUserById(currentUser.ID.Hex(), &models.User{
		Username: user.Username,
	})
	if err != nil {
		utils.SendError(w, m.App.Logger, err, http.StatusInternalServerError)
		return
	}

	utils.Send(w, m.App.Logger, http.StatusOK, updatedUser)
}
