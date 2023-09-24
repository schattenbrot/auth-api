package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/schattenbrot/auth/internal/models"
	"github.com/schattenbrot/auth/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func (m *Repository) UpdateMePassword(w http.ResponseWriter, r *http.Request) {
	currentUser := r.Context().Value(m.App.UserContextKey).(*models.User)

	var user models.UpdateMePasswordUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.SendError(w, m.App.Logger, err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(currentUser.Password), []byte(user.OldPassword))
	if err != nil {
		utils.SendError(w, m.App.Logger, err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		utils.SendError(w, m.App.Logger, err, http.StatusInternalServerError)
		return
	}

	updatedUser, err := m.DB.UpdateUserById(currentUser.ID.Hex(), &models.User{
		Password: string(hashedPassword),
	})
	if err != nil {
		utils.SendError(w, m.App.Logger, err, http.StatusInternalServerError)
		return
	}

	utils.Send(w, m.App.Logger, http.StatusOK, updatedUser)
}
