package di

import (
	handler "ProyectoFinal/internal/handler/warehouse"
	repository "ProyectoFinal/internal/repository/warehouse"
	service "ProyectoFinal/internal/service/warehouse"
	"database/sql"
)

func GetWarehouseHandler(db *sql.DB) *handler.WarehouseHandler {
	warehouseRepository := repository.NewSqlWarehouseRepository(db)
	warehouseService := service.NewWarehouseService(warehouseRepository)
	warehouseHandler := handler.NewWarehouseHandler(warehouseService)
	return warehouseHandler
}
