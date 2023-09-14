package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/schattenbrot/auth/internal/models"
	"github.com/schattenbrot/auth/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func (m *Repository) SignUp(w http.ResponseWriter, r *http.Request) {
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

	var user *models.User
	user, err = m.DB.GetUserByEmail(authUser.Email)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			m.App.Logger.Println(err)
			utils.SendError(w, err)
			return
		}
	}
	if user != nil {
		err = errors.New("user exists already")
		m.App.Logger.Println(err)
		utils.SendError(w, err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(authUser.Password), 12)
	if err != nil {
		m.App.Logger.Println(err)
		utils.SendError(w, err, http.StatusInternalServerError)
		return
	}

	currentTime := time.Now().UTC()

	user = &models.User{
		Email:     authUser.Email,
		Password:  string(hashedPassword),
		Roles:     []string{m.App.Config.Roles.Default},
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	// save user into database
	// m.DB.CreateUser(user)

	// create cookie

	// send back the user

	utils.Send(w, http.StatusOK, user)
}
