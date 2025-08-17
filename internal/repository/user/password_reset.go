package user

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/kucingscript/go-tweets/internal/model"
)

func (r *userRepository) SetPasswordResetToken(ctx context.Context, userID int64, token string, expiresAt time.Time) error {
	query := `UPDATE users SET password_reset_token = $1, password_reset_token_expires_at = $2 WHERE id = $3`
	_, err := r.db.Exec(ctx, query, token, expiresAt, userID)
	return err
}

func (r *userRepository) GetUserByResetToken(ctx context.Context, token string) (*model.UserModel, error) {
	query := `SELECT id, username, email, password_reset_token_expires_at FROM users WHERE password_reset_token = $1`
	row := r.db.QueryRow(ctx, query, token)

	var user model.UserModel
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordResetTokenExpiresAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdatePassword(ctx context.Context, userID int64, password string) error {
	query := `UPDATE users SET password = $1, password_reset_token = NULL, password_reset_token_expires_at = NULL WHERE id = $2`
	_, err := r.db.Exec(ctx, query, password, userID)
	return err
}
