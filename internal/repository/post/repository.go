package post

import (
	"context"

	"github.com/kucingscript/go-tweets/internal/model"
)

type PostReader interface {
	GetPostByID(ctx context.Context, postID int64) (*model.PostModel, error)
}

type PostWriter interface {
	StorePost(ctx context.Context, post *model.PostModel) (*model.PostModel, error)
	UpdatePost(ctx context.Context, post *model.PostModel) (*model.PostModel, error)
	SoftDeletePost(ctx context.Context, postID int64) error
}

type PostRepository interface {
	PostReader
	PostWriter
}
