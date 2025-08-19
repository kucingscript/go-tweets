package post

import (
	"context"

	"github.com/kucingscript/go-tweets/internal/dto"
	"github.com/kucingscript/go-tweets/internal/model"
)

type PostService interface {
	CreatePost(ctx context.Context, req *dto.CreateOrUpdatePostRequest, userID int64) (*model.PostModel, int, error)
	UpdatePost(ctx context.Context, req *dto.CreateOrUpdatePostRequest, postID, userID int64) (*model.PostModel, int, error)
	DeletePost(ctx context.Context, postID, userID int64) (int, error)
}
