package middlewares

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/schattenbrot/auth/internal/utils"
)

func (m *Repository) IsAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookieClaims, err := utils.GetIssuerFromCookie(r, m.App.Config.Cookie.Name, m.App.Config.JWT)
		if err != nil {
			err := errors.New("authorization cookie not found")
			utils.SendError(w, m.App.Logger, err, http.StatusUnauthorized)
			return
		}

		if time.Now().Unix() >= cookieClaims.ExpiresAt {
			err := errors.New("authorization cookie expired")
			utils.SendError(w, m.App.Logger, err, http.StatusUnauthorized)
			return
		}

		user, err := m.DB.GetUserById(cookieClaims.Issuer)
		if err != nil {
			err := errors.New("authorization user not found")
			utils.SendError(w, m.App.Logger, err, http.StatusUnauthorized)
			return
		}

		// set user in context
		ctxWithUser := context.WithValue(r.Context(), m.App.UserContextKey, user)

		next.ServeHTTP(w, r.WithContext(ctxWithUser))
	})
}
