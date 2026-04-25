package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT          string
	DB_HOST       string
	DB_PORT       string
	DB_USER       string
	DB_PASSWORD   string
	DB_NAME       string
	JWT_SECRET    string
	COOKIE_DOMAIN string
	S3_BUCKET     string
}

var Env *Config

func Init() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	switch {
	case os.Getenv("PORT") == "":
		return fmt.Errorf("PORT environment variable is required")

	case os.Getenv("DB_HOST") == "":
		return fmt.Errorf("DB_HOST environment variable is required")

	case os.Getenv("DB_PORT") == "":
		return fmt.Errorf("DB_PORT environment variable is required")

	case os.Getenv("DB_USER") == "":
		return fmt.Errorf("DB_USER environment variable is required")

	case os.Getenv("DB_PASSWORD") == "":
		return fmt.Errorf("DB_PASSWORD environment variable is required")

	case os.Getenv("DB_NAME") == "":
		return fmt.Errorf("DB_NAME environment variable is required")

	case os.Getenv("JWT_SECRET") == "":
		return fmt.Errorf("JWT_SECRET environment variable is required")

	case os.Getenv("COOKIE_DOMAIN") == "":
		return fmt.Errorf("COOKIE_DOMAIN environment variable is required")

	case os.Getenv("S3_BUCKET") == "":
		return fmt.Errorf("S3_BUCKET environment variable is required")

	default:
		Env = &Config{
			PORT:          os.Getenv("PORT"),
			DB_HOST:       os.Getenv("DB_HOST"),
			DB_PORT:       os.Getenv("DB_PORT"),
			DB_USER:       os.Getenv("DB_USER"),
			DB_PASSWORD:   os.Getenv("DB_PASSWORD"),
			DB_NAME:       os.Getenv("DB_NAME"),
			JWT_SECRET:    os.Getenv("JWT_SECRET"),
			COOKIE_DOMAIN: os.Getenv("COOKIE_DOMAIN"),
			S3_BUCKET:     os.Getenv("S3_BUCKET"),
		}
	}
	return nil
}
