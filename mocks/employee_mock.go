package mocks

import (
	"ProyectoFinal/pkg/models"

	"github.com/stretchr/testify/mock"
)

// MockEmployeeService is a mock implementation of the employee.Service interface
type MockEmployeeService struct {
	mock.Mock
}

func (m *MockEmployeeService) GetAll() ([]models.Employee, error) {
	args := m.Called()
	return args.Get(0).([]models.Employee), args.Error(1)
}

func (m *MockEmployeeService) GetById(id int) (models.Employee, error) {
	args := m.Called(id)
	return args.Get(0).(models.Employee), args.Error(1)
}

func (m *MockEmployeeService) Create(employee models.Employee) (models.Employee, error) {
	args := m.Called(employee)
	return args.Get(0).(models.Employee), args.Error(1)
}

func (m *MockEmployeeService) PatchUpdate(id int, updateRequest *models.EmployeeUpdateRequest) (models.Employee, error) {
	args := m.Called(id, updateRequest)
	return args.Get(0).(models.Employee), args.Error(1)
}

func (m *MockEmployeeService) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// MockEmployeeRepository is a mock implementation of the employee.Repository interface
type MockEmployeeRepository struct {
	mock.Mock
}

func (m *MockEmployeeRepository) GetAll() ([]models.Employee, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Employee), args.Error(1)
}

func (m *MockEmployeeRepository) GetById(id int) (models.Employee, error) {
	args := m.Called(id)
	return args.Get(0).(models.Employee), args.Error(1)
}

func (m *MockEmployeeRepository) Create(employee *models.Employee) error {
	args := m.Called(employee)
	return args.Error(0)
}

func (m *MockEmployeeRepository) Update(id int, employee models.Employee) error {
	args := m.Called(id, employee)
	return args.Error(0)
}

func (m *MockEmployeeRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
