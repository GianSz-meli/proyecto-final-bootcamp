package buyer

import (
	"ProyectoFinal/mocks/buyer"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
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
				assert.Contains(t, rr.Body.String(), "CardNumberId")
				body := rr.Body.String()
				assert.True(t,
					strings.Contains(body, "CardNumberId") || strings.Contains(body, "card_number_id"),
					"Response should contain field name: %s", body)
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
				conflictErr := errors.WrapErrConflict("buyer", "card_number_id", "12345678")
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
				assert.Contains(t, message, "12345678")
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
