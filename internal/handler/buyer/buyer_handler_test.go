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
			expectedStatusCode: http.StatusUnprocessableEntity,
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
			expectedStatusCode: http.StatusConflict,
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
		{
			name:        "create_invalid_json",
			requestBody: "invalid json string", // ← JSON malformado
			setupMock: func(mockService *buyer.MockBuyerService) {
			},
			expectedStatusCode: http.StatusBadRequest,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				body := rr.Body.String()
				assert.Contains(t, body, "bad request")
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

func TestBuyerHandler_GetAll(t *testing.T) {
	tests := []struct {
		name               string
		setupMock          func(*buyer.MockBuyerService)
		expectedStatusCode int
		validateResponse   func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "get_all_buyers_success",
			setupMock: func(mockService *buyer.MockBuyerService) {
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
				mockService.On("GetAll").Return(buyers, nil)
			},
			expectedStatusCode: http.StatusOK,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
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
			},
		},
		{
			name: "get_all_buyers_empty_list",
			setupMock: func(mockService *buyer.MockBuyerService) {
				var emptyList []*models.Buyer
				mockService.On("GetAll").Return(emptyList, nil)
			},
			expectedStatusCode: http.StatusOK,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				var response models.SuccessResponse
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				assert.NoError(t, err)

				buyersData, ok := response.Data.([]interface{})
				assert.True(t, ok, "Response data should be an array")
				assert.Len(t, buyersData, 0, "Should return empty array")
			},
		},
		{
			name: "get_all_buyers_service_error",
			setupMock: func(mockService *buyer.MockBuyerService) {
				mockService.On("GetAll").Return(([]*models.Buyer)(nil), errors.ErrGeneral)
			},
			expectedStatusCode: http.StatusInternalServerError,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				body := rr.Body.String()
				assert.Contains(t, body, "Internal Server Error")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockService := new(buyer.MockBuyerService)
			handler := NewBuyerHandler(mockService)
			tt.setupMock(mockService)

			req := httptest.NewRequest(http.MethodGet, "/api/v1/buyers", nil)
			rr := httptest.NewRecorder()

			// Act
			handlerFunc := handler.GetAll()
			handlerFunc(rr, req)

			// Assert
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			tt.validateResponse(t, rr)
			mockService.AssertExpectations(t)
		})
	}
}

func TestBuyerHandler_Update(t *testing.T) {
	tests := []struct {
		name               string
		urlParam           string
		requestBody        interface{}
		setupMock          func(*buyer.MockBuyerService)
		expectedStatusCode int
		validateResponse   func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:     "update_ok",
			urlParam: "1",
			requestBody: models.BuyerUpdateDTO{
				CardNumberId: stringPtr("1"),
				FirstName:    stringPtr("Pepito"),
				LastName:     stringPtr("Perez"),
			},
			setupMock: func(mockService *buyer.MockBuyerService) {
				updatedBuyer := &models.Buyer{
					Id:           1,
					CardNumberId: "1",
					FirstName:    "Pepito",
					LastName:     "Perez",
				}
				mockService.On("PatchUpdate", 1, mock.MatchedBy(func(dto *models.BuyerUpdateDTO) bool {
					return dto != nil &&
						dto.CardNumberId != nil && *dto.CardNumberId == "1" &&
						dto.FirstName != nil && *dto.FirstName == "Pepito" &&
						dto.LastName != nil && *dto.LastName == "Perez"
				})).Return(updatedBuyer, nil)
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
			name:     "update_buyer_partial_card_number_only",
			urlParam: "1",
			requestBody: models.BuyerUpdateDTO{
				CardNumberId: stringPtr("1"),
			},
			setupMock: func(mockService *buyer.MockBuyerService) {
				updatedBuyer := &models.Buyer{
					Id:           1,
					CardNumberId: "1",
					FirstName:    "Pepito",
					LastName:     "Perez",
				}
				mockService.On("PatchUpdate", 1, mock.MatchedBy(func(dto *models.BuyerUpdateDTO) bool {
					return dto != nil &&
						dto.CardNumberId != nil && *dto.CardNumberId == "1" &&
						dto.FirstName == nil &&
						dto.LastName == nil
				})).Return(updatedBuyer, nil)
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
			name:     "update_buyer_partial_names_only",
			urlParam: "1",
			requestBody: models.BuyerUpdateDTO{
				FirstName: stringPtr("pepito"),
				LastName:  stringPtr("perez"),
			},
			setupMock: func(mockService *buyer.MockBuyerService) {
				updatedBuyer := &models.Buyer{
					Id:           1,
					CardNumberId: "1",
					FirstName:    "pepito",
					LastName:     "perez",
				}
				mockService.On("PatchUpdate", 1, mock.MatchedBy(func(dto *models.BuyerUpdateDTO) bool {
					return dto != nil &&
						dto.CardNumberId == nil &&
						dto.FirstName != nil && *dto.FirstName == "pepito" &&
						dto.LastName != nil && *dto.LastName == "perez"
				})).Return(updatedBuyer, nil)
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
				assert.Equal(t, "pepito", buyerData["first_name"])
				assert.Equal(t, "perez", buyerData["last_name"])
			},
		},
		{
			name:     "update_non_existent",
			urlParam: "999",
			requestBody: models.BuyerUpdateDTO{
				CardNumberId: stringPtr("1"),
				FirstName:    stringPtr("pepito"),
				LastName:     stringPtr("perez"),
			},
			setupMock: func(mockService *buyer.MockBuyerService) {
				notFoundErr := errors.WrapErrNotFound("buyer", "id", 999)
				mockService.On("PatchUpdate", 999, mock.MatchedBy(func(dto *models.BuyerUpdateDTO) bool {
					return dto != nil &&
						dto.CardNumberId != nil && *dto.CardNumberId == "1" &&
						dto.FirstName != nil && *dto.FirstName == "pepito" &&
						dto.LastName != nil && *dto.LastName == "perez"
				})).Return((*models.Buyer)(nil), notFoundErr)
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
			name:     "invalid_id_format",
			urlParam: "abc",
			requestBody: models.BuyerUpdateDTO{
				CardNumberId: stringPtr("123456789"),
			},
			setupMock: func(mockService *buyer.MockBuyerService) {
			},
			expectedStatusCode: http.StatusBadRequest,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				assert.Contains(t, rr.Body.String(), "bad request")
			},
		},
		{
			name:     "negative_id",
			urlParam: "-1",
			requestBody: models.BuyerUpdateDTO{
				CardNumberId: stringPtr("1"),
			},
			setupMock: func(mockService *buyer.MockBuyerService) {
			},
			expectedStatusCode: http.StatusBadRequest,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				assert.Contains(t, rr.Body.String(), "bad request")
			},
		},
		{
			name:     "zero_id",
			urlParam: "0",
			requestBody: models.BuyerUpdateDTO{
				CardNumberId: stringPtr("1"),
			},
			setupMock: func(mockService *buyer.MockBuyerService) {
			},
			expectedStatusCode: http.StatusBadRequest,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				assert.Contains(t, rr.Body.String(), "bad request")
			},
		},
		{
			name:     "validation_error_empty_fields",
			urlParam: "1",
			requestBody: models.BuyerUpdateDTO{
				CardNumberId: stringPtr(""),
				FirstName:    stringPtr(""),
				LastName:     stringPtr(""),
			},
			setupMock: func(mockService *buyer.MockBuyerService) {
			},
			expectedStatusCode: http.StatusUnprocessableEntity,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				var errorResponse map[string]interface{}
				err := json.Unmarshal(rr.Body.Bytes(), &errorResponse)
				assert.NoError(t, err)
				message, exists := errorResponse["message"]
				assert.True(t, exists)
				assert.Contains(t, message, "unprocessable entity")
				assert.Contains(t, message, "CardNumberId")
				assert.Contains(t, message, "FirstName")
				assert.Contains(t, message, "LastName")
				assert.Contains(t, message, "min")
			},
		},
		{
			name:        "update_invalid_json",
			urlParam:    "1",
			requestBody: "invalid json string",
			setupMock: func(mockService *buyer.MockBuyerService) {
			},
			expectedStatusCode: http.StatusBadRequest,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				body := rr.Body.String()
				assert.Contains(t, body, "bad request")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockService := new(buyer.MockBuyerService)
			handler := NewBuyerHandler(mockService)
			tt.setupMock(mockService)

			var reqBody []byte
			var err error
			if str, ok := tt.requestBody.(string); ok {
				reqBody = []byte(str)
			} else {
				reqBody, err = json.Marshal(tt.requestBody)
				assert.NoError(t, err)
			}

			req := httptest.NewRequest(http.MethodPatch, "/api/v1/buyers/"+tt.urlParam, bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tt.urlParam)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			rr := httptest.NewRecorder()

			// Act
			handlerFunc := handler.Update()
			handlerFunc(rr, req)

			// Assert
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			tt.validateResponse(t, rr)
			mockService.AssertExpectations(t)
		})
	}
}

func stringPtr(s string) *string {
	return &s
}

func TestBuyerHandler_Delete(t *testing.T) {
	tests := []struct {
		name               string
		urlParam           string
		setupMock          func(*buyer.MockBuyerService)
		expectedStatusCode int
		validateResponse   func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:     "delete_ok",
			urlParam: "1",
			setupMock: func(mockService *buyer.MockBuyerService) {
				mockService.On("Delete", 1).Return(nil)
			},
			expectedStatusCode: http.StatusNoContent,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				assert.Empty(t, rr.Body.String())
			},
		},
		{
			name:     "delete_non_existent",
			urlParam: "999",
			setupMock: func(mockService *buyer.MockBuyerService) {
				notFoundErr := errors.WrapErrNotFound("buyer", "id", 999)
				mockService.On("Delete", 999).Return(notFoundErr)
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
			name:     "invalid_id_format",
			urlParam: "abc",
			setupMock: func(mockService *buyer.MockBuyerService) {
			},
			expectedStatusCode: http.StatusBadRequest,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				assert.Contains(t, rr.Body.String(), "bad request")
			},
		},
		{
			name:     "negative_id",
			urlParam: "-1",
			setupMock: func(mockService *buyer.MockBuyerService) {
			},
			expectedStatusCode: http.StatusBadRequest,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				assert.Contains(t, rr.Body.String(), "bad request")
			},
		},
		{
			name:     "zero_id",
			urlParam: "0",
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

			req := httptest.NewRequest(http.MethodDelete, "/api/v1/buyers/"+tt.urlParam, nil)

			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tt.urlParam)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			rr := httptest.NewRecorder()

			// Act
			handlerFunc := handler.Delete()
			handlerFunc(rr, req)

			// Assert
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			tt.validateResponse(t, rr)
			mockService.AssertExpectations(t)
		})
	}
}

func TestBuyerHandler_GetAllOrByIdWithOrderCount(t *testing.T) {
	tests := []struct {
		name               string
		queryParam         string
		setupMock          func(*buyer.MockBuyerService)
		expectedStatusCode int
		validateResponse   func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:       "get_all_buyers_with_order_count_success",
			queryParam: "",
			setupMock: func(mockService *buyer.MockBuyerService) {
				expectedBuyers := []*models.BuyerWithOrderCount{
					{
						Id:                  1,
						CardNumberId:        "1",
						FirstName:           "Pepito",
						LastName:            "Perez",
						PurchaseOrdersCount: 5,
					},
					{
						Id:                  2,
						CardNumberId:        "2",
						FirstName:           "Juan",
						LastName:            "Lopez",
						PurchaseOrdersCount: 0,
					},
				}
				mockService.On("GetAllWithOrderCount").Return(expectedBuyers, nil)
			},
			expectedStatusCode: http.StatusOK,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				var response models.SuccessResponse
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				assert.NoError(t, err)

				buyersData, ok := response.Data.([]interface{})
				assert.True(t, ok, "Response data should be an array")
				assert.Len(t, buyersData, 2)
				firstBuyer := buyersData[0].(map[string]interface{})
				assert.Equal(t, float64(1), firstBuyer["id"])
				assert.Equal(t, "1", firstBuyer["card_number_id"])
				assert.Equal(t, "Pepito", firstBuyer["first_name"])
				assert.Equal(t, "Perez", firstBuyer["last_name"])
				assert.Equal(t, float64(5), firstBuyer["purchase_orders_count"])
				secondBuyer := buyersData[1].(map[string]interface{})
				assert.Equal(t, float64(2), secondBuyer["id"])
				assert.Equal(t, float64(0), secondBuyer["purchase_orders_count"])
			},
		},
		{
			name:       "get_buyer_by_id_with_order_count_success",
			queryParam: "?id=1",
			setupMock: func(mockService *buyer.MockBuyerService) {
				expectedBuyer := &models.BuyerWithOrderCount{
					Id:                  1,
					CardNumberId:        "1",
					FirstName:           "Pepito",
					LastName:            "Perez",
					PurchaseOrdersCount: 5,
				}
				mockService.On("GetByIdWithOrderCount", 1).Return(expectedBuyer, nil)
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
				assert.Equal(t, float64(5), buyerData["purchase_orders_count"])
			},
		},
		{
			name:       "get_buyer_by_id_with_zero_orders",
			queryParam: "?id=2",
			setupMock: func(mockService *buyer.MockBuyerService) {
				expectedBuyer := &models.BuyerWithOrderCount{
					Id:                  2,
					CardNumberId:        "2",
					FirstName:           "Ana",
					LastName:            "Garcia",
					PurchaseOrdersCount: 0,
				}
				mockService.On("GetByIdWithOrderCount", 2).Return(expectedBuyer, nil)
			},
			expectedStatusCode: http.StatusOK,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				var response models.SuccessResponse
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				assert.NoError(t, err)

				buyerData, ok := response.Data.(map[string]interface{})
				assert.True(t, ok, "Response data should be an object")

				assert.Equal(t, float64(2), buyerData["id"])
				assert.Equal(t, float64(0), buyerData["purchase_orders_count"]) // Validar cero
			},
		},
		{
			name:       "get_buyer_by_id_non_existent",
			queryParam: "?id=999",
			setupMock: func(mockService *buyer.MockBuyerService) {
				notFoundErr := errors.WrapErrNotFound("buyer", "id", 999)
				mockService.On("GetByIdWithOrderCount", 999).Return((*models.BuyerWithOrderCount)(nil), notFoundErr)
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
			name:       "invalid_id_query_parameter_format",
			queryParam: "?id=abc",
			setupMock: func(mockService *buyer.MockBuyerService) {
			},
			expectedStatusCode: http.StatusBadRequest,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				body := rr.Body.String()
				assert.Contains(t, body, "bad request")
			},
		},
		{
			name:       "negative_id_query_parameter",
			queryParam: "?id=-1",
			setupMock: func(mockService *buyer.MockBuyerService) {
			},
			expectedStatusCode: http.StatusBadRequest,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				body := rr.Body.String()
				assert.Contains(t, body, "bad request")
			},
		},
		{
			name:       "zero_id_query_parameter",
			queryParam: "?id=0",
			setupMock: func(mockService *buyer.MockBuyerService) {
			},
			expectedStatusCode: http.StatusBadRequest,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				body := rr.Body.String()
				assert.Contains(t, body, "bad request")
			},
		},
		{
			name:       "get_all_buyers_service_error",
			queryParam: "",
			setupMock: func(mockService *buyer.MockBuyerService) {
				mockService.On("GetAllWithOrderCount").Return(([]*models.BuyerWithOrderCount)(nil), errors.ErrGeneral)
			},
			expectedStatusCode: http.StatusInternalServerError,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				body := rr.Body.String()
				assert.Contains(t, body, "internal server error")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockService := new(buyer.MockBuyerService)
			handler := NewBuyerHandler(mockService)
			tt.setupMock(mockService)

			req := httptest.NewRequest(http.MethodGet, "/api/v1/buyers/reportPurchaseOrders"+tt.queryParam, nil)

			rr := httptest.NewRecorder()

			// Act
			handlerFunc := handler.GetAllOrByIdWithOrderCount()
			handlerFunc(rr, req)

			// Assert
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			tt.validateResponse(t, rr)
			mockService.AssertExpectations(t)
		})
	}
}
