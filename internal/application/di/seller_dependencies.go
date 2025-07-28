package di

import (
	"ProyectoFinal/internal/handler/seller"
	repository "ProyectoFinal/internal/repository/seller"
	service "ProyectoFinal/internal/service/seller"
	"database/sql"
)

// GetSellerHandler initializes and returns a SellerHandler with the provided database connection.
func GetSellerHandler(db *sql.DB) *seller.SellerHandler {
	sellerRepository := repository.NewSellerMysqlRepository(db)
	sellerService := service.NewSellerService(sellerRepository)
	sellerHandler := seller.NewSellerHandler(sellerService)
	return sellerHandler
}
