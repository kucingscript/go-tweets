package user

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/kucingscript/go-tweets/internal/model"
)

func (r *userRepository) GetUserByEmailOrUsername(ctx context.Context, email, username string) (*model.UserModel, error) {
	query := `SELECT id, username, email, password, created_at, updated_at
			FROM users
			WHERE email = $1
			OR username = $2`

	row := r.db.QueryRow(ctx, query, email, username)

	var user model.UserModel
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
