package product_batch

import (
	mocks "ProyectoFinal/mocks/product_batch"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestProductBatchHandler_Create_ValidRequest_Success(t *testing.T) {
	// Arrange
	mockService := new(mocks.MockProductBatchService)
	handler := NewProductBatchHandler(mockService)

	requestBody := models.ProductBatchCreateRequest{
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

	requestBodyBytes, _ := json.Marshal(requestBody)
	expectedResponse := models.SuccessResponse{
		Data: createdProductBatch.ModelToDoc(),
	}
	expectedResponseBytes, _ := json.Marshal(expectedResponse)

	mockService.On("Create", requestBody.CreateRequestToModel()).Return(createdProductBatch, nil)

	// Act
	req := httptest.NewRequest(http.MethodPost, "/product-batches", bytes.NewBuffer(requestBodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler.Create()(w, req)

	// Assert
	require.Equal(t, http.StatusCreated, w.Code)
	require.Equal(t, string(expectedResponseBytes), w.Body.String())
	mockService.AssertExpectations(t)
}

func TestProductBatchHandler_Create_InvalidJSON_BadRequest(t *testing.T) {
	// Arrange
	mockService := new(mocks.MockProductBatchService)
	handler := NewProductBatchHandler(mockService)

	invalidJSON := `{"batch_number": "BATCH001", "current_quantity": "invalid"}`

	// Act
	req := httptest.NewRequest(http.MethodPost, "/product-batches", bytes.NewBufferString(invalidJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler.Create()(w, req)

	// Assert
	require.Equal(t, http.StatusBadRequest, w.Code)
	mockService.AssertNotCalled(t, "Create")
}

func TestProductBatchHandler_Create_InvalidRequestData_UnprocessableEntity(t *testing.T) {
	// Arrange
	mockService := new(mocks.MockProductBatchService)
	handler := NewProductBatchHandler(mockService)

	// Missing required fields
	requestBody := models.ProductBatchCreateRequest{
		BatchNumber:        "",
		CurrentQuantity:    -1, // Invalid negative value
		CurrentTemperature: 15.5,
		DueDate:            time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
		InitialQuantity:    100,
		ManufacturingDate:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		ManufacturingHour:  25, // Invalid hour > 23
		MinimumTemperature: 10.0,
		ProductID:          1,
		SectionID:          1,
	}

	requestBodyBytes, _ := json.Marshal(requestBody)

	// Act
	req := httptest.NewRequest(http.MethodPost, "/product-batches", bytes.NewBuffer(requestBodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler.Create()(w, req)

	// Assert
	require.Equal(t, http.StatusUnprocessableEntity, w.Code)
	mockService.AssertNotCalled(t, "Create")
}

func TestProductBatchHandler_Create_ServiceError_InternalServerError(t *testing.T) {
	// Arrange
	mockService := new(mocks.MockProductBatchService)
	handler := NewProductBatchHandler(mockService)

	requestBody := models.ProductBatchCreateRequest{
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

	requestBodyBytes, _ := json.Marshal(requestBody)

	mockService.On("Create", requestBody.CreateRequestToModel()).Return(models.ProductBatch{}, errors.ErrGeneral)

	// Act
	req := httptest.NewRequest(http.MethodPost, "/product-batches", bytes.NewBuffer(requestBodyBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler.Create()(w, req)

	// Assert
	require.Equal(t, http.StatusInternalServerError, w.Code)
	mockService.AssertExpectations(t)
}

func TestProductBatchHandler_GetProductCountBySection_ValidSectionID_Success(t *testing.T) {
	// Arrange
	mockService := new(mocks.MockProductBatchService)
	handler := NewProductBatchHandler(mockService)

	sectionID := 1
	reports := []models.SectionProductReport{
		{
			SectionID:     1,
			SectionNumber: "SEC001",
			ProductsCount: 5,
		},
	}

	expectedResponse := models.SuccessResponse{
		Data: reports,
	}
	expectedResponseBytes, _ := json.Marshal(expectedResponse)

	mockService.On("GetProductCountBySection", &sectionID).Return(reports, nil)

	// Act
	req := httptest.NewRequest(http.MethodGet, "/sections/1/product-count", nil)
	chiCtx := chi.NewRouteContext()
	chiCtx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
	w := httptest.NewRecorder()
	handler.GetProductCountBySection()(w, req)

	// Assert
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, string(expectedResponseBytes), w.Body.String())
	mockService.AssertExpectations(t)
}

func TestProductBatchHandler_GetProductCountBySection_InvalidSectionID_BadRequest(t *testing.T) {
	// Arrange
	mockService := new(mocks.MockProductBatchService)
	handler := NewProductBatchHandler(mockService)

	// Act
	req := httptest.NewRequest(http.MethodGet, "/sections/invalid/product-count", nil)
	chiCtx := chi.NewRouteContext()
	chiCtx.URLParams.Add("id", "invalid")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
	w := httptest.NewRecorder()
	handler.GetProductCountBySection()(w, req)

	// Assert
	require.Equal(t, http.StatusBadRequest, w.Code)
	mockService.AssertNotCalled(t, "GetProductCountBySection")
}

func TestProductBatchHandler_GetProductCountBySection_ServiceError_InternalServerError(t *testing.T) {
	// Arrange
	mockService := new(mocks.MockProductBatchService)
	handler := NewProductBatchHandler(mockService)

	sectionID := 1

	mockService.On("GetProductCountBySection", &sectionID).Return(nil, errors.ErrGeneral)

	// Act
	req := httptest.NewRequest(http.MethodGet, "/sections/1/product-count", nil)
	chiCtx := chi.NewRouteContext()
	chiCtx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiCtx))
	w := httptest.NewRecorder()
	handler.GetProductCountBySection()(w, req)

	// Assert
	require.Equal(t, http.StatusInternalServerError, w.Code)
	mockService.AssertExpectations(t)
}

func TestProductBatchHandler_GetProductCountByAllSections_Success(t *testing.T) {
	// Arrange
	mockService := new(mocks.MockProductBatchService)
	handler := NewProductBatchHandler(mockService)

	reports := []models.SectionProductReport{
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
	}

	expectedResponse := models.SuccessResponse{
		Data: reports,
	}
	expectedResponseBytes, _ := json.Marshal(expectedResponse)

	mockService.On("GetProductCountBySection", (*int)(nil)).Return(reports, nil)

	// Act
	req := httptest.NewRequest(http.MethodGet, "/sections/product-count", nil)
	w := httptest.NewRecorder()
	handler.GetProductCountByAllSections()(w, req)

	// Assert
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, string(expectedResponseBytes), w.Body.String())
	mockService.AssertExpectations(t)
}

func TestProductBatchHandler_GetProductCountByAllSections_ServiceError_InternalServerError(t *testing.T) {
	// Arrange
	mockService := new(mocks.MockProductBatchService)
	handler := NewProductBatchHandler(mockService)

	mockService.On("GetProductCountBySection", (*int)(nil)).Return(nil, errors.ErrGeneral)

	// Act
	req := httptest.NewRequest(http.MethodGet, "/sections/product-count", nil)
	w := httptest.NewRecorder()
	handler.GetProductCountByAllSections()(w, req)

	// Assert
	require.Equal(t, http.StatusInternalServerError, w.Code)
	mockService.AssertExpectations(t)
}
