package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Logout(c *gin.Context) {
	refreshToken, err := c.Cookie(refreshTokenCookieName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not logged in"})
		return
	}

	statusCode, err := h.userService.Logout(c.Request.Context(), refreshToken)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	h.clearTokenCookies(c)

	c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}
