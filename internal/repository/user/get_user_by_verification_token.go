package user

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/kucingscript/go-tweets/internal/model"
)

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
