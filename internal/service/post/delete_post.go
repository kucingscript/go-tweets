package post

import (
	"context"
	"errors"
	"net/http"
)

func (s *postService) DeletePost(ctx context.Context, postID, userID int64) (int, error) {
	postExist, err := s.postRepository.GetPostByID(ctx, postID)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if postExist == nil {
		return http.StatusNotFound, errors.New("post not found")
	}

	if postExist.UserID != userID {
		return http.StatusUnauthorized, errors.New("user is not authorized to delete this post")
	}

	err = s.postRepository.SoftDeletePost(ctx, postID)
	if err != nil {
		if err.Error() == "post not found or already deleted" {
			return http.StatusNotFound, err
		}

		return http.StatusInternalServerError, err
	}

	return http.StatusNoContent, nil
}
