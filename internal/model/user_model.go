package model

import "time"

type (
	UserModel struct {
		ID       int64  `json:"id"`
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"-"`

		IsVerified        bool       `json:"is_verified"`
		VerificationToken string     `json:"-"`
		VerifiedAt        *time.Time `json:"verified_at,omitempty"`

		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	RefreshTokenModel struct {
		ID           int64     `json:"id"`
		UserID       int64     `json:"user_id"`
		RefreshToken string    `json:"refresh_token"`
		ExpiredAt    time.Time `json:"expired_at"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
)
