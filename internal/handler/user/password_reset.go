package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kucingscript/go-tweets/internal/dto"
)

func (h *Handler) ForgotPassword(c *gin.Context) {
	var req dto.ForgotPasswordRequest
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

	statusCode, err := h.userService.ForgotPassword(c.Request.Context(), req.Email)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(statusCode, gin.H{"message": "if you have an account with this email, you will receive an email to reset your password."})
}

func (h *Handler) ResetPassword(c *gin.Context) {
	var req dto.ResetPasswordRequest
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

	statusCode, err := h.userService.ResetPassword(c.Request.Context(), &req)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(statusCode, gin.H{"message": "password reset successful"})
}
