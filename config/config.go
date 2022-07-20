package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Config(key string) string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Can not Load .env File.")
	}

	return os.Getenv(key)
}
