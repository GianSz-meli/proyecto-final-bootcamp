package di

import (
	"ProyectoFinal/internal/handler"
	carrierRepo "ProyectoFinal/internal/repository/carrier"
	carrierService "ProyectoFinal/internal/service/carrier"
	"database/sql"
)

func GetCarrierHandler(sqlDB *sql.DB) *handler.CarrierHandler {
	carrierRepository := carrierRepo.NewSqlCarrierRepository(sqlDB)
	carrierServiceImpl := carrierService.NewCarrierService(carrierRepository)
	return handler.NewCarrierHandler(carrierServiceImpl)
}
