package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/kucingscript/go-tweets/internal/model"
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

func (r *userRepository) GetUserByVerificationToken(ctx context.Context, token string) (*model.UserModel, error) {
	query := `SELECT id, email, username, is_verified FROM users WHERE verification_token = $1`
	row := r.db.QueryRow(ctx, query, token)

	var user model.UserModel
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.IsVerified)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

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
