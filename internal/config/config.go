package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT           string
	DBUrlMigration string
	JWTSecret      string

	AllowedOrigins string

	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string

	SMTPHost   string
	SMTPPort   int
	SMTPUser   string
	SMTPPass   string
	SMTPSender string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		log.Println("Warning: Could not load .env file. Falling back to environment variables.")
	} else {
		log.Println("Config loaded from .env file")
	}

	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return nil, fmt.Errorf("invalid SMTP_PORT: %w", err)
	}

	log.Println("Config loaded")

	return &Config{
		PORT:           os.Getenv("APP_PORT"),
		DBUrlMigration: os.Getenv("DATABASE_URL"),
		JWTSecret:      os.Getenv("JWT_SECRET"),

		AllowedOrigins: os.Getenv("ALLOWED_ORIGINS"),

		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBUser: os.Getenv("POSTGRES_USER"),
		DBPass: os.Getenv("POSTGRES_PASSWORD"),
		DBName: os.Getenv("POSTGRES_DB"),

		SMTPHost:   os.Getenv("SMTP_HOST"),
		SMTPPort:   smtpPort,
		SMTPUser:   os.Getenv("SMTP_USER"),
		SMTPPass:   os.Getenv("SMTP_PASS"),
		SMTPSender: os.Getenv("SMTP_SENDER"),
	}, nil
}
