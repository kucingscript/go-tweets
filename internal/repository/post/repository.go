package post

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kucingscript/go-tweets/internal/model"
)

type PostRepository interface {
	StorePost(ctx context.Context, post *model.PostModel) (*model.PostModel, error)
	GetPostByID(ctx context.Context, postID int64) (*model.PostModel, error)
	UpdatePost(ctx context.Context, post *model.PostModel) (*model.PostModel, error)
	SoftDeletePost(ctx context.Context, postID int64) error
}

type postRepository struct {
	db *pgxpool.Pool
}

func NewPostRepository(db *pgxpool.Pool) PostRepository {
	return &postRepository{
		db: db,
	}
}
