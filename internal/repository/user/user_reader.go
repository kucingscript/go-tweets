package user

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/kucingscript/go-tweets/internal/model"
)

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*model.UserModel, error) {
	query := `SELECT id, email, username, password, is_verified, created_at, updated_at
			FROM users
			WHERE email = $1`

	row := r.db.QueryRow(ctx, query, email)

	var user model.UserModel
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
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
