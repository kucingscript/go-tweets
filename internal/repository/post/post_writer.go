package post

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
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

func (r *postRepository) SoftDeletePost(ctx context.Context, postID int64) error {
	query := `UPDATE posts SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL`

	cmdTag, err := r.db.Exec(ctx, query, postID)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return errors.New("post not found or already deleted")
	}

	return nil
}
