package controllers

import (
	"net/http"
	"time"

	"github.com/schattenbrot/auth/internal/utils"
)

func (m *Repository) SignOut(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Add(-2 * 24 * time.Hour)
	cookie, err := utils.CreateCookie(currentTime, "123", m.App.Config.JWT, m.App.Config.Cookie.Name, m.App.Config.Cookie.SameSite)
	if err != nil {
		utils.SendError(w, m.App.Logger, err, http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, cookie)

	type jsonResponse struct {
		Message string `json:"message"`
	}

	utils.Send(w, m.App.Logger, http.StatusOK, &jsonResponse{Message: "user got signed out"})
}
