package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kucingscript/go-tweets/internal/config"
	"github.com/kucingscript/go-tweets/internal/middleware"
	"github.com/kucingscript/go-tweets/internal/service/user"
)

type Handler struct {
	gin         *gin.Engine
	validate    *validator.Validate
	userService user.UserService
	cfg         *config.Config
}

func NewUserHandler(gin *gin.Engine, validate *validator.Validate, userService user.UserService, cfg *config.Config) *Handler {
	return &Handler{
		gin:         gin,
		validate:    validate,
		userService: userService,
		cfg:         cfg,
	}
}

func (h *Handler) RouteList() {
	api := h.gin.Group("/api")
	v1 := api.Group("/v1")

	authRoute := v1.Group("/auth")
	{
		authRoute.POST("/register", h.Register)
		authRoute.GET("/verify-email", h.VerifyEmail)

		authRoute.POST("/login", h.Login)
		authRoute.POST("/logout", h.Logout)
		authRoute.POST("/refresh-token", h.RefreshToken)

		authRoute.POST("/forgot-password", h.ForgotPassword)
		authRoute.POST("/reset-password", h.ResetPassword)
	}

	userRoute := v1.Group("/user")
	userRoute.Use(middleware.AuthMiddleware(h.cfg))
	{
		userRoute.GET("/profile", h.GetProfile)
	}
}
