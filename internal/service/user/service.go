package user

import (
	"context"

	"github.com/kucingscript/go-tweets/internal/config"
	"github.com/kucingscript/go-tweets/internal/dto"
	"github.com/kucingscript/go-tweets/internal/mailer"
	"github.com/kucingscript/go-tweets/internal/model"
	"github.com/kucingscript/go-tweets/internal/repository/user"
	"github.com/microcosm-cc/bluemonday"
)

type UserService interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (*model.UserModel, int, error)
	VerifyEmail(ctx context.Context, token string) (int, error)

	Login(ctx context.Context, req *dto.LoginRequest) (string, string, int, error)
	Logout(ctx context.Context, refreshToken string) (int, error)
	RefreshToken(ctx context.Context, refreshToken string) (string, int, error)
	CleanUpExpiredTokens(ctx context.Context)

	ForgotPassword(ctx context.Context, email string) (int, error)
	ResetPassword(ctx context.Context, req *dto.ResetPasswordRequest) (int, error)
}

type userService struct {
	cfg            *config.Config
	userRepository user.UserRepository
	mailer         *mailer.Mailer
	htmlSanitizer  *bluemonday.Policy
}

func NewUserService(cfg *config.Config, userRepository user.UserRepository, mailer *mailer.Mailer) UserService {
	sanitizer := bluemonday.UGCPolicy()

	return &userService{
		cfg:            cfg,
		userRepository: userRepository,
		mailer:         mailer,
		htmlSanitizer:  sanitizer,
	}
}
