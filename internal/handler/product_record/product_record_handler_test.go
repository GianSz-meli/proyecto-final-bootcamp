package product_record_test

import (
	"ProyectoFinal/internal/handler/product_record"
	"ProyectoFinal/mocks"
	pkgErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateProductRecord(t *testing.T) {
	const validProductRecordJSON = `{
		"last_update_date": "2024-06-06",
		"purchase_price": 50.50,
		"sale_price": 60.60,
		"product_id": 1234
	}`

	const invalidProductRecordJSON = `{
		"purchase_price": 50.50,
		"sale_price": 60.60,
		"product_id": 1234
	}`

	inputProductRecord := models.ProductRecord{
		LastUpdateDate: "2024-06-06",
		PurchasePrice:  50.50,
		SalePrice:      60.60,
		ProductID:      1234,
	}

	returnedProductRecord := models.ProductRecord{
		ID:             1,
		LastUpdateDate: "2024-06-06",
		PurchasePrice:  50.50,
		SalePrice:      60.60,
		ProductID:      1234,
	}

	t.Run("create_product_record_ok", func(t *testing.T) {
		mockService := &mocks.MockProductRecordService{}
		mockService.On("CreateProductRecord", inputProductRecord).Return(returnedProductRecord, nil)
		hd := product_record.NewProductRecordHandler(mockService)

		request := httptest.NewRequest("POST", "/productRecords", strings.NewReader(validProductRecordJSON))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		hd.CreateProductRecord(response, request)

		require.Equal(t, http.StatusCreated, response.Code)
		require.Equal(t, "application/json", response.Header().Get("Content-Type"))

		var resp struct {
			Data models.ProductRecordDoc `json:"data"`
		}

		err := json.Unmarshal(response.Body.Bytes(), &resp)
		require.NoError(t, err)
		require.Equal(t, returnedProductRecord, resp.Data.DocToModel())
		mockService.AssertExpectations(t)

	})

	t.Run("create_product_record_fail", func(t *testing.T) {
		mockService := &mocks.MockProductRecordService{}
		hd := product_record.NewProductRecordHandler(mockService)

		request := httptest.NewRequest("POST", "/productRecords", strings.NewReader(invalidProductRecordJSON))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()

		hd.CreateProductRecord(response, request)

		require.Equal(t, http.StatusUnprocessableEntity, response.Code)
		require.Equal(t, "application/json", response.Header().Get("Content-Type"))
		mockService.AssertNotCalled(t, "CreateProductRecord", mock.Anything)

		mockService.AssertExpectations(t)
	})

	t.Run("product_record_conflict", func(t *testing.T) {

		mockService := &mocks.MockProductRecordService{}
		mockService.On("CreateProductRecord", mock.Anything).Return(models.ProductRecord{}, pkgErrors.ErrConflict)
		hd := product_record.NewProductRecordHandler(mockService)

		request := httptest.NewRequest("POST", "/productRecords", strings.NewReader(validProductRecordJSON))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()

		hd.CreateProductRecord(response, request)

		require.Equal(t, http.StatusConflict, response.Code)
		require.Equal(t, "application/json", response.Header().Get("Content-Type"))
		mockService.AssertExpectations(t)

	})

	t.Run("invalid_json_product_record", func(t *testing.T) {

		const malformedJSON = `{
			"last_update_date": "2024-06-06",
			"purchase_price": 50.50,
			"sale_price": 60.60,
			"product_id": 1234,`

		mockService := &mocks.MockProductRecordService{}
		hd := product_record.NewProductRecordHandler(mockService)

		request := httptest.NewRequest("POST", "/productRecords", strings.NewReader(malformedJSON))
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()

		hd.CreateProductRecord(response, request)

		require.Equal(t, http.StatusBadRequest, response.Code)
		require.Equal(t, "application/json", response.Header().Get("Content-Type"))

		mockService.AssertExpectations(t)
	})
}
