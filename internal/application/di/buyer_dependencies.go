package di

import (
	"ProyectoFinal/internal/handler/buyer"
	repository "ProyectoFinal/internal/repository/buyer"
	service "ProyectoFinal/internal/service/buyer"
	"database/sql"
)

// GetBuyerHandler creates and configures a complete buyer handler with all its dependencies.
func GetBuyerHandler(db *sql.DB) *buyer.BuyerHandler {
	buyerRepository := repository.NewBuyerMySqlRepository(db)
	buyerService := service.NewBuyerService(buyerRepository)
	return buyer.NewBuyerHandler(buyerService)
}
