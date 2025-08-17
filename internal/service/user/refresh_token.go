package user

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/kucingscript/go-tweets/pkg/jwt"
)

func (s *userService) RefreshToken(ctx context.Context, refreshToken string) (string, int, error) {
	user, err := s.userRepository.GetUserByRefreshToken(ctx, refreshToken)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	if user == nil {
		return "", http.StatusUnauthorized, errors.New("invalid or expired refresh token")
	}

	newAccessToken, err := jwt.CreateToken(user.ID, user.Email, s.cfg.JWTSecret)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return newAccessToken, http.StatusOK, nil
}

func (s *userService) CleanUpExpiredTokens(ctx context.Context) {
	deletedCount, err := s.userRepository.DeleteExpiredRefreshTokens(ctx)
	if err != nil {
		log.Printf("ERROR: failed to clean up expired refresh tokens: %v", err)
		return
	}

	if deletedCount > 0 {
		log.Printf("Successfully cleaned up %d expired refresh tokens", deletedCount)
	}
}
