package user

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie(refreshTokenCookieName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "refresh token not found"})
		return
	}

	newAccessToken, statusCode, err := h.userService.RefreshToken(c.Request.Context(), refreshToken)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	accessCookie := &http.Cookie{
		Name:     accessTokenCookieName,
		Value:    newAccessToken,
		Expires:  time.Now().Add(15 * time.Minute),
		Path:     "/api",
		Domain:   "",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(c.Writer, accessCookie)

	c.JSON(http.StatusOK, gin.H{"message": "token refreshed successfully"})
}
