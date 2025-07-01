package di

import (
	handler "ProyectoFinal/internal/handler"
	repository "ProyectoFinal/internal/repository/warehouse"
	service "ProyectoFinal/internal/service/warehouse"
	"ProyectoFinal/pkg/models"
)

func GetWarehouseHandler(db map[int]models.Warehouse) *handler.WarehouseHandler {
	warehouseRepository := repository.NewMemoryWarehouseRepository(db)
	warehouseService := service.NewWarehouseService(warehouseRepository)
	warehouseHandler := handler.NewWarehouseHandler(warehouseService)
	return warehouseHandler
}
