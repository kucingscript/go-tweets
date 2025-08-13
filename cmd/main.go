package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kucingscript/go-tweets/internal/config"
	userHandler "github.com/kucingscript/go-tweets/internal/handler/user"
	userRepository "github.com/kucingscript/go-tweets/internal/repository/user"
	userService "github.com/kucingscript/go-tweets/internal/service/user"
	"github.com/kucingscript/go-tweets/pkg/postgres"
)

func main() {
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

	r.Use(gin.Logger(), gin.Recovery())

	userRepository := userRepository.NewRepository(db)
	userService := userService.NewUserService(cfg, userRepository)
	userHandler := userHandler.NewHandler(r, validate, userService)
	userHandler.RouteList()

	server := fmt.Sprintf(":%s", cfg.PORT)
	r.Run(server)
}
