package post

import (
	"github.com/kucingscript/go-tweets/internal/config"
	"github.com/kucingscript/go-tweets/internal/repository/post"
	"github.com/microcosm-cc/bluemonday"
)

type postService struct {
	cfg            *config.Config
	postRepository post.PostRepository
	htmlSanitizer  *bluemonday.Policy
}

func NewPostService(cfg *config.Config, postRepository post.PostRepository) PostService {
	sanitizer := bluemonday.UGCPolicy()

	return &postService{
		cfg:            cfg,
		postRepository: postRepository,
		htmlSanitizer:  sanitizer,
	}
}
