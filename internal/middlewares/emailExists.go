package middlewares

import (
	"errors"
	"net/http"

	"github.com/schattenbrot/auth/internal/models"
	"github.com/schattenbrot/auth/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *Repository) EmailExists(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var authUser models.AuthUser
		utils.MiddlewareBodyDecoder(r, &authUser)

		var user *models.User
		user, err := m.DB.GetUserByEmail(authUser.Email)
		if err != nil {
			if err != mongo.ErrNoDocuments {
				utils.SendError(w, m.App.Logger, err)
				return
			}
		}
		if user != nil {
			err = errors.New("email is already being used")
			utils.SendError(w, m.App.Logger, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}
