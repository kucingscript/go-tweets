package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/api", "", true, true)
	c.SetCookie("refresh_token", "", -1, "/api/v1/auth", "", true, true)

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}
