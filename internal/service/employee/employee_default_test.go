package employee

import (
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockEmployeeRepository is a mock implementation of the employee.Repository interface
type MockEmployeeRepository struct {
	mock.Mock
}

func (m *MockEmployeeRepository) GetAll() ([]models.Employee, error) {
	args := m.Called()
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

func TestService_Create_Ok(t *testing.T) {
	// Arrange
	mockRepo := new(MockEmployeeRepository)
	service := NewService(mockRepo)

	warehouseID := 1
	employee := models.Employee{
		CardNumberID: "EMP001",
		FirstName:    "John",
		LastName:     "Doe",
		WarehouseID:  &warehouseID,
	}

	mockRepo.On("Create", &employee).Return(nil)

	// Act
	result, err := service.Create(employee)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, employee, result)
	mockRepo.AssertExpectations(t)
}

func TestService_Create_Conflict(t *testing.T) {
	// Arrange
	mockRepo := new(MockEmployeeRepository)
	service := NewService(mockRepo)

	warehouseID := 1
	employee := models.Employee{
		CardNumberID: "EMP001",
		FirstName:    "John",
		LastName:     "Doe",
		WarehouseID:  &warehouseID,
	}

	conflictErr := errors.WrapErrConflict("employee", "card_number_id", "EMP001")
	mockRepo.On("Create", &employee).Return(conflictErr)

	// Act
	result, err := service.Create(employee)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, models.Employee{}, result)
	assert.ErrorIs(t, err, errors.ErrConflict)
	mockRepo.AssertExpectations(t)
}

func TestService_FindAll(t *testing.T) {
	// Arrange
	mockRepo := new(MockEmployeeRepository)
	service := NewService(mockRepo)

	warehouseID1 := 1
	warehouseID2 := 2
	expectedEmployees := []models.Employee{
		{
			ID:           1,
			CardNumberID: "EMP001",
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  &warehouseID1,
		},
		{
			ID:           2,
			CardNumberID: "EMP002",
			FirstName:    "Jane",
			LastName:     "Smith",
			WarehouseID:  &warehouseID2,
		},
	}

	mockRepo.On("GetAll").Return(expectedEmployees, nil)

	// Act
	result, err := service.GetAll()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedEmployees, result)
	assert.Len(t, result, 2)
	mockRepo.AssertExpectations(t)
}

func TestService_FindByIdNonExistent(t *testing.T) {
	// Arrange
	mockRepo := new(MockEmployeeRepository)
	service := NewService(mockRepo)

	employeeID := 999
	notFoundErr := errors.WrapErrNotFound("employee", "id", employeeID)
	mockRepo.On("GetById", employeeID).Return(models.Employee{}, notFoundErr)

	// Act
	result, err := service.GetById(employeeID)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, models.Employee{}, result)
	assert.ErrorIs(t, err, errors.ErrNotFound)
	mockRepo.AssertExpectations(t)
}

func TestService_FindByIdExistent(t *testing.T) {
	// Arrange
	mockRepo := new(MockEmployeeRepository)
	service := NewService(mockRepo)

	employeeID := 1
	warehouseID := 1
	expectedEmployee := models.Employee{
		ID:           employeeID,
		CardNumberID: "EMP001",
		FirstName:    "John",
		LastName:     "Doe",
		WarehouseID:  &warehouseID,
	}

	mockRepo.On("GetById", employeeID).Return(expectedEmployee, nil)

	// Act
	result, err := service.GetById(employeeID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedEmployee, result)
	mockRepo.AssertExpectations(t)
}

func TestService_UpdateExistent(t *testing.T) {
	// Arrange
	mockRepo := new(MockEmployeeRepository)
	service := NewService(mockRepo)

	employeeID := 1
	warehouseID := 1
	newWarehouseID := 2

	existingEmployee := models.Employee{
		ID:           employeeID,
		CardNumberID: "EMP001",
		FirstName:    "John",
		LastName:     "Doe",
		WarehouseID:  &warehouseID,
	}

	newFirstName := "Johnny"
	updateRequest := &models.EmployeeUpdateRequest{
		FirstName:   &newFirstName,
		WarehouseID: &newWarehouseID,
	}

	updatedEmployee := models.Employee{
		ID:           employeeID,
		CardNumberID: "EMP001",
		FirstName:    newFirstName,
		LastName:     "Doe",
		WarehouseID:  &newWarehouseID,
	}

	mockRepo.On("GetById", employeeID).Return(existingEmployee, nil)
	mockRepo.On("Update", employeeID, updatedEmployee).Return(nil)

	// Act
	result, err := service.PatchUpdate(employeeID, updateRequest)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, updatedEmployee, result)
	assert.Equal(t, newFirstName, result.FirstName)
	assert.Equal(t, &newWarehouseID, result.WarehouseID)
	mockRepo.AssertExpectations(t)
}

func TestService_UpdateNonExistent(t *testing.T) {
	// Arrange
	mockRepo := new(MockEmployeeRepository)
	service := NewService(mockRepo)

	employeeID := 999
	notFoundErr := errors.WrapErrNotFound("employee", "id", employeeID)

	newFirstName := "Johnny"
	updateRequest := &models.EmployeeUpdateRequest{
		FirstName: &newFirstName,
	}

	mockRepo.On("GetById", employeeID).Return(models.Employee{}, notFoundErr)

	// Act
	result, err := service.PatchUpdate(employeeID, updateRequest)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, models.Employee{}, result)
	assert.ErrorIs(t, err, errors.ErrNotFound)
	mockRepo.AssertExpectations(t)
}

func TestService_DeleteNonExistent(t *testing.T) {
	// Arrange
	mockRepo := new(MockEmployeeRepository)
	service := NewService(mockRepo)

	employeeID := 999
	notFoundErr := errors.WrapErrNotFound("employee", "id", employeeID)
	mockRepo.On("Delete", employeeID).Return(notFoundErr)

	// Act
	err := service.Delete(employeeID)

	// Assert
	assert.Error(t, err)
	assert.ErrorIs(t, err, errors.ErrNotFound)
	mockRepo.AssertExpectations(t)
}

func TestService_DeleteOk(t *testing.T) {
	// Arrange
	mockRepo := new(MockEmployeeRepository)
	service := NewService(mockRepo)

	employeeID := 1
	mockRepo.On("Delete", employeeID).Return(nil)

	// Act
	err := service.Delete(employeeID)

	// Assert
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestService_UpdateWithNoFields(t *testing.T) {
	// Arrange
	mockRepo := new(MockEmployeeRepository)
	service := NewService(mockRepo)

	employeeID := 1
	warehouseID := 1

	existingEmployee := models.Employee{
		ID:           employeeID,
		CardNumberID: "EMP001",
		FirstName:    "John",
		LastName:     "Doe",
		WarehouseID:  &warehouseID,
	}

	updateRequest := &models.EmployeeUpdateRequest{}

	mockRepo.On("GetById", employeeID).Return(existingEmployee, nil)

	// Act
	result, err := service.PatchUpdate(employeeID, updateRequest)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, models.Employee{}, result)
	assert.ErrorIs(t, err, errors.ErrUnprocessableEntity)
	mockRepo.AssertExpectations(t)
}
