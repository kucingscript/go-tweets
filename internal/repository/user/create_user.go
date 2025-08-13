package user

import (
	"context"
	"fmt"

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
