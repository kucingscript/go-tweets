package post

import (
	"context"
	"errors"
)

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
