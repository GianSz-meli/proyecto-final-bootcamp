package di

import (
	"ProyectoFinal/pkg/models"
	repository "ProyectoFinal/internal/repository/products"
	service "ProyectoFinal/internal/service/products"
	handler "ProyectoFinal/internal/handler"

)

func GetProductsHandler(db map[int]models.Product) *handler.ProductHandler {
	productRepository := repository.NewProductMap(db)
	productService := service.NewProductDefault(productRepository)
	productHandler := handler.NewProductHandler(productService)
	return productHandler
}