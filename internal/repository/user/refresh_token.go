package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/kucingscript/go-tweets/internal/model"
)

func (r *userRepository) GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*model.RefreshTokenModel, error) {
	query := `SELECT id, user_id, refresh_token, expired_at, created_at, updated_at 
			FROM refresh_tokens WHERE user_id = $1 
			AND expired_at > $2`

	row := r.db.QueryRow(ctx, query, userID, now)
	var refreshToken model.RefreshTokenModel
	err := row.Scan(&refreshToken.ID, &refreshToken.UserID, &refreshToken.RefreshToken, &refreshToken.ExpiredAt, &refreshToken.CreatedAt, &refreshToken.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &refreshToken, nil
}

func (r *userRepository) GetUserByRefreshToken(ctx context.Context, token string) (*model.UserModel, error) {
	query := `SELECT u.id, u.email, u.username
			  FROM users u
			  JOIN refresh_tokens rt ON u.id = rt.user_id
			  WHERE rt.refresh_token = $1 AND rt.expired_at > $2`

	row := r.db.QueryRow(ctx, query, token, time.Now())

	var user model.UserModel
	err := row.Scan(&user.ID, &user.Email, &user.Username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
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
