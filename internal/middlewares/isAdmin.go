package middlewares

import (
	"errors"
	"net/http"

	"github.com/schattenbrot/auth/internal/models"
	"github.com/schattenbrot/auth/internal/utils"
)

func (m *Repository) IsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(m.App.UserContextKey).(*models.User)

		isAdmin := false
		for _, role := range user.Roles {
			if role == "admin" {
				isAdmin = true
			}
		}

		if !isAdmin {
			err := errors.New("not admin")
			utils.SendError(w, m.App.Logger, err, http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
