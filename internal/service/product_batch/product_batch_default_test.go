package product_batch

import (
	mocks "ProyectoFinal/mocks/product_batch"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProductBatchService_Create_ValidProductBatch_Success(t *testing.T) {
	// Arrange
	mockRepository := new(mocks.MockProductBatchRepository)
	service := NewProductBatchService(mockRepository)

	productBatch := models.ProductBatch{
		BatchNumber:        "BATCH001",
		CurrentQuantity:    100,
		CurrentTemperature: 15.5,
		DueDate:            time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
		InitialQuantity:    100,
		ManufacturingDate:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		ManufacturingHour:  10,
		MinimumTemperature: 10.0,
		ProductID:          1,
		SectionID:          1,
	}

	createdProductBatch := models.ProductBatch{
		ID:                 1,
		BatchNumber:        "BATCH001",
		CurrentQuantity:    100,
		CurrentTemperature: 15.5,
		DueDate:            time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
		InitialQuantity:    100,
		ManufacturingDate:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		ManufacturingHour:  10,
		MinimumTemperature: 10.0,
		ProductID:          1,
		SectionID:          1,
	}

	mockRepository.On("Create", productBatch).Return(createdProductBatch, nil)

	// Act
	result, err := service.Create(productBatch)

	// Assert
	require.NoError(t, err)
	assert.Equal(t, createdProductBatch, result)
	mockRepository.AssertExpectations(t)
}

func TestProductBatchService_Create_RepositoryError_ReturnsError(t *testing.T) {
	// Arrange
	mockRepository := new(mocks.MockProductBatchRepository)
	service := NewProductBatchService(mockRepository)

	productBatch := models.ProductBatch{
		BatchNumber:        "BATCH001",
		CurrentQuantity:    100,
		CurrentTemperature: 15.5,
		DueDate:            time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
		InitialQuantity:    100,
		ManufacturingDate:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		ManufacturingHour:  10,
		MinimumTemperature: 10.0,
		ProductID:          1,
		SectionID:          1,
	}

	mockRepository.On("Create", productBatch).Return(models.ProductBatch{}, errors.ErrGeneral)

	// Act
	result, err := service.Create(productBatch)

	// Assert
	require.Error(t, err)
	assert.Equal(t, models.ProductBatch{}, result)
	assert.ErrorIs(t, err, errors.ErrGeneral)
	mockRepository.AssertExpectations(t)
}

func TestProductBatchService_GetProductCountBySection_ValidSectionID_Success(t *testing.T) {
	// Arrange
	mockRepository := new(mocks.MockProductBatchRepository)
	service := NewProductBatchService(mockRepository)

	sectionID := 1
	expectedReports := []models.SectionProductReport{
		{
			SectionID:     1,
			SectionNumber: "SEC001",
			ProductsCount: 5,
		},
	}

	mockRepository.On("GetProductCountBySection", &sectionID).Return(expectedReports, nil)

	// Act
	result, err := service.GetProductCountBySection(&sectionID)

	// Assert
	require.NoError(t, err)
	assert.Equal(t, expectedReports, result)
	assert.Len(t, result, 1)
	mockRepository.AssertExpectations(t)
}

func TestProductBatchService_GetProductCountBySection_AllSections_Success(t *testing.T) {
	// Arrange
	mockRepository := new(mocks.MockProductBatchRepository)
	service := NewProductBatchService(mockRepository)

	expectedReports := []models.SectionProductReport{
		{
			SectionID:     1,
			SectionNumber: "SEC001",
			ProductsCount: 5,
		},
		{
			SectionID:     2,
			SectionNumber: "SEC002",
			ProductsCount: 3,
		},
		{
			SectionID:     3,
			SectionNumber: "SEC003",
			ProductsCount: 7,
		},
	}

	mockRepository.On("GetProductCountBySection", (*int)(nil)).Return(expectedReports, nil)

	// Act
	result, err := service.GetProductCountBySection(nil)

	// Assert
	require.NoError(t, err)
	assert.Equal(t, expectedReports, result)
	assert.Len(t, result, 3)
	mockRepository.AssertExpectations(t)
}

func TestProductBatchService_GetProductCountBySection_RepositoryError_ReturnsError(t *testing.T) {
	// Arrange
	mockRepository := new(mocks.MockProductBatchRepository)
	service := NewProductBatchService(mockRepository)

	sectionID := 1

	mockRepository.On("GetProductCountBySection", &sectionID).Return(nil, errors.ErrGeneral)

	// Act
	result, err := service.GetProductCountBySection(&sectionID)

	// Assert
	require.Error(t, err)
	assert.Nil(t, result)
	assert.ErrorIs(t, err, errors.ErrGeneral)
	mockRepository.AssertExpectations(t)
}

func TestProductBatchService_GetProductCountBySection_EmptyResult_Success(t *testing.T) {
	// Arrange
	mockRepository := new(mocks.MockProductBatchRepository)
	service := NewProductBatchService(mockRepository)

	sectionID := 999
	expectedReports := []models.SectionProductReport{}

	mockRepository.On("GetProductCountBySection", &sectionID).Return(expectedReports, nil)

	// Act
	result, err := service.GetProductCountBySection(&sectionID)

	// Assert
	require.NoError(t, err)
	assert.Equal(t, expectedReports, result)
	assert.Len(t, result, 0)
	mockRepository.AssertExpectations(t)
}
