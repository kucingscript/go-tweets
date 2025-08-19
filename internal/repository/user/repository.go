package user

import (
	"context"
	"time"

	"github.com/kucingscript/go-tweets/internal/model"
)

type UserReader interface {
	GetUserByEmail(ctx context.Context, email string) (*model.UserModel, error)
	GetUserByVerificationToken(ctx context.Context, token string) (*model.UserModel, error)
	GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*model.RefreshTokenModel, error)
	GetUserByRefreshToken(ctx context.Context, token string) (*model.UserModel, error)
	GetUserByResetToken(ctx context.Context, token string) (*model.UserModel, error)
}

type userWriter interface {
	CreateUser(ctx context.Context, user *model.UserModel) error
	UpdateVerificationToken(ctx context.Context, userID int64, token string) error
	VerifyUser(ctx context.Context, userID int64) error
	StoreRefreshToken(ctx context.Context, refreshToken *model.RefreshTokenModel) error
	DeleteRefreshToken(ctx context.Context, token string) error
	DeleteExpiredRefreshTokens(ctx context.Context) (int64, error)
	SetPasswordResetToken(ctx context.Context, userID int64, token string, expiredAt time.Time) error
	UpdatePassword(ctx context.Context, userID int64, password string) error
}

type UserRepository interface {
	UserReader
	userWriter
}
