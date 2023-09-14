package models

type AuthUser struct {
	Email    string `json:"email,omitempty" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required,min=8,max=20"`
}
