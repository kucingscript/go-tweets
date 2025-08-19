package user

import (
	"github.com/kucingscript/go-tweets/internal/config"
	"github.com/kucingscript/go-tweets/internal/mailer"
	"github.com/kucingscript/go-tweets/internal/repository/user"
	"github.com/microcosm-cc/bluemonday"
)

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
