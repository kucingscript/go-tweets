package post

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/kucingscript/go-tweets/internal/model"
)

func (r *postRepository) GetPostByID(ctx context.Context, postID int64) (*model.PostModel, error) {
	query := `SELECT id, user_id, title, content, created_at, updated_at, deleted_at FROM posts 
			WHERE id = $1
			AND deleted_at IS NULL`

	row := r.db.QueryRow(ctx, query, postID)
	var post model.PostModel
	err := row.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.CreatedAt, &post.UpdatedAt, &post.DeletedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &post, nil
}
