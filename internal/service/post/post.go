package post

import (
	"context"

	"github.com/kucingscript/go-tweets/internal/dto"
)

func (s *postService) CreatePost(ctx context.Context, req *dto.CreatePostRequest, userID int) (int64, int, error) {
	return 0, 0, nil
}
