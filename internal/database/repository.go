package database

import (
	"context"
	"time"

	"github.com/schattenbrot/auth/internal/config"
	"github.com/schattenbrot/auth/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DatabaseRepo is the interface for all repository functions
type DatabaseRepo interface {
	// UserRepo
	CreateUser(user models.User) (*primitive.ObjectID, error)
	GetUsers() ([]*models.User, error)
	GetUserById(id string) (*models.User, error)
	GetInactiveUserById(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByActivationToken(token string) (*models.User, error)
	GetUserByResetPasswordToken(token string) (*models.User, error)
	UpdateUserById(id string, user *models.User) (*models.User, error)
	DeleteUserById(id string) (*models.User, error)
}

type dbRepo struct {
	App *config.AppConfig
	DB  *mongo.Database
}

// NewMongoDBRepo is the function for returning a mongoDBRepo.
func NewDBRepo(app *config.AppConfig, conn *mongo.Database) DatabaseRepo {
	return &dbRepo{
		App: app,
		DB:  conn,
	}
}

// openDB creates a new database connection and returns the Database
func OpenDB(app config.AppConfig) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(app.Config.DB.DSN))
	if err != nil {
		app.Logger.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		app.Logger.Fatal(err)
	}
	db := client.Database(app.Config.DB.Name)

	setIndizes(&app, db)

	return db
}
