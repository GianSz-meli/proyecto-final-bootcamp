package warehouse

import (
	"errors"
	"testing"

	mocks "ProyectoFinal/mocks/warehouse"
	customErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"

	"github.com/stretchr/testify/require"
)

func TestUpdateWarehouse_Success(t *testing.T) {
	mockRepo := new(mocks.MockWarehouseRepository)
	service := NewWarehouseService(mockRepo)

	warehouseId := 1
	updateRequest := models.UpdateWarehouseRequest{
		WarehouseCode: &[]string{"WH001-UPDATED"}[0],
	}
	existingWarehouse := &models.Warehouse{
		ID:                 1,
		WarehouseCode:      "WH001",
		Address:            "Original Address",
		Telephone:          "1234567890",
		MinimumCapacity:    100,
		MinimumTemperature: -10.5,
		LocalityId:         &[]int{1}[0],
	}
	updatedWarehouse := *existingWarehouse
	updatedWarehouse.WarehouseCode = "WH001-UPDATED"

	mockRepo.On("GetById", warehouseId).Return(existingWarehouse, nil)
	mockRepo.On("Update", warehouseId, updatedWarehouse).Return(updatedWarehouse, nil)

	result, err := service.UpdateWarehouse(warehouseId, updateRequest)

	require.NoError(t, err)
	require.Equal(t, updatedWarehouse, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateWarehouse_GetByIdError_ReturnsError(t *testing.T) {
	mockRepo := new(mocks.MockWarehouseRepository)
	service := NewWarehouseService(mockRepo)

	warehouseId := 999
	updateRequest := models.UpdateWarehouseRequest{
		WarehouseCode: &[]string{"WH001-UPDATED"}[0],
	}

	expectedError := errors.New("database connection failed")
	mockRepo.On("GetById", warehouseId).Return(nil, expectedError)

	result, err := service.UpdateWarehouse(warehouseId, updateRequest)

	require.Error(t, err)
	require.Equal(t, expectedError, err)
	require.Equal(t, models.Warehouse{}, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateWarehouse_NoUpdateFieldsProvided_ReturnsUnprocessableEntityError(t *testing.T) {
	mockRepo := new(mocks.MockWarehouseRepository)
	service := NewWarehouseService(mockRepo)

	warehouseId := 1
	existingWarehouse := &models.Warehouse{
		ID:                 1,
		WarehouseCode:      "WH001",
		Address:            "Original Address",
		Telephone:          "1234567890",
		MinimumCapacity:    100,
		MinimumTemperature: -10.5,
		LocalityId:         &[]int{1}[0],
	}
	updateRequest := models.UpdateWarehouseRequest{}

	mockRepo.On("GetById", warehouseId).Return(existingWarehouse, nil)

	result, err := service.UpdateWarehouse(warehouseId, updateRequest)

	require.Error(t, err)
	require.ErrorIs(t, err, customErrors.ErrUnprocessableEntity)
	require.Equal(t, models.Warehouse{}, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateWarehouse_UpdateError_ReturnsError(t *testing.T) {
	mockRepo := new(mocks.MockWarehouseRepository)
	service := NewWarehouseService(mockRepo)

	warehouseId := 1
	updateRequest := models.UpdateWarehouseRequest{
		WarehouseCode: &[]string{"WH001-UPDATED"}[0],
	}
	existingWarehouse := &models.Warehouse{
		ID:                 1,
		WarehouseCode:      "WH001",
		Address:            "Original Address",
		Telephone:          "1234567890",
		MinimumCapacity:    100,
		MinimumTemperature: -10.5,
		LocalityId:         &[]int{1}[0],
	}
	updatedWarehouse := *existingWarehouse
	updatedWarehouse.WarehouseCode = "WH001-UPDATED"
	expectedError := errors.New("update error")

	mockRepo.On("GetById", warehouseId).Return(existingWarehouse, nil)
	mockRepo.On("Update", warehouseId, updatedWarehouse).Return(models.Warehouse{}, expectedError)

	result, err := service.UpdateWarehouse(warehouseId, updateRequest)

	require.Error(t, err)
	require.Equal(t, expectedError, err)
	require.Equal(t, models.Warehouse{}, result)
	mockRepo.AssertExpectations(t)
}

func TestGetAllWarehouses_Success(t *testing.T) {
	mockRepo := new(mocks.MockWarehouseRepository)
	service := NewWarehouseService(mockRepo)

	expectedWarehouses := []models.Warehouse{
		{
			ID:                 1,
			WarehouseCode:      "WH001",
			Address:            "Address 1",
			Telephone:          "1234567890",
			MinimumCapacity:    100,
			MinimumTemperature: -10.5,
			LocalityId:         &[]int{1}[0],
		},
		{
			ID:                 2,
			WarehouseCode:      "WH002",
			Address:            "Address 2",
			Telephone:          "0987654321",
			MinimumCapacity:    200,
			MinimumTemperature: -5.0,
			LocalityId:         nil,
		},
	}

	mockRepo.On("GetAll").Return(expectedWarehouses, nil)

	result, err := service.GetAllWarehouses()

	require.NoError(t, err)
	require.Equal(t, expectedWarehouses, result)
	mockRepo.AssertExpectations(t)
}

func TestGetAllWarehouses_RepositoryError_ReturnsError(t *testing.T) {
	mockRepo := new(mocks.MockWarehouseRepository)
	service := NewWarehouseService(mockRepo)

	expectedError := errors.New("database connection failed")
	mockRepo.On("GetAll").Return([]models.Warehouse{}, expectedError)

	result, err := service.GetAllWarehouses()

	require.Error(t, err)
	require.Equal(t, expectedError, err)
	require.Equal(t, []models.Warehouse{}, result)
	mockRepo.AssertExpectations(t)
}

func TestGetWarehouseById_Success(t *testing.T) {
	mockRepo := new(mocks.MockWarehouseRepository)
	service := NewWarehouseService(mockRepo)

	warehouseId := 1
	expectedWarehouse := &models.Warehouse{
		ID:                 1,
		WarehouseCode:      "WH001",
		Address:            "Address 1",
		Telephone:          "1234567890",
		MinimumCapacity:    100,
		MinimumTemperature: -10.5,
		LocalityId:         &[]int{1}[0],
	}

	mockRepo.On("GetById", warehouseId).Return(expectedWarehouse, nil)

	result, err := service.GetWarehouseById(warehouseId)

	require.NoError(t, err)
	require.Equal(t, *expectedWarehouse, result)
	mockRepo.AssertExpectations(t)
}

func TestGetWarehouseById_RepositoryError_ReturnsError(t *testing.T) {
	mockRepo := new(mocks.MockWarehouseRepository)
	service := NewWarehouseService(mockRepo)

	warehouseId := 999
	expectedError := errors.New("database connection failed")
	mockRepo.On("GetById", warehouseId).Return(nil, expectedError)

	result, err := service.GetWarehouseById(warehouseId)

	require.Error(t, err)
	require.Equal(t, expectedError, err)
	require.Equal(t, models.Warehouse{}, result)
	mockRepo.AssertExpectations(t)
}

func TestCreateWarehouse_Success(t *testing.T) {
	mockRepo := new(mocks.MockWarehouseRepository)
	service := NewWarehouseService(mockRepo)

	inputWarehouse := models.Warehouse{
		WarehouseCode:      "WH001",
		Address:            "Address 1",
		Telephone:          "1234567890",
		MinimumCapacity:    100,
		MinimumTemperature: -10.5,
		LocalityId:         &[]int{1}[0],
	}

	expectedWarehouse := inputWarehouse
	expectedWarehouse.ID = 1

	mockRepo.On("Create", inputWarehouse).Return(expectedWarehouse, nil)

	result, err := service.CreateWarehouse(inputWarehouse)

	require.NoError(t, err)
	require.Equal(t, expectedWarehouse, result)
	mockRepo.AssertExpectations(t)
}

func TestCreateWarehouse_RepositoryError_ReturnsError(t *testing.T) {
	mockRepo := new(mocks.MockWarehouseRepository)
	service := NewWarehouseService(mockRepo)

	inputWarehouse := models.Warehouse{
		WarehouseCode:      "WH001",
		Address:            "Address 1",
		Telephone:          "1234567890",
		MinimumCapacity:    100,
		MinimumTemperature: -10.5,
		LocalityId:         &[]int{1}[0],
	}

	expectedError := errors.New("database connection failed")
	mockRepo.On("Create", inputWarehouse).Return(models.Warehouse{}, expectedError)

	result, err := service.CreateWarehouse(inputWarehouse)

	require.Error(t, err)
	require.Equal(t, expectedError, err)
	require.Equal(t, models.Warehouse{}, result)
	mockRepo.AssertExpectations(t)
}

func TestDeleteWarehouse_Success(t *testing.T) {
	mockRepo := new(mocks.MockWarehouseRepository)
	service := NewWarehouseService(mockRepo)

	warehouseId := 1
	mockRepo.On("Delete", warehouseId).Return(nil)

	err := service.DeleteWarehouse(warehouseId)

	require.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteWarehouse_RepositoryError_ReturnsError(t *testing.T) {
	mockRepo := new(mocks.MockWarehouseRepository)
	service := NewWarehouseService(mockRepo)

	warehouseId := 999
	expectedError := customErrors.WrapErrNotFound("warehouse", "id", warehouseId)
	mockRepo.On("Delete", warehouseId).Return(expectedError)

	err := service.DeleteWarehouse(warehouseId)

	require.Error(t, err)
	require.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}
