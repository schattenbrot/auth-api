package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/schattenbrot/auth/internal/models"
	"github.com/schattenbrot/auth/internal/utils"
)

func (m *Repository) UpdateMeAvatar(w http.ResponseWriter, r *http.Request) {
	currentUser := r.Context().Value(m.App.UserContextKey).(*models.User)
	m.App.Logger.Println(currentUser)

	defer r.Body.Close()
	r.ParseForm()
	fileAvatar, fileHeader, err := r.FormFile("file")
	if err != nil {
		utils.SendError(w, m.App.Logger, err)
		return
	}

	fileType, err := utils.GetFileType(fileHeader)
	if err != nil {
		utils.SendError(w, m.App.Logger, err)
		return
	}
	filename := fmt.Sprintf("%d%s", time.Now().UnixMilli(), fileType)

	defer fileAvatar.Close()
	err = utils.SaveFiles(filename, fileAvatar)
	if err != nil {
		utils.SendError(w, m.App.Logger, err)
		return
	}

	filePath := fmt.Sprintf("%s:%d/users/avatar/%s", m.App.Config.Servername, m.App.Config.Port, filename)
	updatedUser, err := m.DB.UpdateUserById(currentUser.ID.Hex(), &models.User{
		Avatar: filePath,
	})
	if err != nil {
		utils.SendError(w, m.App.Logger, err)
		return
	}

	utils.Send(w, m.App.Logger, http.StatusOK, updatedUser)
}
