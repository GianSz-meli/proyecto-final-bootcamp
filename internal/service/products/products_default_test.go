package service

import (
	"ProyectoFinal/mocks/products"
	pkgErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAllProducts(t *testing.T) {
	mockRepo := &mocks.MockProductRepository{}
	service := NewProductDefault(mockRepo)

	productsMap := map[int]models.Product{
		1: {
			ID:          1,
			ProductCode: "A1000",
			Description: "Coca Cola 2L",
			Width:       5,
			Height:      20,
			Length:      12,
		},
	}

	mockRepo.On("FindAllProducts").Return(productsMap, nil)

	result, err := service.FindAllProducts()
	require.NoError(t, err)
	require.Equal(t, productsMap, result)
	mockRepo.AssertExpectations(t)
}

func ptrInt(v int) *int { return &v }

func TestCreateProduct(t *testing.T) {
	t.Run("should create product successfully when product code does not exist", func(t *testing.T) {
		mockRepo := &mocks.MockProductRepository{}
		service := NewProductDefault(mockRepo)

		product := models.Product{
			ProductCode:    "A1000",
			Description:    "Coca Cola 2L",
			Width:          5.0,
			Height:         20.0,
			Length:         12.0,
			NetWeight:      1.9,
			ExpirationRate: 15.0,
			Temperature:    10.0,
			FreezingRate:   10.0,
			ProductTypeID:  83,
			SellerID:       ptrInt(7),
		}

		mockRepo.On("ExistsProdCode", product.ProductCode).Return(false)
		mockRepo.On("CreateProduct", product).Return(product, nil)

		result, err := service.CreateProduct(product)

		require.NoError(t, err)
		require.Equal(t, product, result)

		mockRepo.AssertCalled(t, "ExistsProdCode", product.ProductCode)
		mockRepo.AssertCalled(t, "CreateProduct", product)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return conflict error when product code already exists", func(t *testing.T) {
		mockRepo := &mocks.MockProductRepository{}
		service := NewProductDefault(mockRepo)

		existingProduct := models.Product{
			ProductCode:    "A1000",
			Description:    "Coca Cola 2L",
			Width:          5,
			Height:         20,
			Length:         12,
			NetWeight:      2.0,
			ExpirationRate: 30,
			Temperature:    4.0,
			FreezingRate:   0.5,
			ProductTypeID:  1,
		}

		mockRepo.On("ExistsProdCode", "A1000").Return(true)

		result, err := service.CreateProduct(existingProduct)

		require.Error(t, err)
		assert.Equal(t, models.Product{}, result)

		mockRepo.AssertNotCalled(t, "CreateProduct")
		mockRepo.AssertExpectations(t)
	})
}

func TestFindProductsById(t *testing.T) {
	t.Run("should return product successfully when product exists", func(t *testing.T) {
		mockRepo := &mocks.MockProductRepository{}
		service := NewProductDefault(mockRepo)

		expectedProduct := models.Product{
			ID:             1,
			ProductCode:    "A1000",
			Description:    "Coca Cola 2L",
			Width:          5,
			Height:         20,
			Length:         12,
			NetWeight:      2.0,
			ExpirationRate: 30,
			Temperature:    4.0,
			FreezingRate:   0.5,
			ProductTypeID:  1,
		}

		mockRepo.On("FindProductsById", 1).Return(expectedProduct, nil)

		result, err := service.FindProductsById(1)

		require.NoError(t, err)
		assert.Equal(t, expectedProduct, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return not found error when product does not exist", func(t *testing.T) {
		mockRepo := &mocks.MockProductRepository{}
		service := NewProductDefault(mockRepo)

		repositoryError := fmt.Errorf("product not found in database")

		mockRepo.On("FindProductsById", 999).Return(models.Product{}, repositoryError)

		result, err := service.FindProductsById(999)

		require.Error(t, err)
		assert.Equal(t, models.Product{}, result)

		assert.ErrorIs(t, err, pkgErrors.ErrNotFound)

		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateProduct(t *testing.T) {

	existingProduct := models.Product{
		ID:          1,
		ProductCode: "A1000",
		Description: "Coca Cola 2L",
		Width:       5,
		Height:      20,
		Length:      12,
	}
	t.Run("should update product successfully when product exists", func(t *testing.T) {
		mockRepo := &mocks.MockProductRepository{}
		service := NewProductDefault(mockRepo)

		mockRepo.On("FindProductsById", 1).Return(existingProduct, nil)
		mockRepo.On("UpdateProduct", 1, existingProduct).Return(existingProduct, nil)

		result, err := service.UpdateProduct(1, models.ProductDocUpdate{
			ProductCode: &existingProduct.ProductCode,
			Description: &existingProduct.Description,
			Width:       &existingProduct.Width,
			Height:      &existingProduct.Height,
			Length:      &existingProduct.Length,
		})

		require.NoError(t, err)
		assert.Equal(t, existingProduct, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return not found error when product does not exist", func(t *testing.T) {
		mockRepo := &mocks.MockProductRepository{}
		service := NewProductDefault(mockRepo)

		repositoryError := pkgErrors.ErrNotFound

		mockRepo.On("FindProductsById", 999).Return(models.Product{}, repositoryError)

		result, err := service.UpdateProduct(999, models.ProductDocUpdate{
			ProductCode: &existingProduct.ProductCode,
			Description: &existingProduct.Description,
			Width:       &existingProduct.Width,
			Height:      &existingProduct.Height,
			Length:      &existingProduct.Length,
		})

		require.Error(t, err)
		assert.Equal(t, models.Product{}, result)

		assert.ErrorIs(t, err, pkgErrors.ErrNotFound)

		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteProduct(t *testing.T) {
	t.Run("should delete product successfully when product exists", func(t *testing.T) {
		mockRepo := &mocks.MockProductRepository{}
		service := NewProductDefault(mockRepo)

		mockRepo.On("FindProductsById", 1).Return(models.Product{}, nil)
		mockRepo.On("DeleteProduct", 1).Return(nil)

		err := service.DeleteProduct(1)

		require.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return not found error when product does not exist", func(t *testing.T) {
		mockRepo := &mocks.MockProductRepository{}
		service := NewProductDefault(mockRepo)

		repositoryError := pkgErrors.ErrNotFound

		mockRepo.On("FindProductsById", 999).Return(models.Product{}, repositoryError)

		err := service.DeleteProduct(999)

		require.Error(t, err)
		assert.ErrorIs(t, err, pkgErrors.ErrNotFound)
		mockRepo.AssertExpectations(t)
	})
}
