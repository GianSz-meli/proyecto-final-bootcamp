package config

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadDotEnv loads environment variables from a .env file.
// The application will terminate if the .env file cannot be loaded.
func LoadDotEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}
