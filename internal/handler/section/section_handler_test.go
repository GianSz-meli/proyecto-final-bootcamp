package section

import (
	"ProyectoFinal/mocks"
	"ProyectoFinal/pkg/models"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestSectionHandler_GetAll_Success(t *testing.T) {
	// Arrange
	mockService := new(mocks.MockSectionService)
	handler := NewSectionDefault(mockService)

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

	sectionsDoc := make([]models.SectionDoc, 0, len(sections))
	for _, section := range sections {
		sectionsDoc = append(sectionsDoc, section.ModelToDoc())
	}

	expectedResponse := models.SuccessResponse{
		Data: sectionsDoc,
	}
	expectedResponseBytes, _ := json.Marshal(expectedResponse)

	mockService.On("GetAll").Return(sections, nil)

	// Act
	req := httptest.NewRequest(http.MethodGet, "/sections", nil)
	w := httptest.NewRecorder()
	handler.GetAll()(w, req)

	// Assert
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, string(expectedResponseBytes), w.Body.String())
	mockService.AssertExpectations(t)
}

func TestSectionHandler_GetById_Existent_Success(t *testing.T) {
	// Arrange
	mockService := new(mocks.MockSectionService)
	handler := NewSectionDefault(mockService)

	section := models.Section{
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

	expectedResponse := models.SuccessResponse{
		Data: section.ModelToDoc(),
	}
	expectedResponseBytes, _ := json.Marshal(expectedResponse)

	mockService.On("GetById", 1).Return(section, nil)

	// Act
	req := httptest.NewRequest(http.MethodGet, "/sections/1", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	handler.GetById()(w, req)

	// Assert
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, string(expectedResponseBytes), w.Body.String())
	mockService.AssertExpectations(t)
}

func TestSectionHandler_Create_ValidRequest_Success(t *testing.T) {
	// Arrange
	mockService := new(mocks.MockSectionService)
	handler := NewSectionDefault(mockService)

	createRequest := models.CreateSectionRequest{
		SectionNumber:      "SEC001",
		CurrentTemperature: 15.5,
		MinimumTemperature: 10.0,
		CurrentCapacity:    50,
		MinimumCapacity:    20,
		MaximumCapacity:    100,
		WarehouseID:        1,
		ProductTypeID:      1,
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

	expectedResponse := models.SuccessResponse{
		Data: createdSection.ModelToDoc(),
	}
	expectedResponseBytes, _ := json.Marshal(expectedResponse)

	body, _ := json.Marshal(createRequest)
	req := httptest.NewRequest(http.MethodPost, "/sections", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	mockService.On("Create", createRequest.DocToModel()).Return(createdSection, nil)

	// Act
	w := httptest.NewRecorder()
	handler.Create()(w, req)

	// Assert
	require.Equal(t, http.StatusCreated, w.Code)
	require.Equal(t, string(expectedResponseBytes), w.Body.String())
	mockService.AssertExpectations(t)
}

func TestSectionHandler_Create_InvalidRequest_UnprocessableEntity(t *testing.T) {
	// Arrange
	mockService := new(mocks.MockSectionService)
	handler := NewSectionDefault(mockService)

	testCases := []struct {
		name        string
		requestBody models.CreateSectionRequest
	}{
		{
			name: "missing section_number",
			requestBody: models.CreateSectionRequest{
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
			name: "invalid maximum_capacity",
			requestBody: models.CreateSectionRequest{
				SectionNumber:      "SEC001",
				CurrentTemperature: 15.5,
				MinimumTemperature: 10.0,
				CurrentCapacity:    50,
				MinimumCapacity:    20,
				MaximumCapacity:    0, // Invalid: must be greater than 0
				WarehouseID:        1,
				ProductTypeID:      1,
			},
		},
		{
			name: "invalid warehouse_id",
			requestBody: models.CreateSectionRequest{
				SectionNumber:      "SEC001",
				CurrentTemperature: 15.5,
				MinimumTemperature: 10.0,
				CurrentCapacity:    50,
				MinimumCapacity:    20,
				MaximumCapacity:    100,
				WarehouseID:        0, // Invalid: must be greater than 0
				ProductTypeID:      1,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			body, _ := json.Marshal(tc.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/sections", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			// Act
			handler.Create()(w, req)

			// Assert
			require.Equal(t, http.StatusUnprocessableEntity, w.Code)
			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			require.NoError(t, err)
			require.Equal(t, "Unprocessable Entity", response["status"])
		})
	}
}

func TestSectionHandler_Update_ValidRequest_Success(t *testing.T) {
	// Arrange
	mockService := new(mocks.MockSectionService)
	handler := NewSectionDefault(mockService)

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

	expectedResponse := models.SuccessResponse{
		Data: updatedSection.ModelToDoc(),
	}
	expectedResponseBytes, _ := json.Marshal(expectedResponse)

	body, _ := json.Marshal(updateRequest)
	req := httptest.NewRequest(http.MethodPut, "/sections/1", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()

	mockService.On("UpdateWithValidation", 1, updateRequest).Return(updatedSection, nil)

	// Act
	handler.Update()(w, req)

	// Assert
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, string(expectedResponseBytes), w.Body.String())
	mockService.AssertExpectations(t)
}

func TestSectionHandler_Delete_Existent_Success(t *testing.T) {
	// Arrange
	mockService := new(mocks.MockSectionService)
	handler := NewSectionDefault(mockService)

	mockService.On("Delete", 1).Return(nil)

	// Act
	req := httptest.NewRequest(http.MethodDelete, "/sections/1", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	handler.Delete()(w, req)

	// Assert
	require.Equal(t, http.StatusNoContent, w.Code)
	mockService.AssertExpectations(t)
}

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
