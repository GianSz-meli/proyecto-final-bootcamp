package main

import (
	"ProyectoFinal/internal/application"
	"ProyectoFinal/internal/application/config"
	"fmt"
)

func main() {
	cfg := &application.ConfigServerChi{
		ServerAddress: ":8080",
	}

	config.LoadDotEnv()

	app := application.NewServerChi(cfg)

	fmt.Printf("Server started in http://localhost%s \n", cfg.ServerAddress)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
