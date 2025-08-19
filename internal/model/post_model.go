package model

import "time"

type (
	PostModel struct {
		ID        int64      `json:"id"`
		UserID    int64      `json:"user_id"`
		Title     string     `json:"title"`
		Content   string     `json:"content"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}
)
