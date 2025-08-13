package user

import (
	"context"
	"errors"
	"net/http"
)

func (s *userService) VerifyEmail(ctx context.Context, token string) (int, error) {
	user, err := s.userRepository.GetUserByVerificationToken(ctx, token)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if user == nil {
		return http.StatusNotFound, errors.New("invalid or expired verification token")
	}

	if user.IsVerified {
		return http.StatusBadRequest, errors.New("email has already been verified")
	}

	err = s.userRepository.VerifyUser(ctx, user.ID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
