package post

import (
	"context"
	"errors"
	"net/http"

	"github.com/kucingscript/go-tweets/internal/dto"
	"github.com/kucingscript/go-tweets/internal/model"
)

func (s *postService) UpdatePost(ctx context.Context, req *dto.CreateOrUpdatePostRequest, postID, userID int64) (*model.PostModel, int, error) {
	postExist, err := s.postRepository.GetPostByID(ctx, postID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if postExist == nil {
		return nil, http.StatusNotFound, errors.New("post not found")
	}

	if postExist.UserID != userID {
		return nil, http.StatusUnauthorized, errors.New("user is not authorized to update this post")
	}

	sanitizedTitle := s.htmlSanitizer.Sanitize(req.Title)
	sanitizedContent := s.htmlSanitizer.Sanitize(req.Content)

	postExist.Title = sanitizedTitle
	postExist.Content = sanitizedContent

	updatedPost, err := s.postRepository.UpdatePost(ctx, postExist)
	if err != nil {
		if err.Error() == "post not found or already deleted" {
			return nil, http.StatusNotFound, err
		}

		return nil, http.StatusInternalServerError, err
	}

	return updatedPost, http.StatusOK, nil
}
