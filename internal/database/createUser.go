package database

import (
	"context"
	"time"

	"github.com/schattenbrot/auth/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *dbRepo) CreateUser(user models.User) (*primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := m.DB.Collection("users")

	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	// oid := res.InsertedID.(primitive.ObjectID).Hex()
	oid := res.InsertedID.(primitive.ObjectID)

	return &oid, nil
}
