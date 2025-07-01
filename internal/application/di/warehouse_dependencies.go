package di

import (
	"ProyectoFinal/pkg/models"
	repository "ProyectoFinal/internal/repository/warehouse"
	service "ProyectoFinal/internal/service/warehouse"
	handler "ProyectoFinal/internal/handler"

)

func GetWarehouseHandler(db map[int]models.Warehouse) *handler.WarehouseHandler {
	warehouseRepository := repository.NewMemoryWarehouseRepository(db)
	warehouseService := service.NewWarehouseService(warehouseRepository)
	warehouseHandler := handler.NewWarehouseHandler(warehouseService)
	return warehouseHandler
}