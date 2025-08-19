package user

import (
	"context"

	"github.com/kucingscript/go-tweets/internal/dto"
	"github.com/kucingscript/go-tweets/internal/model"
)

type UserService interface {
	// Account Service
	Register(ctx context.Context, req *dto.RegisterRequest) (*model.UserModel, int, error)
	VerifyEmail(ctx context.Context, token string) (int, error)

	// Auth Service
	Login(ctx context.Context, req *dto.LoginRequest) (string, string, int, error)
	Logout(ctx context.Context, refreshToken string) (int, error)
	RefreshToken(ctx context.Context, refreshToken string) (string, int, error)
	CleanUpExpiredTokens(ctx context.Context)

	// Password Service
	ForgotPassword(ctx context.Context, email string) (int, error)
	ResetPassword(ctx context.Context, req *dto.ResetPasswordRequest) (int, error)
}
