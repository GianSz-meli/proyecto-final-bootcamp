package handler_test

import (
	"ProyectoFinal/internal/handler"
	"ProyectoFinal/pkg/models"
	pkgErrors "ProyectoFinal/pkg/errors"
	"encoding/json"
	// "errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type MockProductService struct {
	CreateProductFunc    func(models.Product) (models.Product, error)
	FindAllProductsFunc  func() (map[int]models.Product, error)
	FindProductsByIdFunc func(int) (models.Product, error)
	UpdateProductFunc    func(int, models.Product) (models.Product, error)
	DeleteProductFunc    func(int) error
}

func (m *MockProductService) CreateProduct(p models.Product) (models.Product, error) {
	if m.CreateProductFunc != nil {
		return m.CreateProductFunc(p)
	}
	return models.Product{}, nil
}
func (m *MockProductService) FindAllProducts() (map[int]models.Product, error) {
	if m.FindAllProductsFunc != nil {
		return m.FindAllProductsFunc()
	}
	return nil, nil
}
func (m *MockProductService) FindProductsById(id int) (models.Product, error) {
	if m.FindProductsByIdFunc != nil {
		return m.FindProductsByIdFunc(id)
	}
	return models.Product{}, nil
}
func (m *MockProductService) UpdateProduct(id int, p models.Product) (models.Product, error) {
	if m.UpdateProductFunc != nil {
		return m.UpdateProductFunc(id, p)
	}
	return models.Product{}, nil
}
func (m *MockProductService) DeleteProduct(id int) error {
	if m.DeleteProductFunc != nil {
		return m.DeleteProductFunc(id)
	}
	return nil
}

func TestCreateProduct(t *testing.T) {
	const validProductJSON = `{
    "product_code": "A1000",
    "description": "Coca Cola 2L",
    "width": 5,
    "height": 20,
    "length": 12,
    "net_weight": 1.9,
    "expiration_rate": 15,
    "recommended_freezing_temperature": 10,
    "freezing_rate": 10,
    "product_type_id": 83,
    "seller_id": 7
}`

	t.Run("create_ok", func(t *testing.T) {
		expected := models.Product{
			ID:             1,
			ProductCode:    "A1000",
			Description:    "Coca Cola 2L",
			Width:          5,
			Height:         20,
			Length:         12,
			NetWeight:      1.9,
			ExpirationRate: 15,
			Temperature:    10,
			FreezingRate:   10,
			ProductTypeID:  83,
		}
		mockService := &MockProductService{
			CreateProductFunc: func(p models.Product) (models.Product, error) {
				return expected, nil
			},
		}
		hd := handler.NewProductHandler(mockService)

		request := httptest.NewRequest("POST", "/products", strings.NewReader(validProductJSON))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()

		hd.CreateProduct(response, request)

		require.Equal(t, http.StatusCreated, response.Code)

		var respBody map[string]interface{}
		require.NoError(t, json.Unmarshal(response.Body.Bytes(), &respBody))
		data, ok := respBody["data"].(map[string]interface{})
		// cambiar texto
		require.True(t, ok, "la respuesta debe tener 'data' como objeto")

		require.Equal(t, float64(1), data["id"])
		require.Equal(t, "A1000", data["product_code"])
		require.Equal(t, "Coca Cola 2L", data["description"])
	})

	t.Run("create_fail", func(t *testing.T) {
		expected := models.Product{
			ID:          2,
			Description: "Coca Cola 2L",
			Width:       5,
			Height:      20,
			Length:      12,
			NetWeight:   1.9,
		}
		mockService := &MockProductService{
			CreateProductFunc: func(p models.Product) (models.Product, error) {
				return expected, nil
			},
		}
		hd := handler.NewProductHandler(mockService)
		request := httptest.NewRequest("POST", "/products", strings.NewReader(validProductJSON))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()

		hd.CreateProduct(response, request)

		var respBody map[string]interface{}
		require.NoError(t, json.Unmarshal(response.Body.Bytes(), &respBody))
		data, ok := respBody["data"].(map[string]interface{})
		require.True(t, ok, "la respuesta debe tener 'data' como objeto")

		require.Equal(t, float64(2), data["id"])
		require.Equal(t, "Coca Cola 2L", data["description"])

	})

	t.Run("create_conflict", func(t *testing.T) {
		mockService := &MockProductService{
			CreateProductFunc: func(p models.Product) (models.Product, error) {
				return models.Product{}, pkgErrors.WrapErrConflict("product", "product_code", "A1000")
			},
		}
		hd := handler.NewProductHandler(mockService)

		request := httptest.NewRequest("POST", "/products", strings.NewReader(validProductJSON))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()

		hd.CreateProduct(response, request)

		require.Equal(t, http.StatusConflict, response.Code)
		require.Contains(t, response.Body.String(), "product_code A1000 already exists")
	})

}
