package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kucingscript/go-tweets/internal/dto"
)

func (h *Handler) Register(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req dto.RegisterRequest
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

	createdUser, statusCode, err := h.userService.Register(ctx, &req)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	response := dto.RegisterResponse{
		ID:         createdUser.ID,
		Email:      createdUser.Email,
		Username:   createdUser.Username,
		IsVerified: createdUser.IsVerified,
		CreatedAt:  createdUser.CreatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *Handler) VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "verification token is required"})
		return
	}

	statusCode, err := h.userService.VerifyEmail(c.Request.Context(), token)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(statusCode, gin.H{"message": "Email verified successfully"})
}
