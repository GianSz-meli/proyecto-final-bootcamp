package handler

import (
	"ProyectoFinal/mocks/inbound_order"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestInboundOrderHandler_Create(t *testing.T) {
	tests := []struct {
		name               string
		requestBody        interface{}
		setupMock          func(mockService *inbound_order.MockInboundOrderService)
		expectedStatusCode int
		validateResponse   func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "create_success",
			requestBody: models.InboundOrderRequest{
				OrderDate:      "2023-10-15",
				OrderNumber:    "ORD001",
				EmployeeID:     1,
				ProductBatchID: 1,
				WarehouseID:    1,
			},
			setupMock: func(mockService *inbound_order.MockInboundOrderService) {
				expectedInboundOrder := models.InboundOrder{
					ID:             1,
					OrderDate:      time.Date(2023, 10, 15, 0, 0, 0, 0, time.UTC),
					OrderNumber:    "ORD001",
					EmployeeID:     1,
					ProductBatchID: 1,
					WarehouseID:    1,
				}
				mockService.On("Create", mock.MatchedBy(func(inboundOrder models.InboundOrder) bool {
					return inboundOrder.OrderNumber == "ORD001" &&
						inboundOrder.EmployeeID == 1 &&
						inboundOrder.ProductBatchID == 1 &&
						inboundOrder.WarehouseID == 1
				})).Return(expectedInboundOrder, nil)
			},
			expectedStatusCode: http.StatusCreated,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				var response models.SuccessResponse
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				require.NoError(t, err)

				responseData, ok := response.Data.(map[string]interface{})
				require.True(t, ok)
				assert.Equal(t, float64(1), responseData["id"])
				assert.Equal(t, "ORD001", responseData["order_number"])
				assert.Equal(t, float64(1), responseData["employee_id"])
				assert.Equal(t, float64(1), responseData["product_batch_id"])
				assert.Equal(t, float64(1), responseData["warehouse_id"])
			},
		},
		{
			name:        "create_invalid_json",
			requestBody: "invalid json",
			setupMock: func(mockService *inbound_order.MockInboundOrderService) {

			},
			expectedStatusCode: http.StatusBadRequest,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				body := rr.Body.String()
				assert.Contains(t, body, "bad request")
			},
		},
		{
			name: "create_validation_error",
			requestBody: models.InboundOrderRequest{
				OrderDate:      "",
				OrderNumber:    "",
				EmployeeID:     0,
				ProductBatchID: 0,
				WarehouseID:    0,
			},
			setupMock: func(mockService *inbound_order.MockInboundOrderService) {

			},
			expectedStatusCode: http.StatusUnprocessableEntity,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				body := rr.Body.String()
				assert.Contains(t, body, "unprocessable entity")
			},
		},
		{
			name: "create_service_error",
			requestBody: models.InboundOrderRequest{
				OrderDate:      "2023-10-15",
				OrderNumber:    "ORD001",
				EmployeeID:     999,
				ProductBatchID: 1,
				WarehouseID:    1,
			},
			setupMock: func(mockService *inbound_order.MockInboundOrderService) {
				conflictErr := errors.WrapErrConflict("employees", "id", 999)
				mockService.On("Create", mock.MatchedBy(func(inboundOrder models.InboundOrder) bool {
					return inboundOrder.OrderNumber == "ORD001" &&
						inboundOrder.EmployeeID == 999
				})).Return(models.InboundOrder{}, conflictErr)
			},
			expectedStatusCode: http.StatusConflict,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				body := rr.Body.String()
				assert.Contains(t, body, "conflict")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockService := new(inbound_order.MockInboundOrderService)
			handler := NewInboundOrderHandler(mockService)
			tt.setupMock(mockService)

			var requestBody []byte
			var err error

			if strBody, ok := tt.requestBody.(string); ok {
				requestBody = []byte(strBody)
			} else {
				requestBody, err = json.Marshal(tt.requestBody)
				require.NoError(t, err)
			}

			req := httptest.NewRequest(http.MethodPost, "/api/v1/inbound-orders", bytes.NewBuffer(requestBody))
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

func TestInboundOrderHandler_GetEmployeeInboundOrdersReport(t *testing.T) {
	tests := []struct {
		name               string
		queryParams        string
		setupMock          func(mockService *inbound_order.MockInboundOrderService)
		expectedStatusCode int
		validateResponse   func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name:        "get_report_by_employee_id_success",
			queryParams: "id=1",
			setupMock: func(mockService *inbound_order.MockInboundOrderService) {
				expectedReport := models.EmployeeInboundOrdersReport{
					ID:                 1,
					CardNumberID:       "EMP001",
					FirstName:          "John",
					LastName:           "Doe",
					WarehouseID:        1,
					InboundOrdersCount: 5,
				}
				mockService.On("GetEmployeeInboundOrdersReportByEmployeeId", 1).Return(expectedReport, nil)
			},
			expectedStatusCode: http.StatusOK,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				var response models.SuccessResponse
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				require.NoError(t, err)

				responseData, ok := response.Data.([]interface{})
				require.True(t, ok)
				require.Len(t, responseData, 1)

				reportData, ok := responseData[0].(map[string]interface{})
				require.True(t, ok)
				assert.Equal(t, float64(1), reportData["id"])
				assert.Equal(t, "EMP001", reportData["card_number_id"])
				assert.Equal(t, "John", reportData["first_name"])
				assert.Equal(t, "Doe", reportData["last_name"])
				assert.Equal(t, float64(1), reportData["warehouse_id"])
				assert.Equal(t, float64(5), reportData["inbound_orders_count"])
			},
		},
		{
			name:        "get_report_by_employee_id_not_found",
			queryParams: "id=999",
			setupMock: func(mockService *inbound_order.MockInboundOrderService) {
				notFoundErr := errors.WrapErrNotFound("employee", "id", 999)
				mockService.On("GetEmployeeInboundOrdersReportByEmployeeId", 999).Return(models.EmployeeInboundOrdersReport{}, notFoundErr)
			},
			expectedStatusCode: http.StatusNotFound,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				body := rr.Body.String()
				assert.Contains(t, body, "not found")
			},
		},
		{
			name:        "get_report_by_employee_id_invalid_id",
			queryParams: "id=invalid",
			setupMock: func(mockService *inbound_order.MockInboundOrderService) {
				// No expectations set - should not be called
			},
			expectedStatusCode: http.StatusBadRequest,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				body := rr.Body.String()
				assert.Contains(t, body, "bad request")
			},
		},
		{
			name:        "get_all_reports_success",
			queryParams: "",
			setupMock: func(mockService *inbound_order.MockInboundOrderService) {
				expectedReports := []models.EmployeeInboundOrdersReport{
					{
						ID:                 1,
						CardNumberID:       "EMP001",
						FirstName:          "John",
						LastName:           "Doe",
						WarehouseID:        1,
						InboundOrdersCount: 5,
					},
					{
						ID:                 2,
						CardNumberID:       "EMP002",
						FirstName:          "Jane",
						LastName:           "Smith",
						WarehouseID:        2,
						InboundOrdersCount: 3,
					},
				}
				mockService.On("GetEmployeeInboundOrdersReportAll").Return(expectedReports, nil)
			},
			expectedStatusCode: http.StatusOK,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				var response models.SuccessResponse
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				require.NoError(t, err)

				responseData, ok := response.Data.([]interface{})
				require.True(t, ok)
				require.Len(t, responseData, 2)

				firstReport, ok := responseData[0].(map[string]interface{})
				require.True(t, ok)
				assert.Equal(t, float64(1), firstReport["id"])
				assert.Equal(t, "EMP001", firstReport["card_number_id"])

				secondReport, ok := responseData[1].(map[string]interface{})
				require.True(t, ok)
				assert.Equal(t, float64(2), secondReport["id"])
				assert.Equal(t, "EMP002", secondReport["card_number_id"])
			},
		},
		{
			name:        "get_all_reports_service_error",
			queryParams: "",
			setupMock: func(mockService *inbound_order.MockInboundOrderService) {
				mockService.On("GetEmployeeInboundOrdersReportAll").Return(nil, errors.ErrGeneral)
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
			mockService := new(inbound_order.MockInboundOrderService)
			handler := NewInboundOrderHandler(mockService)
			tt.setupMock(mockService)

			reqURL := "/api/v1/inbound-orders/report"
			if tt.queryParams != "" {
				reqURL += "?" + tt.queryParams
			}

			req := httptest.NewRequest(http.MethodGet, reqURL, nil)

			// Parse query parameters
			if tt.queryParams != "" {
				u, err := url.Parse(reqURL)
				require.NoError(t, err)
				req.URL = u
			}

			rr := httptest.NewRecorder()

			// Act
			handlerFunc := handler.GetEmployeeInboundOrdersReport()
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
