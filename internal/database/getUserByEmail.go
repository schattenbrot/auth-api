package database

import (
	"context"
	"time"

	"github.com/schattenbrot/auth/internal/models"
)

func (m *dbRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.DB.Collection("users")

	var user models.User
	err := collection.FindOne(ctx, models.User{Email: email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
