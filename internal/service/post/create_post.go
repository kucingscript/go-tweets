package post

import (
	"context"
	"net/http"

	"github.com/kucingscript/go-tweets/internal/dto"
	"github.com/kucingscript/go-tweets/internal/model"
)

func (s *postService) CreatePost(ctx context.Context, req *dto.CreateOrUpdatePostRequest, userID int64) (*model.PostModel, int, error) {
	sanitizedTitle := s.htmlSanitizer.Sanitize(req.Title)
	sanitizedContent := s.htmlSanitizer.Sanitize(req.Content)

	post := &model.PostModel{
		UserID:  userID,
		Title:   sanitizedTitle,
		Content: sanitizedContent,
	}

	createdPost, err := s.postRepository.StorePost(ctx, post)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return createdPost, http.StatusCreated, nil
}
