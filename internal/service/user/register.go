package user

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/kucingscript/go-tweets/internal/dto"
	"github.com/kucingscript/go-tweets/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func (s *userService) Register(ctx context.Context, req *dto.RegisterRequest) (int64, int, error) {
	// check user exist
	userExist, err := s.userRepository.GetUserByEmailOrUsername(ctx, req.Email, req.Username)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}

	if userExist != nil {
		return 0, http.StatusBadRequest, errors.New("user already exist")
	}

	// hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}

	// create user
	now := time.Now()
	user := &model.UserModel{
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(passwordHash),
		CreatedAt: now,
		UpdatedAt: now,
	}

	userID, err := s.userRepository.CreateUser(ctx, user)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}

	return userID, http.StatusCreated, nil
}
