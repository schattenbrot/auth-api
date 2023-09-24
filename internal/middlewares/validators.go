package middlewares

import (
	"net/http"

	"github.com/schattenbrot/auth/internal/models"
	"github.com/schattenbrot/auth/internal/utils"
)

func (m *Repository) ValidateAuthUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var authUser models.AuthUser
		utils.MiddlewareBodyDecoder(r, &authUser)

		err := m.App.Validator.Struct(authUser)
		if err != nil {
			utils.SendError(w, m.App.Logger, err)
			return
		}
		err = utils.PasswordValidator(authUser.Password)
		if err != nil {
			utils.SendError(w, m.App.Logger, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (m *Repository) ValidateUpdateMeUsername(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user models.UpdateMeUsernameUser
		utils.MiddlewareBodyDecoder(r, &user)

		err := m.App.Validator.Struct(user)
		if err != nil {
			utils.SendError(w, m.App.Logger, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (m *Repository) ValidateUpdateMeEmail(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user models.UpdateMeEmailUser
		utils.MiddlewareBodyDecoder(r, &user)

		err := m.App.Validator.Struct(user)
		if err != nil {
			utils.SendError(w, m.App.Logger, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (m *Repository) ValidateUpdateMePassword(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user models.UpdateMePasswordUser
		utils.MiddlewareBodyDecoder(r, &user)

		err := m.App.Validator.Struct(user)
		if err != nil {
			utils.SendError(w, m.App.Logger, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}
