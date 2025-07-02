package di

import (
	"ProyectoFinal/internal/handler"
	repository "ProyectoFinal/internal/repository/seller"
	service "ProyectoFinal/internal/service/seller"
	"ProyectoFinal/pkg/models"
)

func GetSellerHandler(db map[int]models.Seller) *handler.SellerHandler {
	sellerRepository := repository.NewSellerRepository(db)
	sellerService := service.NewSellerService(sellerRepository)
	sellerHandler := handler.NewSellerHandler(sellerService)
	return sellerHandler
}
