package user

import (
	"context"

	"github.com/kucingscript/go-tweets/internal/config"
	"github.com/kucingscript/go-tweets/internal/dto"
	"github.com/kucingscript/go-tweets/internal/repository/user"
)

type UserService interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (int64, int, error)
}

type userService struct {
	cfg            *config.Config
	userRepository user.UserRepository
}

func NewUserService(cfg *config.Config, userRepository user.UserRepository) UserService {
	return &userService{cfg: cfg, userRepository: userRepository}
}
