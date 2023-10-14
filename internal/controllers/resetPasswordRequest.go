package controllers

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/schattenbrot/auth/internal/models"
	"github.com/schattenbrot/auth/internal/utils"
)

func (m *Repository) ResetPasswordRequest(w http.ResponseWriter, r *http.Request) {
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

	resetTokenBuffer := make([]byte, 128)
	_, err = rand.Read(resetTokenBuffer)
	if err != nil {
		utils.SendError(w, m.App.Logger, err, http.StatusInternalServerError)
		return
	}
	resetToken := fmt.Sprintf("%x", resetTokenBuffer)

	updatedUser, err := m.DB.UpdateUserById(user.ID.Hex(), &models.User{
		ResetPasswordExpires: time.Now().Add(time.Hour * 12),
		ResetPasswordToken:   &resetToken,
	})
	if err != nil {
		utils.SendError(w, m.App.Logger, err, http.StatusInternalServerError)
		return
	}

	mail := &models.Mail{
		From:    m.App.Config.EmailProvider.Email,
		To:      []string{user.Email},
		Subject: "Password reset requested",
		Body:    "If you were the one issueing the password reset please use http://localhost:8080/auth/reset-password?token=" + *updatedUser.ResetPasswordToken,
	}
	err = mail.Send(m.App.Config.EmailProvider.Email, m.App.Config.EmailProvider.Password)
	if err != nil {
		utils.SendError(w, m.App.Logger, err, http.StatusInternalServerError)
		return
	}

	type resp struct {
		Message string `json:"message"`
	}

	utils.Send(w, m.App.Logger, http.StatusOK, &resp{
		Message: "The password reset got successfully requested",
	})
}
