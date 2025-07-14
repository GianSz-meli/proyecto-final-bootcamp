package di

import (
	handler "ProyectoFinal/internal/handler"
	repository "ProyectoFinal/internal/repository/products"
	service "ProyectoFinal/internal/service/products"
	//"ProyectoFinal/pkg/models"
	"database/sql"
)

// func GetProductsHandler(db map[int]models.Product) *handler.ProductHandler {
// 	productRepository := repository.NewProductMap(db)
// 	productService := service.NewProductDefault(productRepository)
// 	productHandler := handler.NewProductHandler(productService)
// 	return productHandler
// }

func GetProductsHandler(db *sql.DB) *handler.ProductHandler {
	productRepository := repository.NewProductSQL(db)
	productService := service.NewProductDefault(productRepository)
	productHandler := handler.NewProductHandler(productService)
	return productHandler
}
