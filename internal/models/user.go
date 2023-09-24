package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                   *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username             string              `json:"username,omitempty" bson:"username,omitempty"`
	Email                string              `json:"email,omitempty" bson:"email,omitempty"`
	EmailActiated        bool                `json:"-" bson:"emailActivated,omitempty"`
	Password             string              `json:"-" bson:"password,omitempty"`
	Avatar               string              `json:"avatar,omitempty" bson:"avatar,omitempty"`
	Roles                []string            `json:"roles,omitempty" bson:"roles,omitempty"`
	Inactive             *bool               `json:"-" bson:"inactive,omitempty"`
	ResetPasswordToken   string              `json:"-" bson:"resetPasswordToken,omitempty"`
	ResetPasswordExpires time.Time           `json:"-" bson:"resetPasswordExpires,omitempty"`
	CreatedAt            time.Time           `json:"-" bson:"createdAt,omitempty"`
	UpdatedAt            time.Time           `json:"-" bson:"updatedAt,omitempty"`
}

type UpdateMeUsernameUser struct {
	Username string `json:"username" validate:"required,min=3"`
}

type UpdateMeEmailUser struct {
	Email string `json:"email" validate:"required,email"`
}
