package user

import (
	"context"
	"net/http"
)

func (s *userService) Logout(ctx context.Context, refreshToken string) (int, error) {
	err := s.userRepository.DeleteRefreshToken(ctx, refreshToken)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
