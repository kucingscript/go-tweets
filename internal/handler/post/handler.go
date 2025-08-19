package post

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kucingscript/go-tweets/internal/config"
	"github.com/kucingscript/go-tweets/internal/middleware"
	"github.com/kucingscript/go-tweets/internal/service/post"
)

type Handler struct {
	validate    *validator.Validate
	postService post.PostService
	cfg         *config.Config
}

func NewPostHandler(validate *validator.Validate, postService post.PostService, cfg *config.Config) *Handler {
	return &Handler{
		validate:    validate,
		postService: postService,
		cfg:         cfg,
	}
}

func (h *Handler) RouteList(r *gin.RouterGroup) {
	postRoute := r.Group("/posts")
	postRoute.Use(middleware.AuthMiddleware(h.cfg))
	{
		postRoute.POST("", h.CreatePost)
		postRoute.GET("/:id", h.UpdatePost)
		postRoute.DELETE("/:id", h.DeletePost)
	}
}
