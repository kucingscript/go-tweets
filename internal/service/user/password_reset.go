// internal/service/user/password_reset.go
package user

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/kucingscript/go-tweets/internal/dto"
	"github.com/kucingscript/go-tweets/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

const resetTokenExpiry = 15 * time.Minute

func (s *userService) ForgotPassword(ctx context.Context, email string) (int, error) {
	user, err := s.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if user == nil {
		return http.StatusOK, nil
	}

	token, err := utils.GenerateSecureToken(32)
	if err != nil {
		log.Printf("CRITICAL: could not generate verification token for user %d: %v", user.ID, err)
	}

	expiresAt := time.Now().Add(resetTokenExpiry)
	err = s.userRepository.SetPasswordResetToken(ctx, user.ID, token, expiresAt)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	go func() {
		resetLink := fmt.Sprintf("http://localhost:8080/reset-password?token=%s", token)

		emailData := map[string]interface{}{
			"Username":  user.Username,
			"ResetLink": resetLink,
		}
		err := s.mailer.Send(user.Email, "password_reset.tmpl", emailData)
		if err != nil {
			log.Printf("ERROR: could not send password reset email to %s: %s", user.Email, err)
		}
	}()

	return http.StatusOK, nil
}

func (s *userService) ResetPassword(ctx context.Context, req *dto.ResetPasswordRequest) (int, error) {
	user, err := s.userRepository.GetUserByResetToken(ctx, req.Token)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if user == nil || user.PasswordResetTokenExpiresAt == nil || user.PasswordResetTokenExpiresAt.Before(time.Now()) {
		return http.StatusBadRequest, errors.New("invalid or expired reset token")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	err = s.userRepository.UpdatePassword(ctx, user.ID, string(password))
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
