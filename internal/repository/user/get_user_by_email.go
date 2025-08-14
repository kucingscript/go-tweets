package user

import (
	"context"
	"errors"

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
