package di

import (
	products "ProyectoFinal/internal/handler/products"
	repository "ProyectoFinal/internal/repository/products"
	service "ProyectoFinal/internal/service/products"

	"database/sql"
)

func GetProductsHandler(sqlDB *sql.DB) *products.ProductHandler {
	productRepository := repository.NewProductSQL(sqlDB)
	productService := service.NewProductDefault(productRepository)
	productHandler := products.NewProductHandler(productService)
	return productHandler
}
