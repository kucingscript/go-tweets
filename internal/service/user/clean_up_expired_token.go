package user

import (
	"context"
	"log"
)

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
