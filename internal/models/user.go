package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                   *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username             string              `json:"username,omitempty" bson:"username, omitempty"`
	Email                string              `json:"email,omitempty" bson:"email,omitempty"`
	Password             string              `json:"password,omitempty" bson:"password,omitempty"`
	Avatar               string              `json:"avatar,omitempty" bson:"avatar,omitempty"`
	Roles                []string            `json:"roles,omitempty" bson:"roles,omitempty"`
	Deactivated          bool                `json:"deactivated,omitempty" bson:"deactiated,omitempty"`
	ResetPasswordToken   string              `json:"resetPasswordToken,omitempty" bson:"resetPasswordToken,omitempty"`
	ResetPasswordExpires *time.Time          `json:"resetPasswordExpires,omitempty" bson:"resetPasswordExpires,omitempty"`
	CreatedAt            time.Time           `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt            time.Time           `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	// GoogleAccessToken string             `json:"googleAccessToken,omitempty" bson:"googleAccessToken,omitempty"`
}
