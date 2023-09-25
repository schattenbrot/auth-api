package database

import (
	"context"
	"time"

	"github.com/schattenbrot/auth/internal/models"
	"github.com/schattenbrot/auth/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *dbRepo) DeleteUserById(id string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.DB.Collection("users")

	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := models.User{ID: &userId}

	var updatedUser models.User
	err = collection.FindOneAndUpdate(ctx, filter, models.User{Inactive: utils.BoolPointer(true)}, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&updatedUser)
	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}
