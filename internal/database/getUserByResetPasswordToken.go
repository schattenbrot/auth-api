package database

import (
	"context"
	"time"

	"github.com/schattenbrot/auth/internal/models"
)

func (m *dbRepo) GetUserByResetPasswordToken(token string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.DB.Collection("users")

	filter := models.User{ResetPasswordToken: &token}

	var user models.User
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
