package dto

type (
	CreatePostRequest struct {
		Title   string `json:"title" validate:"required,min=3,max=100"`
		Content string `json:"content" validate:"required,min=3"`
	}

	CreatePostResponse struct {
		ID int64 `json:"id"`
	}
)
