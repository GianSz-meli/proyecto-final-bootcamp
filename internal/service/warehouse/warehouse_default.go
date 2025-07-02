package warehouse

import (
	"ProyectoFinal/internal/repository/warehouse"
	"ProyectoFinal/pkg/models"
)

type WarehouseServiceImpl struct {
	warehouseRepo warehouse.WarehouseRepository
}

func NewWarehouseService(warehouseRepo warehouse.WarehouseRepository) WarehouseService {
	return &WarehouseServiceImpl{
		warehouseRepo: warehouseRepo,
	}
}

func (s *WarehouseServiceImpl) GetAllWarehouses() []models.Warehouse {
	warehouses := s.warehouseRepo.GetAll()
	return warehouses
}
