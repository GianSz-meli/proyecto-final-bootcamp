package warehouse

import (
	"ProyectoFinal/internal/repository/warehouse"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"fmt"
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

func (s *WarehouseServiceImpl) GetWarehouseById(id int) (*models.Warehouse, error) {
	warehouse := s.warehouseRepo.GetById(id)
	if warehouse == nil {
		return nil, fmt.Errorf("%w: The warehouse with id '%d' does not exist", errors.ErrNotFound, id)
	}
	return warehouse, nil
}

func (s *WarehouseServiceImpl) CreateWarehouse(warehouse models.Warehouse) (*models.Warehouse, error) {
	if s.warehouseRepo.ExistsByCode(warehouse.WarehouseCode) {
		return nil, fmt.Errorf("%w: warehouse with code '%s' already exists", errors.ErrAlreadyExists, warehouse.WarehouseCode)
	}

	createdWarehouse := s.warehouseRepo.Create(warehouse)
	return createdWarehouse, nil
}
