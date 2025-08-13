package user

import (
	"context"
	"fmt"
	"time"
)

func (r *userRepository) VerifyUser(ctx context.Context, userID int64) error {
	query := `UPDATE users 
			  SET is_verified = TRUE, verified_at = $1, verification_token = NULL
			  WHERE id = $2`

	cmdTag, err := r.db.Exec(ctx, query, time.Now(), userID)
	if err != nil {
		return fmt.Errorf("unable to verify user: %w", err)
	}
	if cmdTag.RowsAffected() != 1 {
		return fmt.Errorf("no user found to verify")
	}

	return nil
}
