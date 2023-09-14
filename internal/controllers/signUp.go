package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
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
		m.App.Logger.Println(errors.New("user already exists"))
		utils.SendError(w, err)
		return
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
	userID, err := m.DB.CreateUser(*user)
	if err != nil {
		m.App.Logger.Println(err)
		utils.SendError(w, err, http.StatusInternalServerError)
		return
	}
	user.ID = userID
	user.Password = ""

	// create cookie
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: currentTime.Add(time.Hour * 24).Unix(),
		Id:        userID.Hex(),
		IssuedAt:  currentTime.Unix(),
		Issuer:    user.Email,
	})

	tokenString, err := token.SignedString(m.App.Config.JWT)
	if err != nil {
		m.App.Logger.Println(err)
		utils.SendError(w, err, http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:     m.App.Config.Cookie.Name,
		Path:     "/",
		Value:    tokenString,
		Expires:  currentTime.Add(24 * time.Hour),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	if m.App.Config.Cookie.SameSite == "none" {
		cookie.SameSite = http.SameSiteNoneMode
		cookie.Secure = true
	}

	http.SetCookie(w, cookie)

	// send back the user
	utils.Send(w, http.StatusOK, user)
}
