package user

import (
	"context"
	"fmt"
	"time"

	"github.com/kucingscript/go-tweets/internal/model"
)

func (r *userRepository) CreateUser(ctx context.Context, user *model.UserModel) error {
	query := `INSERT INTO users (email, username, password) VALUES ($1, $2, $3)
			RETURNING id, created_at, updated_at, is_verified`

	row := r.db.QueryRow(ctx, query, user.Email, user.Username, user.Password)

	err := row.Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt, &user.IsVerified)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

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

func (r *userRepository) StoreRefreshToken(ctx context.Context, refreshToken *model.RefreshTokenModel) error {
	query := `INSERT INTO refresh_tokens (user_id, refresh_token, expired_at) VALUES ($1, $2, $3)
			RETURNING id, created_at, updated_at`

	err := r.db.QueryRow(ctx, query, refreshToken.UserID, refreshToken.RefreshToken, refreshToken.ExpiredAt).Scan(&refreshToken.ID, &refreshToken.CreatedAt, &refreshToken.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to store refresh token: %w", err)
	}

	return nil

}

func (r *userRepository) DeleteRefreshToken(ctx context.Context, token string) error {
	query := "DELETE FROM refresh_tokens WHERE refresh_token = $1"

	cmdTag, err := r.db.Exec(ctx, query, token)
	if err != nil {
		return fmt.Errorf("failed to delete refresh token: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return nil
	}

	return nil
}

func (r *userRepository) DeleteExpiredRefreshTokens(ctx context.Context) (int64, error) {
	query := "DELETE FROM refresh_tokens WHERE expired_at < $1"

	cmdTag, err := r.db.Exec(ctx, query, time.Now())
	if err != nil {
		return 0, fmt.Errorf("failed to delete expired refresh tokens: %w", err)
	}

	return cmdTag.RowsAffected(), nil
}

func (r *userRepository) SetPasswordResetToken(ctx context.Context, userID int64, token string, expiresAt time.Time) error {
	query := `UPDATE users SET password_reset_token = $1, password_reset_token_expires_at = $2 WHERE id = $3`
	_, err := r.db.Exec(ctx, query, token, expiresAt, userID)
	return err
}

func (r *userRepository) UpdatePassword(ctx context.Context, userID int64, password string) error {
	query := `UPDATE users SET password = $1, password_reset_token = NULL, password_reset_token_expires_at = NULL WHERE id = $2`
	_, err := r.db.Exec(ctx, query, password, userID)
	return err
}
