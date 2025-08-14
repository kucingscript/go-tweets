package middleware

import (
	"log"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kucingscript/go-tweets/internal/config"
)

func CORSMiddleware(cfg *config.Config) gin.HandlerFunc {
	allowedOrigins := strings.Split(cfg.AllowedOrigins, ",")
	if len(allowedOrigins) == 0 || cfg.AllowedOrigins == "" {
		log.Println("Peringatan: Variabel lingkungan ALLOWED_ORIGINS tidak diatur. CORS mungkin tidak akan berfungsi.")
	}

	return cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
