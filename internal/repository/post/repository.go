package post

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kucingscript/go-tweets/internal/model"
)

type PostRepository interface {
	StorePost(ctx context.Context, post *model.PostModel) (int64, error)
}

type postRepository struct {
	db *pgxpool.Pool
}

func NewPostRepository(db *pgxpool.Pool) PostRepository {
	return &postRepository{
		db: db,
	}
}
