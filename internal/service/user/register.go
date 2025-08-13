package user

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/kucingscript/go-tweets/internal/dto"
	"github.com/kucingscript/go-tweets/internal/model"
	"github.com/kucingscript/go-tweets/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func (s *userService) Register(ctx context.Context, req *dto.RegisterRequest) (*model.UserModel, int, error) {
	// check user exist
	userExist, err := s.userRepository.GetUserByEmailOrUsername(ctx, req.Email, req.Username)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if userExist != nil {
		return nil, http.StatusBadRequest, errors.New("user already exist")
	}

	// hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// create user
	user := &model.UserModel{
		Email:    req.Email,
		Username: req.Username,
		Password: string(passwordHash),
	}

	err = s.userRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// Verifikasi Email
	// Token Verifikasi Email
	token, err := utils.GenerateSecureToken(32)
	if err != nil {
		log.Printf("CRITICAL: could not generate verification token for user %d: %v", user.ID, err)
	}

	// Save token to database
	err = s.userRepository.UpdateVerificationToken(ctx, user.ID, token)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// prepare email
	emailData := map[string]interface{}{
		"Username":         user.Username,
		"VerificationLink": fmt.Sprintf("http://localhost:%s/api/v1/auth/verify-email?token=%s", s.cfg.PORT, token),
	}

	// Send verification email asynchronously
	go func() {
		err := s.mailer.Send(user.Email, "user_welcome.tmpl", emailData)
		if err != nil {
			log.Printf("ERROR: could not send verification email to %s: %s", user.Email, err)
		}
	}()

	return user, http.StatusCreated, nil
}
