package post

import (
	"github.com/go-playground/validator/v10"
	"github.com/kucingscript/go-tweets/internal/config"
	"github.com/kucingscript/go-tweets/internal/service/post"
)

type Handler struct {
	validate    *validator.Validate
	postService post.PostService
	cfg         *config.Config
}

func NewPostHandler(validate *validator.Validate, postService post.PostService, cfg *config.Config) *Handler {
	return &Handler{
		validate:    validate,
		postService: postService,
		cfg:         cfg,
	}
}
