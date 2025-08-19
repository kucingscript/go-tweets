package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/kucingscript/go-tweets/internal/config"
	"github.com/kucingscript/go-tweets/internal/service/user"
)

type Handler struct {
	validate    *validator.Validate
	userService user.UserService
	cfg         *config.Config
}

func NewUserHandler(validate *validator.Validate, userService user.UserService, cfg *config.Config) *Handler {
	return &Handler{
		validate:    validate,
		userService: userService,
		cfg:         cfg,
	}
}
