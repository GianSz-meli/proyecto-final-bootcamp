package di

import (
	handler "ProyectoFinal/internal/handler"
	repository "ProyectoFinal/internal/repository/products"
	service "ProyectoFinal/internal/service/products"
	"ProyectoFinal/pkg/models"
)

func GetProductsHandler(db map[int]models.Product) *handler.ProductHandler {
	productRepository := repository.NewProductMap(db)
	productService := service.NewProductDefault(productRepository)
	productHandler := handler.NewProductHandler(productService)
	return productHandler
}
