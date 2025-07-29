package mocks

import (
	"ProyectoFinal/pkg/models"

	"github.com/stretchr/testify/mock"
)

type MockWarehouseRepository struct {
	mock.Mock
}

func (m *MockWarehouseRepository) GetAll() ([]models.Warehouse, error) {
	args := m.Called()
	return args.Get(0).([]models.Warehouse), args.Error(1)
}

func (m *MockWarehouseRepository) GetById(id int) (*models.Warehouse, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Warehouse), args.Error(1)
}

func (m *MockWarehouseRepository) Create(warehouse models.Warehouse) (models.Warehouse, error) {
	args := m.Called(warehouse)
	return args.Get(0).(models.Warehouse), args.Error(1)
}

func (m *MockWarehouseRepository) Update(id int, warehouse models.Warehouse) (models.Warehouse, error) {
	args := m.Called(id, warehouse)
	return args.Get(0).(models.Warehouse), args.Error(1)
}

func (m *MockWarehouseRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
