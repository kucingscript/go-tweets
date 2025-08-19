package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kucingscript/go-tweets/internal/config"
	postHandler "github.com/kucingscript/go-tweets/internal/handler/post"
	userHandler "github.com/kucingscript/go-tweets/internal/handler/user"
	"github.com/kucingscript/go-tweets/internal/mailer"
	"github.com/kucingscript/go-tweets/internal/middleware"
	postRepository "github.com/kucingscript/go-tweets/internal/repository/post"
	userRepository "github.com/kucingscript/go-tweets/internal/repository/user"
	postService "github.com/kucingscript/go-tweets/internal/service/post"
	userService "github.com/kucingscript/go-tweets/internal/service/user"
	"github.com/kucingscript/go-tweets/pkg/postgres"
	"github.com/robfig/cron/v3"
)

func main() {

	// gin.SetMode(gin.ReleaseMode) Enable for production

	r := gin.Default()
	validate := validator.New()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.ConnectPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	r.Use(middleware.CORSMiddleware(cfg))
	r.Use(gin.Logger(), gin.Recovery())

	mailer := mailer.NewMailer(cfg)

	userRepository := userRepository.NewUserRepository(db)
	postRepository := postRepository.NewPostRepository(db)

	userService := userService.NewUserService(cfg, userRepository, mailer)
	postService := postService.NewPostService(cfg, postRepository)

	v1 := r.Group("/api/v1")

	userHandler := userHandler.NewUserHandler(validate, userService, cfg)
	postHandler := postHandler.NewPostHandler(validate, postService, cfg)

	userHandler.RouteList(v1)
	postHandler.RouteList(v1)

	c := cron.New()
	c.AddFunc("@daily", func() {
		log.Println("Running scheduled task to clean up expired tokens...")
		userService.CleanUpExpiredTokens(context.Background())
	})
	c.Start()

	server := fmt.Sprintf(":%s", cfg.PORT)
	r.Run(server)
}
