package controllers

import (
	"errors"
	"net/http"

	"github.com/schattenbrot/auth/internal/utils"
)

func (m *Repository) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := m.DB.GetUsers()
	if err != nil {
		utils.SendError(w, m.App.Logger, err)
		return
	}

	if users == nil {
		err = errors.New("users not found")
		utils.SendError(w, m.App.Logger, err)
		return
	}

	utils.Send(w, m.App.Logger, http.StatusOK, users)
}
