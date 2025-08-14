package user

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/kucingscript/go-tweets/internal/dto"
	"github.com/kucingscript/go-tweets/internal/model"
	"github.com/kucingscript/go-tweets/pkg/jwt"
	"github.com/kucingscript/go-tweets/pkg/refreshtoken"
	"golang.org/x/crypto/bcrypt"
)

func (s *userService) Login(ctx context.Context, req *dto.LoginRequest) (string, string, int, error) {
	// check user exist
	userExist, err := s.userRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if userExist == nil {
		return "", "", http.StatusNotFound, errors.New("wrong email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(req.Password))
	if err != nil {
		return "", "", http.StatusNotFound, errors.New("wrong email or password")
	}

	// generate access token
	token, err := jwt.CreateToken(userExist.ID, userExist.Email, s.cfg.JWTSecret)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	// get refresh token if exist
	now := time.Now()
	refreshTokenExist, err := s.userRepository.GetRefreshToken(ctx, userExist.ID, now)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if refreshTokenExist != nil {
		return token, refreshTokenExist.RefreshToken, http.StatusOK, nil
	}

	// generate & store refresh token
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
