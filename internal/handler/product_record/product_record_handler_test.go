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

func TestGetProductRecordsCount(t *testing.T) {

	t.Run("get_product_records_by_id_ok", func(t *testing.T) {
		expectedReport := models.ReportProductData{
			ProductID:    1,
			Description:  "Test Product",
			RecordsCount: 5,
		}

		productID := 1
		mockService := &mocks.MockProductRecordService{}
		mockService.On("GetRecordsProduct", &productID).Return(expectedReport, nil)
		hd := product_record.NewProductRecordHandler(mockService)

		request := httptest.NewRequest("GET", "/products/reportRecords?id=1", nil)
		response := httptest.NewRecorder()

		hd.GetProductRecordsCount(response, request)

		require.Equal(t, http.StatusOK, response.Code)
		require.Equal(t, "application/json", response.Header().Get("Content-Type"))

		var resp struct {
			Data models.ReportProductData `json:"data"`
		}

		err := json.Unmarshal(response.Body.Bytes(), &resp)
		require.NoError(t, err)
		require.Equal(t, expectedReport, resp.Data)
		mockService.AssertExpectations(t)
	})

	t.Run("get_all_product_records_ok", func(t *testing.T) {
		expectedReports := []models.ReportProductData{
			{ProductID: 1, Description: "Product 1", RecordsCount: 3},
			{ProductID: 2, Description: "Product 2", RecordsCount: 5},
		}

		mockService := &mocks.MockProductRecordService{}
		mockService.On("GetRecordsProductAll").Return(expectedReports, nil)
		hd := product_record.NewProductRecordHandler(mockService)

		request := httptest.NewRequest("GET", "/products/reportRecords", nil)
		response := httptest.NewRecorder()

		hd.GetProductRecordsCount(response, request)

		require.Equal(t, http.StatusOK, response.Code)
		require.Equal(t, "application/json", response.Header().Get("Content-Type"))

		var resp struct {
			Data []models.ReportProductData `json:"data"`
		}

		err := json.Unmarshal(response.Body.Bytes(), &resp)
		require.NoError(t, err)
		require.Equal(t, expectedReports, resp.Data)
		mockService.AssertExpectations(t)
	})

	t.Run("get_all_product_records_service_error", func(t *testing.T) {
		mockService := &mocks.MockProductRecordService{}
		mockService.On("GetRecordsProductAll").Return([]models.ReportProductData{}, pkgErrors.ErrNotFound)
		hd := product_record.NewProductRecordHandler(mockService)

		request := httptest.NewRequest("GET", "/products/reportRecords", nil)
		response := httptest.NewRecorder()

		hd.GetProductRecordsCount(response, request)

		require.Equal(t, http.StatusNotFound, response.Code)
		require.Equal(t, "application/json", response.Header().Get("Content-Type"))
		require.Contains(t, response.Body.String(), pkgErrors.ErrNotFound.Error())
		mockService.AssertExpectations(t)
	})

	t.Run("get_product_records_invalid_id", func(t *testing.T) {
		mockService := &mocks.MockProductRecordService{}
		hd := product_record.NewProductRecordHandler(mockService)

		request := httptest.NewRequest("GET", "/products/reportRecords?id=abc", nil)
		response := httptest.NewRecorder()

		hd.GetProductRecordsCount(response, request)

		require.Equal(t, http.StatusBadRequest, response.Code)
		require.Equal(t, "application/json", response.Header().Get("Content-Type"))
		require.Contains(t, response.Body.String(), "invalid param")
		mockService.AssertExpectations(t)
	})

	t.Run("get_product_records_service_error", func(t *testing.T) {
		productID := 1
		mockService := &mocks.MockProductRecordService{}
		mockService.On("GetRecordsProduct", &productID).Return(models.ReportProductData{}, pkgErrors.ErrNotFound)
		hd := product_record.NewProductRecordHandler(mockService)

		request := httptest.NewRequest("GET", "/products/reportRecords?id=1", nil)
		response := httptest.NewRecorder()

		hd.GetProductRecordsCount(response, request)

		require.Equal(t, http.StatusNotFound, response.Code)
		require.Contains(t, response.Body.String(), pkgErrors.ErrNotFound.Error())
		mockService.AssertExpectations(t)
	})
}
