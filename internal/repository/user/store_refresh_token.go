package user

import (
	"context"
	"fmt"

	"github.com/kucingscript/go-tweets/internal/model"
)

func (r *userRepository) StoreRefreshToken(ctx context.Context, refreshToken *model.RefreshTokenModel) error {
	query := `INSERT INTO refresh_tokens (user_id, refresh_token, expired_at) VALUES ($1, $2, $3)
			RETURNING id, created_at, updated_at`

	err := r.db.QueryRow(ctx, query, refreshToken.UserID, refreshToken.RefreshToken, refreshToken.ExpiredAt).Scan(&refreshToken.ID, &refreshToken.CreatedAt, &refreshToken.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to store refresh token: %w", err)
	}

	return nil

}
