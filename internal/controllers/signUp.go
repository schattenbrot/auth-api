package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/schattenbrot/auth/internal/models"
	"github.com/schattenbrot/auth/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func (m *Repository) SignUp(w http.ResponseWriter, r *http.Request) {
	var authUser models.AuthUser
	err := json.NewDecoder(r.Body).Decode(&authUser)
	if err != nil {
		utils.SendError(w, m.App.Logger, err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(authUser.Password), 12)
	if err != nil {
		utils.SendError(w, m.App.Logger, err, http.StatusInternalServerError)
		return
	}

	currentTime := time.Now().UTC()

	user := &models.User{
		Email:     authUser.Email,
		Password:  string(hashedPassword),
		Roles:     []string{m.App.Config.Roles.Default},
		Inactive:  utils.BoolPointer(false),
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	userID, err := m.DB.CreateUser(*user)
	if err != nil {
		utils.SendError(w, m.App.Logger, err, http.StatusInternalServerError)
		return
	}
	user.ID = userID
	user.Password = ""

	cookie, err := utils.CreateCookie(currentTime, userID.Hex(), m.App.Config.JWT, m.App.Config.Cookie.Name, m.App.Config.Cookie.SameSite)
	if err != nil {
		utils.SendError(w, m.App.Logger, err, http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, cookie)

	utils.Send(w, m.App.Logger, http.StatusOK, user)
}
