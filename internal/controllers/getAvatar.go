package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/schattenbrot/auth/internal/utils"
)

func (m *Repository) GetAvatar(w http.ResponseWriter, r *http.Request) {
	filename := chi.URLParam(r, "filename")

	file, err := utils.GetFile(filename)
	if err != nil {
		utils.SendError(w, m.App.Logger, err)
		return
	}

	utils.SendFile(w, m.App.Logger, file)
}
