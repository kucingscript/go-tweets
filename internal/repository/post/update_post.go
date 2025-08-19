package post

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/kucingscript/go-tweets/internal/model"
)

func (r *postRepository) UpdatePost(ctx context.Context, post *model.PostModel) (*model.PostModel, error) {
	query := `UPDATE posts SET title = $1, content = $2, updated_at = NOW()
			WHERE id = $3 AND deleted_at IS NULL
			RETURNING title, content, updated_at`

	row := r.db.QueryRow(ctx, query, post.Title, post.Content, post.ID)

	err := row.Scan(&post.Title, &post.Content, &post.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("post not found")
		}
		return nil, err
	}

	return post, nil
}
