package di

import (
	handler "ProyectoFinal/internal/handler"
	localityRepository "ProyectoFinal/internal/repository/locality"
	repository "ProyectoFinal/internal/repository/warehouse"
	service "ProyectoFinal/internal/service/warehouse"
	"database/sql"
)

func GetWarehouseHandler(db *sql.DB) *handler.WarehouseHandler {
	warehouseRepository := repository.NewSqlWarehouseRepository(db)
	localityRepo := localityRepository.NewSqlLocalityRepository(db)
	warehouseService := service.NewWarehouseService(warehouseRepository, localityRepo)
	warehouseHandler := handler.NewWarehouseHandler(warehouseService)
	return warehouseHandler
}
