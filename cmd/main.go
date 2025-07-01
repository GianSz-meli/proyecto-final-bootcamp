package main

import (
	"ProyectoFinal/internal/application"
	"ProyectoFinal/internal/application/loader"
	"fmt"
)

func main() {
	cfg := &application.ConfigServerChi{
		ServerAddress: ":8080",
		LoaderFilePath: map[string]string{
			loader.Seller:    "docs/db/sellers.json",
			loader.Section:   "docs/db/sections_10.json",
			loader.Employee:  "docs/db/employees.json",
			loader.Warehouse: "docs/db/warehouses.json",
		},
	}
	app := application.NewServerChi(cfg)

	fmt.Printf("Server started in http://localhost%s \n", cfg.ServerAddress)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
