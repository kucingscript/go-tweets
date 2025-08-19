package user

import (
	"github.com/gin-gonic/gin"
	"github.com/kucingscript/go-tweets/internal/middleware"
)

func (h *Handler) RouteList(r *gin.RouterGroup) {
	authRoute := r.Group("/auth")
	{
		authRoute.POST("/login", h.Login)
		authRoute.POST("/refresh-token", h.RefreshToken)
		authRoute.POST("/logout", h.Logout)
	}

	accountRoute := r.Group("/account")
	{
		accountRoute.POST("/register", h.Register)
		accountRoute.GET("/verify-email", h.VerifyEmail)
	}

	passwordRoute := r.Group("/password")
	{
		passwordRoute.POST("/forgot", h.ForgotPassword)
		passwordRoute.POST("/reset", h.ResetPassword)
	}

	userRoute := r.Group("/users")
	userRoute.Use(middleware.AuthMiddleware(h.cfg))
	{
		userRoute.GET("/profile", h.GetProfile)
	}
}
