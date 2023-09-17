package controllers

import (
	"net/http"
	"time"

	"github.com/schattenbrot/auth/internal/utils"
)

type apiStatus struct {
	Status      string        `json:"status"`
	Uptime      time.Duration `json:"uptime"`
	Environment string        `json:"environment"`
	Version     string        `json:"version"`
}

func (m *Repository) ApiStatus(w http.ResponseWriter, r *http.Request) {
	status := apiStatus{
		Status:      "Available",
		Uptime:      time.Duration(time.Since(m.App.ServerStartTime).Minutes()),
		Environment: m.App.Config.Env,
		Version:     m.App.Version,
	}

	utils.Send(w, m.App.Logger, http.StatusOK, status)
}
