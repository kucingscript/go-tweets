package post

import (
	"context"

	"github.com/kucingscript/go-tweets/internal/config"
	"github.com/kucingscript/go-tweets/internal/dto"
	"github.com/kucingscript/go-tweets/internal/repository/post"
)

type PostService interface {
	CreatePost(ctx context.Context, req *dto.CreatePostRequest, userID int) (int64, int, error)
}

type postService struct {
	cfg            *config.Config
	postRepository post.PostRepository
}

func NewPostService(cfg *config.Config, postRepository post.PostRepository) PostService {
	return &postService{
		cfg:            cfg,
		postRepository: postRepository,
	}
}
