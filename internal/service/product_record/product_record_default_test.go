package service

import (
	"ProyectoFinal/mocks"
	pkgErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateProductRecord(t *testing.T) {
	t.Run("should create product record successfully", func(t *testing.T) {
		mockRepo := &mocks.MockProductRecordRepository{}
		service := NewProductRecordDefault(mockRepo)

		productRecord := models.ProductRecord{
			LastUpdateDate: "2024-01-15",
			PurchasePrice:  10.50,
			SalePrice:      15.99,
			ProductID:      1,
		}

		expectedRecord := models.ProductRecord{
			ID:             1,
			LastUpdateDate: "2024-01-15",
			PurchasePrice:  10.50,
			SalePrice:      15.99,
			ProductID:      1,
		}

		mockRepo.On("CreateProductRecord", productRecord).Return(expectedRecord, nil)

		result, err := service.CreateProductRecord(productRecord)

		require.NoError(t, err)
		assert.Equal(t, expectedRecord, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		mockRepo := &mocks.MockProductRecordRepository{}
		service := NewProductRecordDefault(mockRepo)

		productRecord := models.ProductRecord{
			LastUpdateDate: "2024-01-15",
			PurchasePrice:  10.50,
			SalePrice:      15.99,
			ProductID:      999,
		}

		repositoryError := pkgErrors.ErrNotFound

		mockRepo.On("CreateProductRecord", productRecord).Return(models.ProductRecord{}, repositoryError)

		result, err := service.CreateProductRecord(productRecord)

		require.Error(t, err)
		assert.Equal(t, models.ProductRecord{}, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetRecordsProduct(t *testing.T) {
	t.Run("should get records product successfully", func(t *testing.T) {
		mockRepo := &mocks.MockProductRecordRepository{}
		service := NewProductRecordDefault(mockRepo)

		productID := 1
		mockRepo.On("GetRecordsProduct", productID).Return(models.ReportProductData{}, nil)

		result, err := service.GetRecordsProduct(&productID)

		require.NoError(t, err)
		assert.Equal(t, models.ReportProductData{}, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error when repository fails", func(t *testing.T) {
		mockRepo := &mocks.MockProductRecordRepository{}
		service := NewProductRecordDefault(mockRepo)

		productID := 1
		mockRepo.On("GetRecordsProduct", productID).Return(models.ReportProductData{}, pkgErrors.ErrNotFound)

		result, err := service.GetRecordsProduct(&productID)

		require.Error(t, err)
		assert.Equal(t, models.ReportProductData{}, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetRecordsProductAll(t *testing.T) {
	t.Run("should get records product all successfully", func(t *testing.T) {
		mockRepo := &mocks.MockProductRecordRepository{}
		service := NewProductRecordDefault(mockRepo)

		mockRepo.On("GetRecordsProductAll").Return([]models.ReportProductData{}, nil)

		result, err := service.GetRecordsProductAll()

		require.NoError(t, err)
		assert.Equal(t, []models.ReportProductData{}, result)
		mockRepo.AssertExpectations(t)
	})
}
