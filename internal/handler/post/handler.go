package post

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kucingscript/go-tweets/internal/service/post"
)

type Handler struct {
	gin         *gin.Engine
	validate    *validator.Validate
	postService post.PostService
}

func NewPostHandler(gin *gin.Engine, validate *validator.Validate, postService post.PostService) *Handler {
	return &Handler{
		gin:         gin,
		validate:    validate,
		postService: postService,
	}
}

func (h *Handler) RouteList() {

}
