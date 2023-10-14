package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/schattenbrot/auth/internal/models"
	"github.com/schattenbrot/auth/internal/utils"
)

func (m *Repository) ResetPasswordRevoke(w http.ResponseWriter, r *http.Request) {
	var requestingUser models.ResetPasswordRequestUser
	err := json.NewDecoder(r.Body).Decode(&requestingUser)
	if err != nil {
		utils.SendError(w, m.App.Logger, err)
		return
	}

	user, err := m.DB.GetUserByEmail(requestingUser.Email)
	if err != nil {
		utils.SendError(w, m.App.Logger, err)
		return
	}

	resetToken := ""
	m.DB.UpdateUserById(user.ID.Hex(), &models.User{
		ResetPasswordToken:   &resetToken,
		ResetPasswordExpires: time.Now(),
	})

	type resp struct {
		Message string `json:"message"`
	}
	utils.Send(w, m.App.Logger, http.StatusOK, &resp{
		Message: "The password reset has been revoked",
	})
}
