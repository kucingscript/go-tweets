package post

import (
	"github.com/gin-gonic/gin"
	"github.com/kucingscript/go-tweets/internal/middleware"
)

func (h *Handler) RouteList(r *gin.RouterGroup) {
	postRoute := r.Group("/posts")
	postRoute.Use(middleware.AuthMiddleware(h.cfg))
	{
		postRoute.POST("", h.CreatePost)
		postRoute.GET("/:id", h.UpdatePost)
		postRoute.DELETE("/:id", h.DeletePost)
	}
}
