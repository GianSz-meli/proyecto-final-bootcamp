package mocks

import (
	"ProyectoFinal/pkg/models"

	"github.com/stretchr/testify/mock"
)

// MockWarehouseService is a mock implementation of the warehouse.WarehouseService interface
type MockWarehouseService struct {
	mock.Mock
}

func (m *MockWarehouseService) GetAllWarehouses() ([]models.Warehouse, error) {
	args := m.Called()
	return args.Get(0).([]models.Warehouse), args.Error(1)
}

func (m *MockWarehouseService) GetWarehouseById(id int) (models.Warehouse, error) {
	args := m.Called(id)
	return args.Get(0).(models.Warehouse), args.Error(1)
}

func (m *MockWarehouseService) CreateWarehouse(warehouse models.Warehouse) (models.Warehouse, error) {
	args := m.Called(warehouse)
	return args.Get(0).(models.Warehouse), args.Error(1)
}

func (m *MockWarehouseService) UpdateWarehouse(id int, warehouse models.Warehouse) (models.Warehouse, error) {
	args := m.Called(id, warehouse)
	return args.Get(0).(models.Warehouse), args.Error(1)
}

func (m *MockWarehouseService) DeleteWarehouse(id int) error {
	args := m.Called(id)
	return args.Error(0)
}