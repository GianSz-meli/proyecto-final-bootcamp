package warehouse

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	customErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"

	mocks "ProyectoFinal/mocks/warehouse"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestGetAllWarehouses_ValidRequest_Success(t *testing.T) {
	mockService := new(mocks.MockWarehouseService)
	handler := NewWarehouseHandler(mockService)

	warehouses := []models.Warehouse{
		{
			ID:                 1,
			WarehouseCode:      "1234567890",
			Address:            "Address 1",
			Telephone:          "1234567890",
			MinimumCapacity:    10,
			MinimumTemperature: 10,
			LocalityId:         nil,
			Locality:           nil,
		},
		{
			ID:                 2,
			WarehouseCode:      "1234567890",
			Address:            "Address 2",
			Telephone:          "1234567890",
			MinimumCapacity:    10,
			MinimumTemperature: 10,
			LocalityId:         nil,
			Locality:           nil,
		},
	}
	expectedResponseBody := `
	{
		"data": [
			{
				"id":1,
				"warehouse_code":"1234567890",
				"address":"Address 1",
				"telephone":"1234567890",
				"minimum_capacity":10,
				"minimum_temperature":10,
				"locality":null
			},
			{
				"id":2,
				"warehouse_code":"1234567890",
				"address":"Address 2",
				"telephone":"1234567890",
				"minimum_capacity":10,
				"minimum_temperature":10,
				"locality":null
			}
		]
	}`

	mockService.On("GetAllWarehouses").Return(warehouses, nil)

	req := httptest.NewRequest(http.MethodGet, "/warehouses", nil)
	w := httptest.NewRecorder()

	handler.GetAllWarehouses(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	require.JSONEq(t, expectedResponseBody, w.Body.String())
	mockService.AssertExpectations(t)
}

func TestGetAllWarehouses_ServiceError_InternalServerError(t *testing.T) {
	mockService := new(mocks.MockWarehouseService)
	handler := NewWarehouseHandler(mockService)

	// Mock an unexpected service error (not a predefined error like NotFound, Conflict, etc.)
	unexpectedError := errors.New("database connection failed")
	expectedResponseBody := `{"status":"Internal Server Error","message":"Internal Server Error"}`

	mockService.On("GetAllWarehouses").Return([]models.Warehouse{}, unexpectedError)

	req := httptest.NewRequest(http.MethodGet, "/warehouses", nil)
	w := httptest.NewRecorder()

	handler.GetAllWarehouses(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code)
	require.Equal(t, expectedResponseBody, w.Body.String())
	mockService.AssertExpectations(t)
}

func TestGetWarehouseById_ValidId_Success(t *testing.T) {
	mockService := new(mocks.MockWarehouseService)
	handler := NewWarehouseHandler(mockService)

	warehouse := models.Warehouse{
		ID:                 1,
		WarehouseCode:      "WH001",
		Address:            "Address 1",
		Telephone:          "1234567890",
		MinimumCapacity:    100,
		MinimumTemperature: -10.5,
		LocalityId:         nil,
		Locality:           nil,
	}
	expectedResponse := `{
		"data": {
			"id": 1,
			"warehouse_code": "WH001",
			"address": "Address 1",
			"telephone": "1234567890",
			"minimum_capacity": 100,
			"minimum_temperature": -10.5,
			"locality": null
		}
	}`

	mockService.On("GetWarehouseById", 1).Return(warehouse, nil)

	req := httptest.NewRequest(http.MethodGet, "/warehouses/1", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()

	handler.GetWarehouseById(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	require.JSONEq(t, expectedResponse, w.Body.String())
	mockService.AssertExpectations(t)
}

func TestGetWarehouseById_NonExistingId_NotFound(t *testing.T) {
	mockService := new(mocks.MockWarehouseService)
	handler := NewWarehouseHandler(mockService)

	nonExistingId := 999
	expectedError := customErrors.WrapErrNotFound("warehouse", "id", nonExistingId)
	expectedResponseBody := `{"status":"Not Found","message":"not found : warehouse with id 999 not found"}`

	mockService.On("GetWarehouseById", nonExistingId).Return(models.Warehouse{}, expectedError)

	req := httptest.NewRequest(http.MethodGet, "/warehouses/999", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "999")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()

	handler.GetWarehouseById(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	require.Equal(t, expectedResponseBody, w.Body.String())
	mockService.AssertExpectations(t)
}

func TestGetWarehouseById_InvalidId_BadRequest(t *testing.T) {
	mockService := new(mocks.MockWarehouseService)
	handler := NewWarehouseHandler(mockService)

	expectedResponseBody := `{"status":"Bad Request","message":"bad request : id must be a number"}`

	req := httptest.NewRequest(http.MethodGet, "/warehouses/abc", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "abc")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()

	handler.GetWarehouseById(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
	require.Equal(t, expectedResponseBody, w.Body.String())
}

func TestCreateWarehouse_ValidRequest_Success(t *testing.T) {
	mockService := new(mocks.MockWarehouseService)
	handler := NewWarehouseHandler(mockService)
	createRequest := models.CreateWarehouseRequest{
		WarehouseCode:      "WH001",
		Address:            "Address 1",
		Telephone:          "1234567890",
		MinimumCapacity:    &[]int{100}[0],
		MinimumTemperature: &[]float64{-10.5}[0],
		LocalityId:         nil,
	}
	expectedWarehouse := models.Warehouse{
		ID:                 1,
		WarehouseCode:      "WH001",
		Address:            "Address 1",
		Telephone:          "1234567890",
		MinimumCapacity:    100,
		MinimumTemperature: -10.5,
		LocalityId:         nil,
		Locality:           nil,
	}
	expectedResponse := `{
		"data": {
			"id": 1,
			"warehouse_code": "WH001",
			"address": "Address 1",
			"telephone": "1234567890",
			"minimum_capacity": 100,
			"minimum_temperature": -10.5,
			"locality_id": null
		}
	}`

	body, _ := json.Marshal(createRequest)
	mockService.On("CreateWarehouse", createRequest.DocToModel()).Return(expectedWarehouse, nil)

	req := httptest.NewRequest(http.MethodPost, "/warehouses", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.CreateWarehouse(w, req)

	require.Equal(t, http.StatusCreated, w.Code)
	require.JSONEq(t, expectedResponse, w.Body.String())
	mockService.AssertExpectations(t)
}

func TestCreateWarehouse_InvalidRequest_UnprocessableEntity(t *testing.T) {
	mockService := new(mocks.MockWarehouseService)
	handler := NewWarehouseHandler(mockService)

	testCases := []struct {
		name                 string
		requestBody          models.CreateWarehouseRequest
		expectedResponseBody string
	}{
		{
			name: "warehouse_code is required",
			requestBody: models.CreateWarehouseRequest{
				Address:            "Address 1",
				Telephone:          "1234567890",
				MinimumCapacity:    &[]int{100}[0],
				MinimumTemperature: &[]float64{-10.5}[0],
				LocalityId:         nil,
			},
			expectedResponseBody: `
			{ 
				"status":"Unprocessable Entity",
				"message": "unprocessable entity : Key: 'CreateWarehouseRequest.WarehouseCode' Error:Field validation for 'WarehouseCode' failed on the 'required' tag"
			}`,
		},
		{
			name: "warehouse_code minimum length validation",
			requestBody: models.CreateWarehouseRequest{
				WarehouseCode:      "",
				Address:            "Address 1",
				Telephone:          "1234567890",
				MinimumCapacity:    &[]int{100}[0],
				MinimumTemperature: &[]float64{-10.5}[0],
				LocalityId:         nil,
			},
			expectedResponseBody: `
			{ 
				"status":"Unprocessable Entity",
				"message": "unprocessable entity : Key: 'CreateWarehouseRequest.WarehouseCode' Error:Field validation for 'WarehouseCode' failed on the 'required' tag"
			}`,
		},
		{
			name: "address is required",
			requestBody: models.CreateWarehouseRequest{
				WarehouseCode:      "WH001",
				Telephone:          "1234567890",
				MinimumCapacity:    &[]int{100}[0],
				MinimumTemperature: &[]float64{-10.5}[0],
				LocalityId:         nil,
			},
			expectedResponseBody: `
			{ 
				"status":"Unprocessable Entity",
				"message": "unprocessable entity : Key: 'CreateWarehouseRequest.Address' Error:Field validation for 'Address' failed on the 'required' tag"
			}`,
		},
		{
			name: "address minimum length validation",
			requestBody: models.CreateWarehouseRequest{
				WarehouseCode:      "WH001",
				Address:            "",
				Telephone:          "1234567890",
				MinimumCapacity:    &[]int{100}[0],
				MinimumTemperature: &[]float64{-10.5}[0],
				LocalityId:         nil,
			},
			expectedResponseBody: `
			{ 
				"status":"Unprocessable Entity",
				"message": "unprocessable entity : Key: 'CreateWarehouseRequest.Address' Error:Field validation for 'Address' failed on the 'required' tag"
			}`,
		},
		{
			name: "telephone is required",
			requestBody: models.CreateWarehouseRequest{
				WarehouseCode:      "WH001",
				Address:            "Address 1",
				MinimumCapacity:    &[]int{100}[0],
				MinimumTemperature: &[]float64{-10.5}[0],
				LocalityId:         nil,
			},
			expectedResponseBody: `
			{ 
				"status":"Unprocessable Entity",
				"message": "unprocessable entity : Key: 'CreateWarehouseRequest.Telephone' Error:Field validation for 'Telephone' failed on the 'required' tag"
			}`,
		},
		{
			name: "telephone must be numeric",
			requestBody: models.CreateWarehouseRequest{
				WarehouseCode:      "WH001",
				Address:            "Address 1",
				Telephone:          "12345abc90",
				MinimumCapacity:    &[]int{100}[0],
				MinimumTemperature: &[]float64{-10.5}[0],
				LocalityId:         nil,
			},
			expectedResponseBody: `
			{ 
				"status":"Unprocessable Entity",
				"message": "unprocessable entity : Key: 'CreateWarehouseRequest.Telephone' Error:Field validation for 'Telephone' failed on the 'numeric' tag"
			}`,
		},
		{
			name: "telephone minimum length validation",
			requestBody: models.CreateWarehouseRequest{
				WarehouseCode:      "WH001",
				Address:            "Address 1",
				Telephone:          "123456",
				MinimumCapacity:    &[]int{100}[0],
				MinimumTemperature: &[]float64{-10.5}[0],
				LocalityId:         nil,
			},
			expectedResponseBody: `
			{ 
				"status":"Unprocessable Entity",
				"message": "unprocessable entity : Key: 'CreateWarehouseRequest.Telephone' Error:Field validation for 'Telephone' failed on the 'min' tag"
			}`,
		},
		{
			name: "minimum_capacity is required",
			requestBody: models.CreateWarehouseRequest{
				WarehouseCode:      "WH001",
				Address:            "Address 1",
				Telephone:          "1234567890",
				MinimumTemperature: &[]float64{-10.5}[0],
				LocalityId:         nil,
			},
			expectedResponseBody: `
			{ 
				"status":"Unprocessable Entity",
				"message": "unprocessable entity : Key: 'CreateWarehouseRequest.MinimumCapacity' Error:Field validation for 'MinimumCapacity' failed on the 'required' tag"
			}`,
		},
		{
			name: "minimum_capacity must be greater than 0",
			requestBody: models.CreateWarehouseRequest{
				WarehouseCode:      "WH001",
				Address:            "Address 1",
				Telephone:          "1234567890",
				MinimumCapacity:    &[]int{-1}[0],
				MinimumTemperature: &[]float64{-10.5}[0],
				LocalityId:         nil,
			},
			expectedResponseBody: `
			{ 
				"status":"Unprocessable Entity",
				"message": "unprocessable entity : Key: 'CreateWarehouseRequest.MinimumCapacity' Error:Field validation for 'MinimumCapacity' failed on the 'min' tag"
			}`,
		},
		{
			name: "minimum_temperature is required",
			requestBody: models.CreateWarehouseRequest{
				WarehouseCode:   "WH001",
				Address:         "Address 1",
				Telephone:       "1234567890",
				MinimumCapacity: &[]int{100}[0],
				LocalityId:      nil,
			},
			expectedResponseBody: `
			{ 
				"status":"Unprocessable Entity",
				"message": "unprocessable entity : Key: 'CreateWarehouseRequest.MinimumTemperature' Error:Field validation for 'MinimumTemperature' failed on the 'required' tag"
			}`,
		},
		{
			name: "locality_id must be greater than 0",
			requestBody: models.CreateWarehouseRequest{
				WarehouseCode:      "WH001",
				Address:            "Address 1",
				Telephone:          "1234567890",
				MinimumCapacity:    &[]int{100}[0],
				MinimumTemperature: &[]float64{-10.5}[0],
				LocalityId:         &[]int{0}[0],
			},
			expectedResponseBody: `
			{ 
				"status":"Unprocessable Entity",
				"message": "unprocessable entity : Key: 'CreateWarehouseRequest.LocalityId' Error:Field validation for 'LocalityId' failed on the 'gt' tag"
			}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body, _ := json.Marshal(tc.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/warehouses", bytes.NewReader(body))
			w := httptest.NewRecorder()

			handler.CreateWarehouse(w, req)

			require.Equal(t, http.StatusUnprocessableEntity, w.Code)
			require.JSONEq(t, tc.expectedResponseBody, w.Body.String())
		})
	}
}

func TestCreateWarehouse_DuplicateWarehouseCode_Conflict(t *testing.T) {
	mockService := new(mocks.MockWarehouseService)
	handler := NewWarehouseHandler(mockService)

	createRequest := models.CreateWarehouseRequest{
		WarehouseCode:      "WH001",
		Address:            "Address 1",
		Telephone:          "1234567890",
		MinimumCapacity:    &[]int{100}[0],
		MinimumTemperature: &[]float64{-10.5}[0],
		LocalityId:         nil,
	}
	expectedError := customErrors.WrapErrConflict("warehouses", "warehouse_code", "WH001")
	expectedResponseBody := `{"status":"Conflict","message":"conflict : warehouses with warehouse_code WH001 already exists"}`

	mockService.On("CreateWarehouse", createRequest.DocToModel()).Return(models.Warehouse{}, expectedError)

	body, _ := json.Marshal(createRequest)
	req := httptest.NewRequest(http.MethodPost, "/warehouses", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.CreateWarehouse(w, req)

	require.Equal(t, http.StatusConflict, w.Code)
	require.Equal(t, expectedResponseBody, w.Body.String())
	mockService.AssertExpectations(t)
}

func TestCreateWarehouse_MalformedJson_BadRequest(t *testing.T) {
	mockService := new(mocks.MockWarehouseService)
	handler := NewWarehouseHandler(mockService)

	malformedJson := `{warehouse_code: "WH001", "address": "Address 1"}}`
	expectedResponseBody := `
	{ 
		"status":"Bad Request",
		"message":"bad request : it was not possible to decode json"
	}`

	req := httptest.NewRequest(http.MethodPost, "/warehouses", bytes.NewReader([]byte(malformedJson)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.CreateWarehouse(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
	require.JSONEq(t, expectedResponseBody, w.Body.String())
}

func TestUpdateWarehouse_ValidIdAndRequest_Success(t *testing.T) {
	mockService := new(mocks.MockWarehouseService)
	handler := NewWarehouseHandler(mockService)

	updateRequest := models.UpdateWarehouseRequest{
		WarehouseCode: &[]string{"WH001-UPDATED"}[0],
	}
	updatedWarehouse := models.Warehouse{
		ID:                 1,
		WarehouseCode:      "WH001-UPDATED",
		Address:            "Address 1",
		Telephone:          "1234567890",
		MinimumCapacity:    100,
		MinimumTemperature: -10.5,
		LocalityId:         nil,
		Locality:           nil,
	}
	expectedResponse := `{
		"data": {
			"id": 1,
			"warehouse_code": "WH001-UPDATED",
			"address": "Address 1",
			"telephone": "1234567890",
			"minimum_capacity": 100,
			"minimum_temperature": -10.5,
			"locality_id": null
		}
	}`

	mockService.On("UpdateWarehouse", 1, updateRequest).Return(updatedWarehouse, nil)

	body, _ := json.Marshal(updateRequest)
	req := httptest.NewRequest(http.MethodPut, "/warehouses/1", bytes.NewReader(body))
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()

	handler.UpdateWarehouse(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	require.JSONEq(t, expectedResponse, w.Body.String())
	mockService.AssertExpectations(t)
}

func TestUpdateWarehouse_NonExistingId_NotFound(t *testing.T) {
	mockService := new(mocks.MockWarehouseService)
	handler := NewWarehouseHandler(mockService)

	nonExistingId := 999
	updateRequest := models.UpdateWarehouseRequest{}
	expectedError := customErrors.WrapErrNotFound("warehouse", "id", nonExistingId)
	expectedResponseBody := `{"status":"Not Found","message":"not found : warehouse with id 999 not found"}`

	mockService.On("UpdateWarehouse", nonExistingId, updateRequest).Return(models.Warehouse{}, expectedError)

	body, _ := json.Marshal(updateRequest)
	req := httptest.NewRequest(http.MethodPut, "/warehouses/999", bytes.NewReader(body))
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "999")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()

	handler.UpdateWarehouse(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	require.Equal(t, expectedResponseBody, w.Body.String())
	mockService.AssertExpectations(t)
}

func TestUpdateWarehouse_InvalidRequest_UnprocessableEntity(t *testing.T) {
	mockService := new(mocks.MockWarehouseService)
	handler := NewWarehouseHandler(mockService)

	testCases := []struct {
		name                 string
		requestBody          models.UpdateWarehouseRequest
		expectedResponseBody string
	}{
		{
			name: "warehouse_code minimum length validation",
			requestBody: models.UpdateWarehouseRequest{
				WarehouseCode: &[]string{""}[0],
			},
			expectedResponseBody: `
			{ 
				"status":"Unprocessable Entity",
				"message": "unprocessable entity : Key: 'UpdateWarehouseRequest.WarehouseCode' Error:Field validation for 'WarehouseCode' failed on the 'min' tag"
			}`,
		},
		{
			name: "address minimum length validation",
			requestBody: models.UpdateWarehouseRequest{
				Address: &[]string{""}[0],
			},
			expectedResponseBody: `
			{ 
				"status":"Unprocessable Entity",
				"message": "unprocessable entity : Key: 'UpdateWarehouseRequest.Address' Error:Field validation for 'Address' failed on the 'min' tag"
			}`,
		},
		{
			name: "telephone must be numeric",
			requestBody: models.UpdateWarehouseRequest{
				Telephone: &[]string{"12345abc90"}[0],
			},
			expectedResponseBody: `
			{ 
				"status":"Unprocessable Entity",
				"message": "unprocessable entity : Key: 'UpdateWarehouseRequest.Telephone' Error:Field validation for 'Telephone' failed on the 'numeric' tag"
			}`,
		},
		{
			name: "telephone minimum length validation",
			requestBody: models.UpdateWarehouseRequest{
				Telephone: &[]string{"123456"}[0],
			},
			expectedResponseBody: `
			{ 
				"status":"Unprocessable Entity",
				"message": "unprocessable entity : Key: 'UpdateWarehouseRequest.Telephone' Error:Field validation for 'Telephone' failed on the 'min' tag"
			}`,
		},
		{
			name: "minimum_capacity must be greater than 0",
			requestBody: models.UpdateWarehouseRequest{
				MinimumCapacity: &[]int{-1}[0],
			},
			expectedResponseBody: `
			{ 
				"status":"Unprocessable Entity",
				"message": "unprocessable entity : Key: 'UpdateWarehouseRequest.MinimumCapacity' Error:Field validation for 'MinimumCapacity' failed on the 'min' tag"
			}`,
		},
		{
			name: "locality_id must be greater than 0",
			requestBody: models.UpdateWarehouseRequest{
				LocalityId: &[]int{0}[0],
			},
			expectedResponseBody: `
			{ 
				"status":"Unprocessable Entity",
				"message": "unprocessable entity : Key: 'UpdateWarehouseRequest.LocalityId' Error:Field validation for 'LocalityId' failed on the 'gt' tag"
			}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			body, _ := json.Marshal(tc.requestBody)
			req := httptest.NewRequest(http.MethodPut, "/warehouses/1", bytes.NewReader(body))
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", "1")
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
			w := httptest.NewRecorder()

			handler.UpdateWarehouse(w, req)

			require.Equal(t, http.StatusUnprocessableEntity, w.Code)
			require.JSONEq(t, tc.expectedResponseBody, w.Body.String())
		})
	}
}

func TestUpdateWarehouse_DuplicateWarehouseCode_Conflict(t *testing.T) {
	mockService := new(mocks.MockWarehouseService)
	handler := NewWarehouseHandler(mockService)

	updateRequest := models.UpdateWarehouseRequest{
		WarehouseCode: &[]string{"WH002"}[0],
	}
	expectedError := customErrors.WrapErrConflict("warehouses", "warehouse_code", "WH002")
	expectedResponseBody := `{"status":"Conflict","message":"conflict : warehouses with warehouse_code WH002 already exists"}`

	mockService.On("UpdateWarehouse", 1, updateRequest).Return(models.Warehouse{}, expectedError).Once()

	body, _ := json.Marshal(updateRequest)
	req := httptest.NewRequest(http.MethodPut, "/warehouses/1", bytes.NewReader(body))
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()

	handler.UpdateWarehouse(w, req)

	require.Equal(t, http.StatusConflict, w.Code)
	require.Equal(t, expectedResponseBody, w.Body.String())
	mockService.AssertExpectations(t)
}

func TestUpdateWarehouse_BadData_BadRequest(t *testing.T) {
	mockService := new(mocks.MockWarehouseService)
	handler := NewWarehouseHandler(mockService)

	tests := []struct {
		name                 string
		id                   string
		requestBodyJson      string
		expectedResponseBody string
	}{
		{
			name:                 "invalid id",
			id:                   "abc",
			requestBodyJson:      ``,
			expectedResponseBody: `{"status":"Bad Request","message":"bad request : id must be a number"}`,
		},
		{
			name:            "malformed json",
			id:              "1",
			requestBodyJson: `{warehouse_code: "WH001-UPDATED", "address": "Address Updated"}`,
			expectedResponseBody: `
			{ 
				"status":"Bad Request",
				"message":"bad request : it was not possible to decode json"
			}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPut, "/warehouses/"+test.id, bytes.NewReader([]byte(test.requestBodyJson)))

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", test.id)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
			w := httptest.NewRecorder()

			handler.UpdateWarehouse(w, req)

			require.Equal(t, http.StatusBadRequest, w.Code)
			require.JSONEq(t, test.expectedResponseBody, w.Body.String())
		})
	}
}

func TestDeleteWarehouse_ValidId_Success(t *testing.T) {
	mockService := new(mocks.MockWarehouseService)
	handler := NewWarehouseHandler(mockService)

	mockService.On("DeleteWarehouse", 1).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/warehouses/1", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()

	handler.DeleteWarehouse(w, req)

	require.Equal(t, http.StatusNoContent, w.Code)
	mockService.AssertExpectations(t)
}

func TestDeleteWarehouse_NonExistingId_NotFound(t *testing.T) {
	mockService := new(mocks.MockWarehouseService)
	handler := NewWarehouseHandler(mockService)

	nonExistingId := 999
	expectedError := customErrors.WrapErrNotFound("warehouse", "id", nonExistingId)
	expectedResponseBody := `{"status":"Not Found","message":"not found : warehouse with id 999 not found"}`

	mockService.On("DeleteWarehouse", nonExistingId).Return(expectedError)

	req := httptest.NewRequest(http.MethodDelete, "/warehouses/999", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "999")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()

	handler.DeleteWarehouse(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
	require.Equal(t, expectedResponseBody, w.Body.String())
	mockService.AssertExpectations(t)
}

func TestDeleteWarehouse_InvalidId_BadRequest(t *testing.T) {
	mockService := new(mocks.MockWarehouseService)
	handler := NewWarehouseHandler(mockService)

	expectedResponseBody := `{"status":"Bad Request","message":"bad request : id must be a number"}`

	req := httptest.NewRequest(http.MethodDelete, "/warehouses/abc", nil)

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "abc")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()

	handler.DeleteWarehouse(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
	require.Equal(t, expectedResponseBody, w.Body.String())
}
