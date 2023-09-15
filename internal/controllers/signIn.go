package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/schattenbrot/auth/internal/models"
	"github.com/schattenbrot/auth/internal/utils"
)

func (m *Repository) SignIn(w http.ResponseWriter, r *http.Request) {
	var authUser models.AuthUser
	err := json.NewDecoder(r.Body).Decode(&authUser)
	if err != nil {
		m.App.Logger.Println(err)
		utils.SendError(w, err)
		return
	}

	err = m.App.Validator.Struct(authUser)
	if err != nil {
		m.App.Logger.Println(err)
		utils.SendError(w, err)
		return
	}

	// check inactive
	user, err := m.DB.GetUserByEmail(authUser.Email)
	if err != nil {
		m.App.Logger.Println(err)
		utils.SendError(w, err)
		return
	}

	currentTime := time.Now()
	cookie, err := utils.CreateCookie(currentTime, user.ID.Hex(), user.Email, m.App.Config.JWT, m.App.Config.Cookie.Name, m.App.Config.Cookie.SameSite)
	if err != nil {
		m.App.Logger.Println(err)
		utils.SendError(w, err, http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, cookie)

	utils.Send(w, http.StatusOK, user)
}
