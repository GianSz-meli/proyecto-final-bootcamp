package buyer

import (
	"ProyectoFinal/mocks/buyer"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBuyerHandler_Create(t *testing.T) {
	tests := []struct {
		name               string
		requestBody        interface{}
		setupMock          func(mockService *buyer.MockBuyerService)
		expectedStatusCode int
		expectedResponse   interface{}
		validateResponse   func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "create_ok",
			requestBody: models.BuyerCreateDTO{
				CardNumberId: "1",
				FirstName:    "Pepito",
				LastName:     "Perez",
			},
			setupMock: func(mockService *buyer.MockBuyerService) {
				expectedBuyer := &models.Buyer{
					Id:           1,
					CardNumberId: "1",
					FirstName:    "Pepito",
					LastName:     "Perez",
				}
				mockService.On("Create", mock.MatchedBy(func(buyer *models.Buyer) bool {
					return buyer.CardNumberId == "1" &&
						buyer.FirstName == "Pepito" &&
						buyer.LastName == "Perez"
				})).Return(expectedBuyer, nil)
			},
			expectedStatusCode: http.StatusCreated,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				var response models.SuccessResponse
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				assert.NoError(t, err)

				responseData, ok := response.Data.(map[string]interface{})
				assert.True(t, ok)
				assert.Equal(t, float64(1), responseData["id"])
				assert.Equal(t, "1", responseData["card_number_id"])
				assert.Equal(t, "Pepito", responseData["first_name"])
				assert.Equal(t, "Perez", responseData["last_name"])
			},
		},
		{
			name: "create_fail",
			requestBody: models.BuyerCreateDTO{
				FirstName: "Pepito",
				LastName:  "Perez",
			},
			setupMock: func(mockService *buyer.MockBuyerService) {
			},
			expectedStatusCode: http.StatusUnprocessableEntity, // 422
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				body := rr.Body.String()
				assert.Contains(t, body, "CardNumberId")
				assert.Contains(t, body, "required")
			},
		},
		{
			name:        "create_fail - All fields missing",
			requestBody: models.BuyerCreateDTO{},
			setupMock: func(mockService *buyer.MockBuyerService) {
			},
			expectedStatusCode: http.StatusUnprocessableEntity,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				body := rr.Body.String()
				assert.Contains(t, body, "CardNumberId")
				assert.Contains(t, body, "FirstName")
				assert.Contains(t, body, "LastName")
				assert.Contains(t, body, "required")
			},
		},
		{
			name: "create_conflict",
			requestBody: models.BuyerCreateDTO{
				CardNumberId: "1",
				FirstName:    "Pepito",
				LastName:     "Perez",
			},
			setupMock: func(mockService *buyer.MockBuyerService) {
				conflictErr := errors.WrapErrConflict("buyer", "card_number_id", "1")
				mockService.On("Create", mock.AnythingOfType("*models.Buyer")).Return((*models.Buyer)(nil), conflictErr)
			},
			expectedStatusCode: http.StatusConflict, // 409
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				var errorResponse map[string]interface{}
				err := json.Unmarshal(rr.Body.Bytes(), &errorResponse)
				assert.NoError(t, err)

				message, exists := errorResponse["message"]
				assert.True(t, exists)
				assert.Contains(t, message, "conflict")
				assert.Contains(t, message, "buyer")
				assert.Contains(t, message, "card_number_id")
				assert.Contains(t, message, "1")
				assert.Contains(t, message, "already exists")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockService := new(buyer.MockBuyerService)
			handler := NewBuyerHandler(mockService)
			tt.setupMock(mockService)

			var requestBody []byte
			var err error

			if strBody, ok := tt.requestBody.(string); ok {
				requestBody = []byte(strBody)
			} else {
				requestBody, err = json.Marshal(tt.requestBody)
				assert.NoError(t, err)
			}

			req := httptest.NewRequest(http.MethodPost, "/api/v1/buyers", bytes.NewBuffer(requestBody))
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()

			// Act
			handlerFunc := handler.Create()
			handlerFunc(rr, req)

			// Assert
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			if tt.validateResponse != nil {
				tt.validateResponse(t, rr)
			}
			mockService.AssertExpectations(t)
		})
	}
}

func TestBuyerHandler_GetById(t *testing.T) {
	tests := []struct {
		name               string
		urlParam           string
		setupMock          func(*buyer.MockBuyerService)
		expectedStatusCode int
		validateResponse   func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:     "find_by_id_existent",
			urlParam: "1",
			setupMock: func(mockService *buyer.MockBuyerService) {
				expectedBuyer := &models.Buyer{
					Id:           1,
					CardNumberId: "1",
					FirstName:    "Pepito",
					LastName:     "Perez",
				}
				mockService.On("GetById", 1).Return(expectedBuyer, nil)
			},
			expectedStatusCode: http.StatusOK,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				var response models.SuccessResponse
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				assert.NoError(t, err)

				buyerData, ok := response.Data.(map[string]interface{})
				assert.True(t, ok, "Response data should be an object")

				assert.Equal(t, float64(1), buyerData["id"])
				assert.Equal(t, "1", buyerData["card_number_id"])
				assert.Equal(t, "Pepito", buyerData["first_name"])
				assert.Equal(t, "Perez", buyerData["last_name"])
			},
		},
		{
			name:     "find_by_id_non_existent",
			urlParam: "999",
			setupMock: func(mockService *buyer.MockBuyerService) {
				notFoundErr := errors.WrapErrNotFound("buyer", "id", 999)
				mockService.On("GetById", 999).Return((*models.Buyer)(nil), notFoundErr)
			},
			expectedStatusCode: http.StatusNotFound,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				var errorResponse map[string]interface{}
				err := json.Unmarshal(rr.Body.Bytes(), &errorResponse)
				assert.NoError(t, err)

				message, exists := errorResponse["message"]
				assert.True(t, exists)
				assert.Contains(t, message, "not found")
				assert.Contains(t, message, "buyer")
				assert.Contains(t, message, "id")
				assert.Contains(t, message, "999")
			},
		},
		{
			name:     "invalid id format",
			urlParam: "abc",
			setupMock: func(mockService *buyer.MockBuyerService) {
			},
			expectedStatusCode: http.StatusBadRequest,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				assert.Contains(t, rr.Body.String(), "bad request")
			},
		},
		{
			name:     "negative id",
			urlParam: "-1",
			setupMock: func(mockService *buyer.MockBuyerService) {
			},
			expectedStatusCode: http.StatusBadRequest,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				assert.Contains(t, rr.Body.String(), "bad request")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockService := new(buyer.MockBuyerService)
			handler := NewBuyerHandler(mockService)
			tt.setupMock(mockService)
			req := httptest.NewRequest(http.MethodGet, "/api/v1/buyers"+tt.urlParam, nil)
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tt.urlParam)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			rr := httptest.NewRecorder()

			// Act
			handlerFunc := handler.GetById()
			handlerFunc(rr, req)

			// Assert
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			tt.validateResponse(t, rr)
			mockService.AssertExpectations(t)
		})
	}
}

func TestBuyerHandler_GetAll_Success(t *testing.T) {
	// Arrange
	buyers := []*models.Buyer{
		{
			Id:           1,
			CardNumberId: "1",
			FirstName:    "Pepito",
			LastName:     "Perez",
		},
		{
			Id:           2,
			CardNumberId: "2",
			FirstName:    "Cosme",
			LastName:     "Fulanito",
		},
		{
			Id:           3,
			CardNumberId: "3",
			FirstName:    "Fulano",
			LastName:     "De Tal",
		},
	}

	mockService := new(buyer.MockBuyerService)
	handler := NewBuyerHandler(mockService)
	mockService.On("GetAll").Return(buyers, nil)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/buyers", nil)
	rr := httptest.NewRecorder()

	// Act
	handlerFunc := handler.GetAll()
	handlerFunc(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	var response models.SuccessResponse
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)
	buyersData, ok := response.Data.([]interface{})
	assert.True(t, ok, "Response data should be an array")
	assert.Len(t, buyersData, 3, "Should return 3 buyers")
	firstBuyer, ok := buyersData[0].(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, float64(1), firstBuyer["id"])
	assert.Equal(t, "1", firstBuyer["card_number_id"])
	assert.Equal(t, "Pepito", firstBuyer["first_name"])
	assert.Equal(t, "Perez", firstBuyer["last_name"])
	secondBuyer, ok := buyersData[1].(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, float64(2), secondBuyer["id"])
	assert.Equal(t, "2", secondBuyer["card_number_id"])
	assert.Equal(t, "Cosme", secondBuyer["first_name"])
	assert.Equal(t, "Fulanito", secondBuyer["last_name"])
	thirdBuyer, ok := buyersData[2].(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, float64(3), thirdBuyer["id"])
	assert.Equal(t, "3", thirdBuyer["card_number_id"])
	assert.Equal(t, "Fulano", thirdBuyer["first_name"])
	assert.Equal(t, "De Tal", thirdBuyer["last_name"])
	mockService.AssertExpectations(t)
}
