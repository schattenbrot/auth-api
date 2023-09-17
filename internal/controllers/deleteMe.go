package controllers

import (
	"net/http"

	"github.com/schattenbrot/auth/internal/models"
	"github.com/schattenbrot/auth/internal/utils"
)

func (m *Repository) DeleteMe(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(m.App.UserContextKey).(*models.User)

	_, err := m.DB.DeleteUserById(user.ID.Hex())
	if err != nil {
		utils.SendError(w, m.App.Logger, err, http.StatusInternalServerError)
		return
	}

	type jsonResponse struct {
		Message string `json:"message"`
	}

	response := &jsonResponse{
		Message: "user deleted",
	}

	utils.Send(w, m.App.Logger, http.StatusOK, response)
}
