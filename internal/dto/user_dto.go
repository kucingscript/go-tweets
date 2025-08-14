package dto

import "time"

type (
	RegisterRequest struct {
		Email           string `json:"email" validate:"required,email"`
		Username        string `json:"username" validate:"required,min=3"`
		Password        string `json:"password" validate:"required,min=8"`
		PasswordConfirm string `json:"password_confirm" validate:"required,eqfield=Password"`
	}

	RegisterResponse struct {
		ID         int64     `json:"id"`
		Email      string    `json:"email"`
		Username   string    `json:"username"`
		IsVerified bool      `json:"is_verified"`
		CreatedAt  time.Time `json:"created_at"`
	}
)

type (
	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}

	LoginResponse struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
)
