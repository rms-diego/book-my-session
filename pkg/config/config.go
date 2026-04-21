package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT string
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

	default:
		Env = &Config{
			PORT: os.Getenv("PORT"),
		}
	}
	return nil
}
