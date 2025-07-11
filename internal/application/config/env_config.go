package config

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadDotEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}
