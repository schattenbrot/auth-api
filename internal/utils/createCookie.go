package utils

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateCookie(currentTime time.Time, userId string, jwtSecret []byte, cookieName string, cookieSameSite string) (*http.Cookie, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: currentTime.Add(time.Hour * 24).Unix(),
		IssuedAt:  currentTime.Unix(),
		Issuer:    userId,
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name:     cookieName,
		Path:     "/",
		Value:    tokenString,
		Expires:  currentTime.Add(24 * time.Hour),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	if cookieSameSite == "none" {
		cookie.SameSite = http.SameSiteNoneMode
		cookie.Secure = true
	}

	return cookie, nil
}
