package post

import (
	"context"

	"github.com/kucingscript/go-tweets/internal/model"
)

func (r *postRepository) StorePost(ctx context.Context, post *model.PostModel) (*model.PostModel, error) {
	query := `INSERT INTO posts (user_id, title, content) VALUES ($1, $2, $3)
			RETURNING id, created_at, updated_at, deleted_at`

	row := r.db.QueryRow(ctx, query, post.UserID, post.Title, post.Content)

	err := row.Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt, &post.DeletedAt)
	if err != nil {
		return nil, err
	}

	return post, nil
}
