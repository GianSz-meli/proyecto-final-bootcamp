package di

import (
	handler "ProyectoFinal/internal/handler"
	repositorySection "ProyectoFinal/internal/repository/section"
	repository "ProyectoFinal/internal/repository/warehouse"
	serviceSection "ProyectoFinal/internal/service/section"
	service "ProyectoFinal/internal/service/warehouse"
	"ProyectoFinal/pkg/models"
)

func GetWarehouseHandler(db map[int]models.Warehouse) *handler.WarehouseHandler {
	warehouseRepository := repository.NewMemoryWarehouseRepository(db)
	warehouseService := service.NewWarehouseService(warehouseRepository)
	warehouseHandler := handler.NewWarehouseHandler(warehouseService)
	return warehouseHandler
}

func GetSectionHandler(db map[int]models.Section) *handler.SectionDefault {
	sectionRepository := repositorySection.NewSectionMap(db)
	sectionService := serviceSection.NewSectionDefault(sectionRepository)
	sectionHandler := handler.NewSectionDefault(sectionService)
	return sectionHandler
}
