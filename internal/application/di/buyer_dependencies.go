package di

import (
	"ProyectoFinal/internal/handler"
	repository "ProyectoFinal/internal/repository/buyer"
	service "ProyectoFinal/internal/service/buyer"
	"database/sql"
)

func GetBuyerHandler(db *sql.DB) *handler.BuyerHandler {
	buyerRepository := repository.NewBuyerMySqlRepository(db)
	buyerService := service.NewBuyerService(buyerRepository)
	return handler.NewBuyerHandler(buyerService)
}
