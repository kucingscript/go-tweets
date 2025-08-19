package dto

import "time"

type (
	CreateOrUpdatePostRequest struct {
		Title   string `json:"title" validate:"required,min=3,max=100"`
		Content string `json:"content" validate:"required,min=3"`
	}

	CreateOrUpdatePostResponse struct {
		ID        int64     `json:"id"`
		UserID    int64     `json:"user_id"`
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
