package user

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kucingscript/go-tweets/internal/model"
)

type UserRepository interface {
	GetUserByEmailOrUsername(ctx context.Context, email, username string) (*model.UserModel, error)
	CreateUser(ctx context.Context, user *model.UserModel) error
	UpdateVerificationToken(ctx context.Context, userID int64, token string) error
	GetUserByVerificationToken(ctx context.Context, token string) (*model.UserModel, error)
	VerifyUser(ctx context.Context, userID int64) error
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{
		db: db,
	}
}
