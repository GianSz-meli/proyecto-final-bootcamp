package di

import (
	"ProyectoFinal/internal/handler"
	repository "ProyectoFinal/internal/repository/seller"
	service "ProyectoFinal/internal/service/seller"
	"database/sql"
)

func GetSellerHandler(db *sql.DB) *handler.SellerHandler {
	sellerRepository := repository.NewSellerMysqlRepository(db)
	sellerService := service.NewSellerService(sellerRepository)
	sellerHandler := handler.NewSellerHandler(sellerService)
	return sellerHandler
}
