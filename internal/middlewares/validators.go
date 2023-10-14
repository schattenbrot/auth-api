package middlewares

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/schattenbrot/auth/internal/models"
	"github.com/schattenbrot/auth/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (m *Repository) ValidateUpdateMeAvatar(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		r.ParseForm()
		_, _, err := r.FormFile("file")
		if err != nil {
			utils.SendError(w, m.App.Logger, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// func (m *Repository) RequiredId() func(http.Handler) http.Handler {
func (m *Repository) RequiredId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		value := chi.URLParam(r, "id")
		if value == "" {
			err := errors.New("mongoid required")
			utils.SendError(w, m.App.Logger, err)
			return
		}

		_, err := primitive.ObjectIDFromHex(value)
		if err != nil {
			err := errors.New("mongoid required")
			utils.SendError(w, m.App.Logger, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (m *Repository) ValidateUpdateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user models.UpdateUserUser
		utils.MiddlewareBodyDecoder(r, &user)

		err := m.App.Validator.Struct(user)
		if err != nil {
			utils.SendError(w, m.App.Logger, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (m *Repository) ValidateResetPasswordRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user models.ResetPasswordRequestUser
		utils.MiddlewareBodyDecoder(r, &user)

		err := m.App.Validator.Struct(user)
		if err != nil {
			utils.SendError(w, m.App.Logger, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (m *Repository) ValidateResetPasswordByToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user models.ResetPasswordByTokenUser
		utils.MiddlewareBodyDecoder(r, &user)

		err := m.App.Validator.Struct(user)
		if err != nil {
			utils.SendError(w, m.App.Logger, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}
