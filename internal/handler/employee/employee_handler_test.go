package employee

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"ProyectoFinal/mocks"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestEmployeeHandler_Create_OK(t *testing.T) {
	// Arrange
	mockService := &mocks.MockEmployeeService{}
	handler := NewEmployeeHandler(mockService)

	warehouseID := 1
	employeeRequest := models.EmployeeRequest{
		CardNumberID: "12345",
		FirstName:    "John",
		LastName:     "Doe",
		WarehouseID:  warehouseID,
	}

	expectedEmployee := models.Employee{
		ID:           1,
		CardNumberID: "12345",
		FirstName:    "John",
		LastName:     "Doe",
		WarehouseID:  &warehouseID,
	}

	mockService.On("Create", mock.AnythingOfType("models.Employee")).Return(expectedEmployee, nil)

	reqBody, _ := json.Marshal(employeeRequest)
	req := httptest.NewRequest(http.MethodPost, "/employees", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	handler.Create()(w, req)

	// Assert
	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.SuccessResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	employeeDoc := response.Data.(map[string]interface{})
	assert.Equal(t, float64(1), employeeDoc["id"])
	assert.Equal(t, "12345", employeeDoc["card_number_id"])
	assert.Equal(t, "John", employeeDoc["first_name"])
	assert.Equal(t, "Doe", employeeDoc["last_name"])

	mockService.AssertExpectations(t)
}

func TestEmployeeHandler_Create_InvalidJSON(t *testing.T) {
	// Arrange
	mockService := &mocks.MockEmployeeService{}
	handler := NewEmployeeHandler(mockService)

	req := httptest.NewRequest(http.MethodPost, "/employees", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	handler.Create()(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockService.AssertExpectations(t)
}

func TestEmployeeHandler_Create_Fail(t *testing.T) {
	// Arrange
	mockService := &mocks.MockEmployeeService{}
	handler := NewEmployeeHandler(mockService)

	invalidRequest := map[string]interface{}{
		"card_number_id": "12345",
		"last_name":      "Doe",
		"warehouse_id":   1,
	}

	reqBody, _ := json.Marshal(invalidRequest)
	req := httptest.NewRequest(http.MethodPost, "/employees", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	handler.Create()(w, req)

	// Assert
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	mockService.AssertExpectations(t)
}

func TestEmployeeHandler_Create_Conflict(t *testing.T) {
	// Arrange
	mockService := &mocks.MockEmployeeService{}
	handler := NewEmployeeHandler(mockService)

	warehouseID := 1
	employeeRequest := models.EmployeeRequest{
		CardNumberID: "12345",
		FirstName:    "John",
		LastName:     "Doe",
		WarehouseID:  warehouseID,
	}

	mockService.On("Create", mock.AnythingOfType("models.Employee")).Return(models.Employee{}, errors.ErrConflict)

	reqBody, _ := json.Marshal(employeeRequest)
	req := httptest.NewRequest(http.MethodPost, "/employees", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	handler.Create()(w, req)

	// Assert
	assert.Equal(t, http.StatusConflict, w.Code)
	mockService.AssertExpectations(t)
}

func TestEmployeeHandler_GetAll(t *testing.T) {
	// Arrange
	mockService := &mocks.MockEmployeeService{}
	handler := NewEmployeeHandler(mockService)

	warehouseID1 := 1
	warehouseID2 := 2
	expectedEmployees := []models.Employee{
		{
			ID:           1,
			CardNumberID: "12345",
			FirstName:    "John",
			LastName:     "Doe",
			WarehouseID:  &warehouseID1,
		},
		{
			ID:           2,
			CardNumberID: "67890",
			FirstName:    "Jane",
			LastName:     "Smith",
			WarehouseID:  &warehouseID2,
		},
	}

	mockService.On("GetAll").Return(expectedEmployees, nil)
	req := httptest.NewRequest(http.MethodGet, "/employees", nil)
	w := httptest.NewRecorder()

	// Act
	handler.GetAll()(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response models.SuccessResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	employeesData := response.Data.([]interface{})
	assert.Len(t, employeesData, 2)

	mockService.AssertExpectations(t)
}

func TestEmployeeHandler_GetAll_ServiceError(t *testing.T) {
	// Arrange
	mockService := &mocks.MockEmployeeService{}
	handler := NewEmployeeHandler(mockService)

	mockService.On("GetAll").Return([]models.Employee{}, errors.ErrGeneral)
	req := httptest.NewRequest(http.MethodGet, "/employees", nil)
	w := httptest.NewRecorder()

	// Act
	handler.GetAll()(w, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockService.AssertExpectations(t)
}

func TestEmployeeHandler_GetById_InvalidParam(t *testing.T) {
	// Arrange
	mockService := &mocks.MockEmployeeService{}
	handler := NewEmployeeHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/employees/invalid", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "invalid")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()

	// Act
	handler.GetById()(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockService.AssertExpectations(t)
}

func TestEmployeeHandler_GetById_NonExistent(t *testing.T) {
	// Arrange
	mockService := &mocks.MockEmployeeService{}
	handler := NewEmployeeHandler(mockService)

	employeeID := 124
	mockService.On("GetById", employeeID).Return(models.Employee{}, errors.ErrNotFound)

	req := httptest.NewRequest(http.MethodGet, "/employees/124", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "124")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()

	// Act
	handler.GetById()(w, req)

	// Assert
	assert.Equal(t, http.StatusNotFound, w.Code)
	mockService.AssertExpectations(t)
}

func TestEmployeeHandler_GetById_Existent(t *testing.T) {
	// Arrange
	mockService := &mocks.MockEmployeeService{}
	handler := NewEmployeeHandler(mockService)

	employeeID := 1
	warehouseID := 1
	expectedEmployee := models.Employee{
		ID:           employeeID,
		CardNumberID: "12345",
		FirstName:    "John",
		LastName:     "Doe",
		WarehouseID:  &warehouseID,
	}

	mockService.On("GetById", employeeID).Return(expectedEmployee, nil)

	req := httptest.NewRequest(http.MethodGet, "/employees/1", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()

	// Act
	handler.GetById()(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response models.SuccessResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	employeeData := response.Data.(map[string]interface{})
	assert.Equal(t, float64(1), employeeData["id"])
	assert.Equal(t, "12345", employeeData["card_number_id"])

	mockService.AssertExpectations(t)
}

func TestEmployeeHandler_Update_OK(t *testing.T) {
	// Arrange
	mockService := &mocks.MockEmployeeService{}
	handler := NewEmployeeHandler(mockService)

	employeeID := 1
	warehouseID := 1

	updatedFirstName := "Johnny"
	updateRequest := models.EmployeeUpdateRequest{
		FirstName: &updatedFirstName,
	}

	expectedUpdatedEmployee := models.Employee{
		ID:           employeeID,
		CardNumberID: "12345",
		FirstName:    "Johnny",
		LastName:     "Doe",
		WarehouseID:  &warehouseID,
	}

	mockService.On("PatchUpdate", employeeID, &updateRequest).Return(expectedUpdatedEmployee, nil)

	reqBody, _ := json.Marshal(updateRequest)
	req := httptest.NewRequest(http.MethodPatch, "/employees/1", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()

	// Act
	handler.Update()(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)

	var response models.SuccessResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	employeeData := response.Data.(map[string]interface{})
	assert.Equal(t, "Johnny", employeeData["first_name"])

	mockService.AssertExpectations(t)
}

func TestEmployeeHandler_Update_InvalidParam(t *testing.T) {
	// Arrange
	mockService := &mocks.MockEmployeeService{}
	handler := NewEmployeeHandler(mockService)

	updatedFirstName := "Johnny"
	updateRequest := models.EmployeeUpdateRequest{
		FirstName: &updatedFirstName,
	}

	reqBody, _ := json.Marshal(updateRequest)
	req := httptest.NewRequest(http.MethodPatch, "/employees/invalid", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "invalid")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()

	// Act
	handler.Update()(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockService.AssertExpectations(t)
}

func TestEmployeeHandler_Update_InvalidJSON(t *testing.T) {
	// Arrange
	mockService := &mocks.MockEmployeeService{}
	handler := NewEmployeeHandler(mockService)

	req := httptest.NewRequest(http.MethodPatch, "/employees/1", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()

	// Act
	handler.Update()(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockService.AssertExpectations(t)
}

func TestEmployeeHandler_Update_NonExistent(t *testing.T) {
	// Arrange
	mockService := &mocks.MockEmployeeService{}
	handler := NewEmployeeHandler(mockService)

	employeeID := 124
	updatedFirstName := "Johnny"
	updateRequest := models.EmployeeUpdateRequest{
		FirstName: &updatedFirstName,
	}

	mockService.On("PatchUpdate", employeeID, &updateRequest).Return(models.Employee{}, errors.ErrNotFound)

	reqBody, _ := json.Marshal(updateRequest)
	req := httptest.NewRequest(http.MethodPatch, "/employees/124", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "124")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()

	// Act
	handler.Update()(w, req)

	// Assert
	assert.Equal(t, http.StatusNotFound, w.Code)
	mockService.AssertExpectations(t)
}

func TestEmployeeHandler_Delete_InvalidParam(t *testing.T) {
	// Arrange
	mockService := &mocks.MockEmployeeService{}
	handler := NewEmployeeHandler(mockService)

	req := httptest.NewRequest(http.MethodDelete, "/employees/invalid", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "invalid")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()

	// Act
	handler.Delete()(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockService.AssertExpectations(t)
}

func TestEmployeeHandler_Delete_NonExistent(t *testing.T) {
	// Arrange
	mockService := &mocks.MockEmployeeService{}
	handler := NewEmployeeHandler(mockService)

	employeeID := 124
	mockService.On("Delete", employeeID).Return(errors.ErrNotFound)

	req := httptest.NewRequest(http.MethodDelete, "/employees/124", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "124")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()

	// Act
	handler.Delete()(w, req)

	// Assert
	assert.Equal(t, http.StatusNotFound, w.Code)
	mockService.AssertExpectations(t)
}

func TestEmployeeHandler_Delete_OK(t *testing.T) {
	// Arrange
	mockService := &mocks.MockEmployeeService{}
	handler := NewEmployeeHandler(mockService)

	employeeID := 1
	mockService.On("Delete", employeeID).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/employees/1", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()

	// Act
	handler.Delete()(w, req)

	// Assert
	assert.Equal(t, http.StatusNoContent, w.Code)
	mockService.AssertExpectations(t)
}
