package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kucingscript/go-tweets/internal/dto"
)

func (h *Handler) Login(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req dto.LoginRequest
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

	token, refreshToken, statusCode, err := h.userService.Login(ctx, &req)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	h.setTokenCookies(c, token, refreshToken)

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
