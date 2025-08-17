package user

import (
	"context"

	"github.com/kucingscript/go-tweets/internal/config"
	"github.com/kucingscript/go-tweets/internal/dto"
	"github.com/kucingscript/go-tweets/internal/mailer"
	"github.com/kucingscript/go-tweets/internal/model"
	"github.com/kucingscript/go-tweets/internal/repository/user"
)

type UserService interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (*model.UserModel, int, error)
	VerifyEmail(ctx context.Context, token string) (int, error)

	Login(ctx context.Context, req *dto.LoginRequest) (string, string, int, error)

	ForgotPassword(ctx context.Context, email string) (int, error)
	ResetPassword(ctx context.Context, req *dto.ResetPasswordRequest) (int, error)
}

type userService struct {
	cfg            *config.Config
	userRepository user.UserRepository
	mailer         *mailer.Mailer
}

func NewUserService(cfg *config.Config, userRepository user.UserRepository, mailer *mailer.Mailer) UserService {
	return &userService{
		cfg:            cfg,
		userRepository: userRepository,
		mailer:         mailer,
	}
}
