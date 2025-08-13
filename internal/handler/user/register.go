package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kucingscript/go-tweets/internal/dto"
)

func (h *Handler) Register(c *gin.Context) {
	var (
		ctx = c.Request.Context()
		req dto.RegisterRequest
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validate.Struct(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	userID, statusCode, err := h.userService.Register(ctx, &req)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, err.Error())
		return
	}

	c.JSON(http.StatusCreated, dto.RegisterResponse{ID: userID})
}
