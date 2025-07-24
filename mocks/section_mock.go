package mocks

import (
	"ProyectoFinal/pkg/models"

	"github.com/stretchr/testify/mock"
)

// MockSectionService - implementation of section interface
type MockSectionService struct {
	mock.Mock
}

func (m *MockSectionService) GetAll() ([]models.Section, error) {
	args := m.Called()
	return args.Get(0).([]models.Section), args.Error(1)
}

func (m *MockSectionService) GetById(id int) (models.Section, error) {
	args := m.Called(id)
	return args.Get(0).(models.Section), args.Error(1)
}

func (m *MockSectionService) Create(section models.Section) (models.Section, error) {
	args := m.Called(section)
	return args.Get(0).(models.Section), args.Error(1)
}

func (m *MockSectionService) Update(id int, section models.Section) (models.Section, error) {
	args := m.Called(id, section)
	return args.Get(0).(models.Section), args.Error(1)
}

func (m *MockSectionService) UpdateWithValidation(id int, updateRequest models.UpdateSectionRequest) (models.Section, error) {
	args := m.Called(id, updateRequest)
	return args.Get(0).(models.Section), args.Error(1)
}

func (m *MockSectionService) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
