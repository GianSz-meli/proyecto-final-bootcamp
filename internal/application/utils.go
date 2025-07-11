package application

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func LoadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}
}
