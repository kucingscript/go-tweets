package user

import (
	"net/http"
	"time"

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
