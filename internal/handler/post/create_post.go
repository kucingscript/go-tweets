package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kucingscript/go-tweets/internal/dto"
)

func (h *Handler) CreatePost(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req dto.CreateOrUpdatePostRequest
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := h.validate.Struct(req); err != nil {
		errorMessages := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			errorMessages[e.Field()] = "Error: " + e.Tag()
		}

		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	userID, ok := c.Get("userID")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	createdPost, statusCode, err := h.postService.CreatePost(ctx, &req, userID.(int64))
	if err != nil {
		c.AbortWithStatusJSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	response := dto.CreateOrUpdatePostResponse{
		ID:        createdPost.ID,
		Title:     createdPost.Title,
		Content:   createdPost.Content,
		CreatedAt: createdPost.CreatedAt,
		UpdatedAt: createdPost.UpdatedAt,
	}

	c.JSON(statusCode, response)
}
