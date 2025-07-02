package warehouse

import (
	"ProyectoFinal/internal/repository/warehouse"
	"ProyectoFinal/pkg/errors"
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

func (s *WarehouseServiceImpl) GetWarehouseById(id int) (*models.Warehouse, error) {
	warehouse := s.warehouseRepo.GetById(id)
	if warehouse == nil {
		return nil, errors.WrapErrNotFound("warehouse", "id", id)
	}
	return warehouse, nil
}

func (s *WarehouseServiceImpl) CreateWarehouse(warehouse models.Warehouse) (*models.Warehouse, error) {
	if s.warehouseRepo.ExistsByCode(*warehouse.WarehouseCode) {
		return nil, errors.WrapErrAlreadyExist("warehouse", "code", *warehouse.WarehouseCode)
	}

	createdWarehouse := s.warehouseRepo.Create(warehouse)
	return createdWarehouse, nil
}

func (s *WarehouseServiceImpl) UpdateWarehouse(id int, warehouse models.Warehouse) (*models.Warehouse, error) {
	existingWarehouse := s.warehouseRepo.GetById(id)
	if existingWarehouse == nil {
		return nil, errors.WrapErrNotFound("warehouse", "id", id)
	}

	if warehouse.WarehouseCode != nil && *warehouse.WarehouseCode != *existingWarehouse.WarehouseCode {
		if s.warehouseRepo.ExistsByCode(*warehouse.WarehouseCode) {
			return nil, errors.WrapErrAlreadyExist("warehouse", "code", *warehouse.WarehouseCode)
		}
	}
	if warehouse.WarehouseCode == nil {
		warehouse.WarehouseCode = existingWarehouse.WarehouseCode
	}
	if warehouse.Address == nil {
		warehouse.Address = existingWarehouse.Address
	}
	if warehouse.Telephone == nil {
		warehouse.Telephone = existingWarehouse.Telephone
	}
	if warehouse.MinimumCapacity == nil {
		warehouse.MinimumCapacity = existingWarehouse.MinimumCapacity
	}
	if warehouse.MinimumTemperature == nil {
		warehouse.MinimumTemperature = existingWarehouse.MinimumTemperature
	}

	warehouse.ID = id
	updatedWarehouse := s.warehouseRepo.Update(id, warehouse)
	return updatedWarehouse, nil
}

func (s *WarehouseServiceImpl) DeleteWarehouse(id int) error {
	panic("unimplemented")
}