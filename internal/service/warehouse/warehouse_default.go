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
	return s.warehouseRepo.GetAll()
}

func (s *WarehouseServiceImpl) GetWarehouseById(id int) (models.Warehouse, error) {
	wh := s.warehouseRepo.GetById(id)
	if wh == nil {
		return models.Warehouse{}, errors.WrapErrNotFound("warehouse", "id", id)
	}
	return *wh, nil
}

func (s *WarehouseServiceImpl) CreateWarehouse(warehouse models.Warehouse) (models.Warehouse, error) {
	if s.warehouseRepo.ExistsByCode(warehouse.WarehouseCode) {
		return models.Warehouse{}, errors.WrapErrConflict("warehouse", "code", warehouse.WarehouseCode)
	}

	wh := s.warehouseRepo.Create(warehouse)
	return wh, nil
}

func (s *WarehouseServiceImpl) UpdateWarehouse(id int, warehouse models.Warehouse) (models.Warehouse, error) {
	existingWarehouse := s.warehouseRepo.GetById(id)
	if existingWarehouse == nil {
		return models.Warehouse{}, errors.WrapErrNotFound("warehouse", "id", id)
	}

	if warehouse.WarehouseCode != existingWarehouse.WarehouseCode {
		if s.warehouseRepo.ExistsByCode(warehouse.WarehouseCode) {
			return models.Warehouse{}, errors.WrapErrConflict("warehouse", "code", warehouse.WarehouseCode)
		}
	}

	warehouse.ID = id
	updatedWarehouse := s.warehouseRepo.Update(id, warehouse)
	return updatedWarehouse, nil
}

func (s *WarehouseServiceImpl) DeleteWarehouse(id int) error {
	existingWarehouse := s.warehouseRepo.GetById(id)
	if existingWarehouse == nil {
		return errors.WrapErrNotFound("warehouse", "id", id)
	}

	s.warehouseRepo.Delete(id)
	return nil
}
