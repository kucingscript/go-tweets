package post

import (
	"context"

	"github.com/kucingscript/go-tweets/internal/config"
	"github.com/kucingscript/go-tweets/internal/dto"
	"github.com/kucingscript/go-tweets/internal/model"
	"github.com/kucingscript/go-tweets/internal/repository/post"
	"github.com/microcosm-cc/bluemonday"
)

type PostService interface {
	CreatePost(ctx context.Context, req *dto.CreateOrUpdatePostRequest, userID int64) (*model.PostModel, int, error)
	UpdatePost(ctx context.Context, req *dto.CreateOrUpdatePostRequest, postID, userID int64) (*model.PostModel, int, error)
	DeletePost(ctx context.Context, postID, userID int64) (int, error)
}

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
