package controllers

import (
	"errors"
	"net/http"
	"time"

	"github.com/schattenbrot/auth/internal/models"
	"github.com/schattenbrot/auth/internal/utils"
)

func (m *Repository) ActivateEmail(w http.ResponseWriter, r *http.Request) {
	emailActivationToken := r.URL.Query().Get("token")

	user, err := m.DB.GetUserByActivationToken(emailActivationToken)
	if err != nil {
		utils.SendError(w, m.App.Logger, err)
		return
	}

	m.App.Logger.Println(user.EmailActivateExpires)
	m.App.Logger.Println(time.Now())

	if user.EmailActivateExpires.Before(time.Now()) {
		err = errors.New("email activation token expired")
		utils.SendError(w, m.App.Logger, err)
		return
	}

	user.EmailActivated = true
	user.EmailActivateToken = ""
	user.EmailActivateExpires = time.Now()

	_, err = m.DB.UpdateUserById(user.ID.Hex(), user)
	if err != nil {
		utils.SendError(w, m.App.Logger, err)
		return
	}

	from := m.App.Config.EmailProvider.Email
	password := m.App.Config.EmailProvider.Password

	mail := &models.Mail{
		From:    from,
		To:      []string{user.Email},
		Subject: "Your email address was successfully activated!",
		Body:    `Your email address was successfully activated!`,
	}
	err = mail.Send(from, password)
	if err != nil {
		utils.SendError(w, m.App.Logger, err)
		return
	}

	type resp struct {
		message string
	}

	utils.Send(w, m.App.Logger, http.StatusOK, &resp{
		message: "Your email address was successfully activated!",
	})
}
