package database

import (
	"context"
	"time"

	"github.com/schattenbrot/auth/internal/models"
	"github.com/schattenbrot/auth/internal/utils"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m *dbRepo) GetUsers() ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := m.DB.Collection("users")

	filters := models.User{Inactive: utils.BoolPointer(false)}

	cursor, err := collection.Find(ctx, filters, options.Find())
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*models.User

	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
