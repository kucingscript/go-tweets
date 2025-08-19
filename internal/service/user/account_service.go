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
	userExist, err := s.userRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if userExist != nil {
		go func() {
			err := s.mailer.Send(userExist.Email, "user_already_exists.tmpl", userExist)
			if err != nil {
				log.Printf("ERROR: could not send duplicate registration email to %s: %s", userExist.Email, err)
			}
		}()

		return userExist, http.StatusCreated, nil
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	sanitizedUsername := s.htmlSanitizer.Sanitize(req.Username)
	user := &model.UserModel{
		Email:    req.Email,
		Username: sanitizedUsername,
		Password: string(passwordHash),
	}

	err = s.userRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	token, err := utils.GenerateSecureToken(32)
	if err != nil {
		log.Printf("CRITICAL: could not generate verification token for user %d: %v", user.ID, err)
	}

	err = s.userRepository.UpdateVerificationToken(ctx, user.ID, token)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	go func() {
		emailData := map[string]interface{}{
			"Username":         user.Username,
			"VerificationLink": fmt.Sprintf("http://localhost:%s/api/v1/auth/verify-email?token=%s", s.cfg.PORT, token),
		}
		err := s.mailer.Send(user.Email, "user_welcome.tmpl", emailData)
		if err != nil {
			log.Printf("ERROR: could not send verification email to %s: %s", user.Email, err)
		}
	}()

	return user, http.StatusCreated, nil
}

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
