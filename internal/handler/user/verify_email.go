package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
