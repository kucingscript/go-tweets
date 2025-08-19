package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProfile(c *gin.Context) {
	userID, userIDExists := c.Get("userID")
	userEmail, userEmailExists := c.Get("userEmail")

	if !userIDExists || !userEmailExists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not logged in"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "welcome to your profile",
		"id":      userID,
		"email":   userEmail,
	})
}
