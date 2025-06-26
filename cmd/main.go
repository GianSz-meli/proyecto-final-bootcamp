package main

import (
	"ProyectoFinal/internal/application"
	"fmt"
)

func main() {
	cfg := &application.ConfigServerChi{
		ServerAddress:  ":8080",
		LoaderFilePath: "",
	}
	app := application.NewServerChi(cfg)

	fmt.Printf("Server started in http://localhost%s \n", cfg.ServerAddress)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
