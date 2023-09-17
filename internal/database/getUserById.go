package database

import (
	"context"
	"time"

	"github.com/schattenbrot/auth/internal/models"
	"github.com/schattenbrot/auth/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *dbRepo) GetUserById(id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.DB.Collection("users")

	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := models.User{ID: &userId, Inactive: utils.BoolPointer(false)}

	var user models.User
	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
