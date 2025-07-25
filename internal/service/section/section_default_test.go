package service

import (
	"ProyectoFinal/mocks"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// Helper functions for creating pointers
func stringPtr(s string) *string {
	return &s
}

func float64Ptr(f float64) *float64 {
	return &f
}

func intPtr(i int) *int {
	return &i
}

func TestSectionService_GetAll_Success(t *testing.T) {
	// Arrange
	mockRepository := new(mocks.MockSectionRepository)
	service := NewSectionDefault(mockRepository)

	sections := []models.Section{
		{
			ID: 1,
			SectionAttributes: models.SectionAttributes{
				SectionNumber:      "SEC001",
				CurrentTemperature: 15.5,
				MinimumTemperature: 10.0,
				CurrentCapacity:    50,
				MinimumCapacity:    20,
				MaximumCapacity:    100,
				WarehouseID:        1,
				ProductTypeID:      1,
			},
		},
		{
			ID: 2,
			SectionAttributes: models.SectionAttributes{
				SectionNumber:      "SEC002",
				CurrentTemperature: 20.0,
				MinimumTemperature: 15.0,
				CurrentCapacity:    75,
				MinimumCapacity:    30,
				MaximumCapacity:    150,
				WarehouseID:        1,
				ProductTypeID:      2,
			},
		},
	}

	mockRepository.On("GetAll").Return(sections, nil)

	// Act
	result, err := service.GetAll()

	// Assert
	require.NoError(t, err)
	require.Equal(t, 2, len(result))
	require.Equal(t, sections[0], result[0])
	require.Equal(t, sections[1], result[1])
	mockRepository.AssertExpectations(t)
}

func TestSectionService_GetById_Existent_Success(t *testing.T) {
	// Arrange
	mockRepository := new(mocks.MockSectionRepository)
	service := NewSectionDefault(mockRepository)

	expectedSection := models.Section{
		ID: 1,
		SectionAttributes: models.SectionAttributes{
			SectionNumber:      "SEC001",
			CurrentTemperature: 15.5,
			MinimumTemperature: 10.0,
			CurrentCapacity:    50,
			MinimumCapacity:    20,
			MaximumCapacity:    100,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
	}

	mockRepository.On("GetById", 1).Return(expectedSection, nil)

	// Act
	result, err := service.GetById(1)

	// Assert
	require.NoError(t, err)
	require.Equal(t, expectedSection, result)
	mockRepository.AssertExpectations(t)
}

func TestSectionService_GetById_NonExistent_ReturnsEmpty(t *testing.T) {
	// Arrange
	mockRepository := new(mocks.MockSectionRepository)
	service := NewSectionDefault(mockRepository)

	notFoundError := errors.WrapErrNotFound("Section", "id", 999)
	mockRepository.On("GetById", 999).Return(models.Section{}, notFoundError)

	// Act
	result, err := service.GetById(999)

	// Assert
	require.Error(t, err)
	require.Equal(t, models.Section{}, result)
	require.Equal(t, notFoundError, err)
	mockRepository.AssertExpectations(t)
}

func TestSectionService_Create_ValidSection_Success(t *testing.T) {
	// Arrange
	mockRepository := new(mocks.MockSectionRepository)
	service := NewSectionDefault(mockRepository)

	sectionToCreate := models.Section{
		SectionAttributes: models.SectionAttributes{
			SectionNumber:      "SEC001",
			CurrentTemperature: 15.5,
			MinimumTemperature: 10.0,
			CurrentCapacity:    50,
			MinimumCapacity:    20,
			MaximumCapacity:    100,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
	}

	createdSection := models.Section{
		ID: 1,
		SectionAttributes: models.SectionAttributes{
			SectionNumber:      "SEC001",
			CurrentTemperature: 15.5,
			MinimumTemperature: 10.0,
			CurrentCapacity:    50,
			MinimumCapacity:    20,
			MaximumCapacity:    100,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
	}

	mockRepository.On("Create", sectionToCreate).Return(createdSection, nil)

	// Act
	result, err := service.Create(sectionToCreate)

	// Assert
	require.NoError(t, err)
	require.Equal(t, createdSection, result)
	require.Equal(t, 1, result.ID)
	mockRepository.AssertExpectations(t)
}

func TestSectionService_Create_Conflict_Error(t *testing.T) {
	// Arrange
	mockRepository := new(mocks.MockSectionRepository)
	service := NewSectionDefault(mockRepository)

	sectionToCreate := models.Section{
		SectionAttributes: models.SectionAttributes{
			SectionNumber:      "SEC001",
			CurrentTemperature: 15.5,
			MinimumTemperature: 10.0,
			CurrentCapacity:    50,
			MinimumCapacity:    20,
			MaximumCapacity:    100,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
	}

	conflictError := errors.WrapErrConflict("Section", "section_number", "SEC001")
	mockRepository.On("Create", sectionToCreate).Return(models.Section{}, conflictError)

	// Act
	result, err := service.Create(sectionToCreate)

	// Assert
	require.Error(t, err)
	require.Equal(t, models.Section{}, result)
	require.Equal(t, conflictError, err)
	mockRepository.AssertExpectations(t)
}

func TestSectionService_Update_Existent_Success(t *testing.T) {
	// Arrange
	mockRepository := new(mocks.MockSectionRepository)
	service := NewSectionDefault(mockRepository)

	existingSection := models.Section{
		ID: 1,
		SectionAttributes: models.SectionAttributes{
			SectionNumber:      "SEC001",
			CurrentTemperature: 15.5,
			MinimumTemperature: 10.0,
			CurrentCapacity:    50,
			MinimumCapacity:    20,
			MaximumCapacity:    100,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
	}

	sectionToUpdate := models.Section{
		ID: 1,
		SectionAttributes: models.SectionAttributes{
			SectionNumber:      "SEC001-UPDATED",
			CurrentTemperature: 20.0,
			MinimumTemperature: 10.0,
			CurrentCapacity:    50,
			MinimumCapacity:    20,
			MaximumCapacity:    150,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
	}

	updatedSection := models.Section{
		ID: 1,
		SectionAttributes: models.SectionAttributes{
			SectionNumber:      "SEC001-UPDATED",
			CurrentTemperature: 20.0,
			MinimumTemperature: 10.0,
			CurrentCapacity:    50,
			MinimumCapacity:    20,
			MaximumCapacity:    150,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
	}

	mockRepository.On("GetById", 1).Return(existingSection, nil)
	mockRepository.On("Update", 1, sectionToUpdate).Return(updatedSection, nil)

	// Act
	result, err := service.Update(1, sectionToUpdate)

	// Assert
	require.NoError(t, err)
	require.Equal(t, updatedSection, result)
	require.Equal(t, 1, result.ID)
	mockRepository.AssertExpectations(t)
}

func TestSectionService_Update_NonExistent_ReturnsEmpty(t *testing.T) {
	// Arrange
	mockRepository := new(mocks.MockSectionRepository)
	service := NewSectionDefault(mockRepository)

	sectionToUpdate := models.Section{
		SectionAttributes: models.SectionAttributes{
			SectionNumber:      "SEC999-UPDATED",
			CurrentTemperature: 20.0,
			MinimumTemperature: 10.0,
			CurrentCapacity:    50,
			MinimumCapacity:    20,
			MaximumCapacity:    150,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
	}

	notFoundError := errors.WrapErrNotFound("Section", "id", 999)
	mockRepository.On("GetById", 999).Return(models.Section{}, notFoundError)

	// Act
	result, err := service.Update(999, sectionToUpdate)

	// Assert
	require.Error(t, err)
	require.Equal(t, models.Section{}, result)
	require.Equal(t, notFoundError, err)
	mockRepository.AssertExpectations(t)
}

func TestSectionService_UpdateWithValidation_Existent_Success(t *testing.T) {
	// Arrange
	mockRepository := new(mocks.MockSectionRepository)
	service := NewSectionDefault(mockRepository)

	existingSection := models.Section{
		ID: 1,
		SectionAttributes: models.SectionAttributes{
			SectionNumber:      "SEC001",
			CurrentTemperature: 15.5,
			MinimumTemperature: 10.0,
			CurrentCapacity:    50,
			MinimumCapacity:    20,
			MaximumCapacity:    100,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
	}

	updateRequest := models.UpdateSectionRequest{
		SectionNumber:      stringPtr("SEC001-UPDATED"),
		CurrentTemperature: float64Ptr(20.0),
		MaximumCapacity:    intPtr(150),
	}

	updatedSection := models.Section{
		ID: 1,
		SectionAttributes: models.SectionAttributes{
			SectionNumber:      "SEC001-UPDATED",
			CurrentTemperature: 20.0,
			MinimumTemperature: 10.0,
			CurrentCapacity:    50,
			MinimumCapacity:    20,
			MaximumCapacity:    150,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
	}

	mockRepository.On("GetById", 1).Return(existingSection, nil)
	mockRepository.On("Update", 1, mock.AnythingOfType("models.Section")).Return(updatedSection, nil)

	// Act
	result, err := service.UpdateWithValidation(1, updateRequest)

	// Assert
	require.NoError(t, err)
	require.Equal(t, updatedSection, result)
	require.Equal(t, 1, result.ID)
	mockRepository.AssertExpectations(t)
}

func TestSectionService_UpdateWithValidation_NonExistent_ReturnsEmpty(t *testing.T) {
	// Arrange
	mockRepository := new(mocks.MockSectionRepository)
	service := NewSectionDefault(mockRepository)

	updateRequest := models.UpdateSectionRequest{
		SectionNumber: stringPtr("SEC999-UPDATED"),
	}

	notFoundError := errors.WrapErrNotFound("Section", "id", 999)
	mockRepository.On("GetById", 999).Return(models.Section{}, notFoundError)

	// Act
	result, err := service.UpdateWithValidation(999, updateRequest)

	// Assert
	require.Error(t, err)
	require.Equal(t, models.Section{}, result)
	require.Equal(t, notFoundError, err)
	mockRepository.AssertExpectations(t)
}

func TestSectionService_Delete_Existent_Success(t *testing.T) {
	// Arrange
	mockRepository := new(mocks.MockSectionRepository)
	service := NewSectionDefault(mockRepository)

	existingSection := models.Section{
		ID: 1,
		SectionAttributes: models.SectionAttributes{
			SectionNumber:      "SEC001",
			CurrentTemperature: 15.5,
			MinimumTemperature: 10.0,
			CurrentCapacity:    50,
			MinimumCapacity:    20,
			MaximumCapacity:    100,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
	}

	mockRepository.On("GetById", 1).Return(existingSection, nil)
	mockRepository.On("Delete", 1).Return(nil)

	// Act
	err := service.Delete(1)

	// Assert
	require.NoError(t, err)
	mockRepository.AssertExpectations(t)
}

func TestSectionService_Delete_NonExistent_ReturnsError(t *testing.T) {
	// Arrange
	mockRepository := new(mocks.MockSectionRepository)
	service := NewSectionDefault(mockRepository)

	notFoundError := errors.WrapErrNotFound("Section", "id", 999)
	mockRepository.On("GetById", 999).Return(models.Section{}, notFoundError)

	// Act
	err := service.Delete(999)

	// Assert
	require.Error(t, err)
	require.Equal(t, notFoundError, err)
	mockRepository.AssertExpectations(t)
}
