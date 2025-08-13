package user

import (
	"context"
	"fmt"
)

func (r *userRepository) UpdateVerificationToken(ctx context.Context, userID int64, token string) error {
	query := `UPDATE users SET verification_token = $1 WHERE id = $2`

	cmdTag, err := r.db.Exec(ctx, query, token, userID)
	if err != nil {
		return fmt.Errorf("unable to update verification token: %w", err)
	}

	if cmdTag.RowsAffected() != 1 {
		return fmt.Errorf("no user found with the given ID to update")
	}

	return nil
}
