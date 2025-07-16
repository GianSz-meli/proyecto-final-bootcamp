package di

import (
	handler "ProyectoFinal/internal/handler"
	repository "ProyectoFinal/internal/repository/products"
	service "ProyectoFinal/internal/service/products"
	"database/sql"
)

func GetProductsHandler(sqlDB *sql.DB) *handler.ProductHandler {
	productRepository := repository.NewProductSQL(sqlDB)
	productService := service.NewProductDefault(productRepository)
	productHandler := handler.NewProductHandler(productService)
	return productHandler
}
