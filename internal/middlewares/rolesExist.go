package middlewares

import (
	"errors"
	"net/http"
	"slices"

	"github.com/schattenbrot/auth/internal/models"
	"github.com/schattenbrot/auth/internal/utils"
)

func (m *Repository) RolesExist(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		possibleRoles := append(m.App.Config.Roles.Additional, "admin")

		var user models.User
		utils.MiddlewareBodyDecoder(r, &user)

		wrongRoles := []string{}
		for _, role := range user.Roles {
			if !slices.Contains(possibleRoles, role) {
				wrongRoles = append(wrongRoles, role)
			}
		}

		if len(wrongRoles) != 0 {
			err := errors.New("selected roles do not exist")
			utils.SendError(w, m.App.Logger, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}
