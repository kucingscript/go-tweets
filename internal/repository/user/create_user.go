package user

import (
	"context"
	"fmt"

	"github.com/kucingscript/go-tweets/internal/model"
)

func (r *userRepository) CreateUser(ctx context.Context, user *model.UserModel) (int64, error) {
	query := `INSERT INTO users (email, username, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)
			RETURNING id;`

	var userID int64

	err := r.db.QueryRow(ctx, query, user.Email, user.Username, user.Password).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("unable to create user: %w", err)
	}

	return userID, nil
}
