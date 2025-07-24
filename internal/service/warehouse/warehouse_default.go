package warehouse

import (
	"ProyectoFinal/internal/repository/warehouse"
	"ProyectoFinal/internal/service/utils"
	customErrors "ProyectoFinal/pkg/errors"
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

func (s *WarehouseServiceImpl) GetAllWarehouses() ([]models.Warehouse, error) {
	warehouses, err := s.warehouseRepo.GetAll()
	if err != nil {
		return []models.Warehouse{}, err
	}
	return warehouses, nil
}

func (s *WarehouseServiceImpl) GetWarehouseById(id int) (models.Warehouse, error) {
	wh, err := s.warehouseRepo.GetById(id)
	if err != nil {
		return models.Warehouse{}, err
	}
	return *wh, nil
}

func (s *WarehouseServiceImpl) CreateWarehouse(warehouse models.Warehouse) (models.Warehouse, error) {
	wh, err := s.warehouseRepo.Create(warehouse)
	if err != nil {
		return models.Warehouse{}, err
	}
	return wh, nil
}

func (s *WarehouseServiceImpl) UpdateWarehouse(id int, warehouseUpdate models.UpdateWarehouseRequest) (models.Warehouse, error) {
	existingWarehouse, err := s.warehouseRepo.GetById(id)
	if err != nil {
		return models.Warehouse{}, err
	}

	if updated := utils.UpdateFields(existingWarehouse, &warehouseUpdate); !updated {
		return models.Warehouse{}, customErrors.WrapErrUnprocessableEntity(fmt.Errorf("no fields provided for update"))
	}

	updatedWarehouse, err := s.warehouseRepo.Update(id, *existingWarehouse)
	if err != nil {
		return models.Warehouse{}, err
	}
	return updatedWarehouse, nil
}

func (s *WarehouseServiceImpl) DeleteWarehouse(id int) error {
	err := s.warehouseRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
