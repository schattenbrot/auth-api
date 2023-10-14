package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/schattenbrot/auth/internal/models"
	"github.com/schattenbrot/auth/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func (m *Repository) ResetPasswordByToken(w http.ResponseWriter, r *http.Request) {
	resetPasswordToken := r.URL.Query().Get("token")

	user, err := m.DB.GetUserByResetPasswordToken(resetPasswordToken)
	if err != nil {
		utils.SendError(w, m.App.Logger, err)
		return
	}

	if user.ResetPasswordExpires.Before(time.Now()) {
		err = errors.New("reset password token expired")
		utils.SendError(w, m.App.Logger, err)
		return
	}

	var requestingUser models.ResetPasswordByTokenUser
	err = json.NewDecoder(r.Body).Decode(&requestingUser)
	if err != nil {
		utils.SendError(w, m.App.Logger, err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestingUser.NewPassword))
	if err == nil {
		err = errors.New("the new password is the same as the old one")
		utils.SendError(w, m.App.Logger, err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestingUser.NewPassword), 12)
	if err != nil {
		utils.SendError(w, m.App.Logger, err, http.StatusInternalServerError)
		return
	}

	resetPasswordToken = ""
	_, err = m.DB.UpdateUserById(user.ID.Hex(), &models.User{
		Password:             string(hashedPassword),
		ResetPasswordToken:   &resetPasswordToken,
		ResetPasswordExpires: time.Now(),
	})
	if err != nil {
		utils.SendError(w, m.App.Logger, err, http.StatusInternalServerError)
		return
	}

	mail := &models.Mail{
		From:    m.App.Config.EmailProvider.Email,
		To:      []string{user.Email},
		Subject: "Password was changed",
		Body:    "Your password has been changed. Please contact an administrator if you were not the initiator!",
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
		Message: "Password successfully changed",
	})
}
