package models

type AuthUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ResetPasswordRequestUser struct {
	Email string `json:"email" validate:"required"`
}

type ResetPasswordByTokenUser struct {
	NewPassword string `json:"newPassword" validate:"required"`
}
