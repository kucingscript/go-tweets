package post

import "github.com/jackc/pgx/v5/pgxpool"

type postRepository struct {
	db *pgxpool.Pool
}

func NewPostRepository(db *pgxpool.Pool) PostRepository {
	return &postRepository{
		db: db,
	}
}
