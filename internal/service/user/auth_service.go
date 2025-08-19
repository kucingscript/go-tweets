package user

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/kucingscript/go-tweets/internal/dto"
	"github.com/kucingscript/go-tweets/internal/model"
	"github.com/kucingscript/go-tweets/pkg/jwt"
	"github.com/kucingscript/go-tweets/pkg/refreshtoken"
	"golang.org/x/crypto/bcrypt"
)

func (s *userService) Login(ctx context.Context, req *dto.LoginRequest) (string, string, int, error) {
	userExist, err := s.userRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if userExist == nil {
		return "", "", http.StatusUnauthorized, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(req.Password))
	if err != nil {
		return "", "", http.StatusUnauthorized, errors.New("invalid credentials")
	}

	if !userExist.IsVerified {
		return "", "", http.StatusUnauthorized, errors.New("please verify your email first")
	}

	token, err := jwt.CreateToken(userExist.ID, userExist.Email, s.cfg.JWTSecret)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	now := time.Now()
	refreshTokenExist, err := s.userRepository.GetRefreshToken(ctx, userExist.ID, now)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if refreshTokenExist != nil {
		return token, refreshTokenExist.RefreshToken, http.StatusOK, nil
	}

	refreshToken, err := refreshtoken.GenerateRefreshToken()
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	err = s.userRepository.StoreRefreshToken(ctx, &model.RefreshTokenModel{
		UserID:       userExist.ID,
		RefreshToken: refreshToken,
		ExpiredAt:    time.Now().Add(24 * time.Hour),
	})

	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	return token, refreshToken, http.StatusOK, nil
}

func (s *userService) Logout(ctx context.Context, refreshToken string) (int, error) {
	err := s.userRepository.DeleteRefreshToken(ctx, refreshToken)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

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
