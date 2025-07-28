package purchase_order

import (
	"ProyectoFinal/mocks/purchase_order"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPurchaseOrderHandler_Create(t *testing.T) {
	tests := []struct {
		name               string
		requestBody        interface{}
		setupMock          func(mockService *purchase_order.MockPurchaseOrderService)
		expectedStatusCode int
		validateResponse   func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "create_ok",
			requestBody: models.PurchaseOrderCreateDTO{
				OrderNumber:   "ORD123",
				OrderDate:     "2023-10-05",
				TrackingCode:  "TRK123",
				BuyerId:       1,
				CarrierId:     2,
				OrderStatusId: 3,
				WarehouseId:   4,
				OrderDetails: []models.OrderDetailsCreateDTO{
					{
						CleanlinessStatus: "good",
						Quantity:          5,
						Temperature:       4.5,
						ProductRecordId:   101,
					},
				},
			},
			setupMock: func(mockService *purchase_order.MockPurchaseOrderService) {
				expectedPurchaseOrder := &models.PurchaseOrder{
					Id:            1,
					OrderNumber:   "ORD123",
					OrderDate:     "2023-10-05",
					TrackingCode:  "TRK123",
					BuyerId:       1,
					CarrierId:     2,
					OrderStatusId: 3,
					WarehouseId:   4,
					OrderDetails: []models.OrderDetails{
						{
							CleanlinessStatus: "good",
							Quantity:          5,
							Temperature:       4.5,
							ProductRecordId:   101,
						},
					},
				}
				mockService.On("Create", mock.MatchedBy(func(po *models.PurchaseOrder) bool {
					if po.OrderNumber != "ORD123" ||
						po.OrderDate != "2023-10-05" ||
						po.TrackingCode != "TRK123" ||
						po.BuyerId != 1 ||
						po.CarrierId != 2 ||
						po.OrderStatusId != 3 ||
						po.WarehouseId != 4 {
						return false
					}
					if len(po.OrderDetails) != 1 {
						return false
					}
					od := po.OrderDetails[0]
					return od.CleanlinessStatus == "good" &&
						od.Quantity == 5 &&
						od.Temperature == 4.5 &&
						od.ProductRecordId == 101
				})).Return(expectedPurchaseOrder, nil)
			},
			expectedStatusCode: http.StatusCreated,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				var response models.SuccessResponse
				err := json.Unmarshal(rr.Body.Bytes(), &response)
				assert.NoError(t, err)
				doc, ok := response.Data.(map[string]interface{})
				require.True(t, ok, "response data should be a map")
				assert.Equal(t, float64(1), doc["id"])
				assert.Equal(t, "ORD123", doc["order_number"])
				assert.Equal(t, "2023-10-05", doc["order_date"])
				assert.Equal(t, "TRK123", doc["tracking_code"])
				assert.Equal(t, float64(1), doc["buyer"])
				assert.Equal(t, float64(2), doc["carrier"])
				assert.Equal(t, float64(3), doc["order_status"])
				assert.Equal(t, float64(4), doc["warehouse"])
				orderDetails, ok := doc["order_details"].([]interface{})
				require.True(t, ok, "order_details should be an array")
				require.Len(t, orderDetails, 1)
				od, ok := orderDetails[0].(map[string]interface{})
				require.True(t, ok, "each order_detail should be a map")
				assert.Equal(t, "good", od["cleanliness_status"])
				assert.Equal(t, float64(5), od["quantity"])
				assert.Equal(t, 4.5, od["temperature"])
				assert.Equal(t, float64(101), od["product_record_id"])
			},
		},
		{
			name: "create_fail",
			requestBody: models.PurchaseOrderCreateDTO{
				OrderDate:     "2023-10-05",
				TrackingCode:  "TRK123",
				BuyerId:       1,
				CarrierId:     2,
				OrderStatusId: 3,
				WarehouseId:   4,
				OrderDetails: []models.OrderDetailsCreateDTO{
					{
						CleanlinessStatus: "good",
						Quantity:          5,
						Temperature:       4.5,
						ProductRecordId:   101,
					},
				},
			},
			setupMock: func(mockService *purchase_order.MockPurchaseOrderService) {
			},
			expectedStatusCode: http.StatusUnprocessableEntity,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				body := rr.Body.String()
				assert.Contains(t, body, "OrderNumber")
				assert.Contains(t, body, "required")
			},
		},
		{
			name:        "create_fail - All fields missing",
			requestBody: models.PurchaseOrderCreateDTO{},
			setupMock: func(mockService *purchase_order.MockPurchaseOrderService) {
			},
			expectedStatusCode: http.StatusUnprocessableEntity,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				body := rr.Body.String()
				assert.Contains(t, body, "OrderNumber")
				assert.Contains(t, body, "OrderDate")
				assert.Contains(t, body, "TrackingCode")
				assert.Contains(t, body, "BuyerId")
				assert.Contains(t, body, "CarrierId")
				assert.Contains(t, body, "OrderStatusId")
				assert.Contains(t, body, "WarehouseId")
				assert.Contains(t, body, "required")
			},
		},
		{
			name: "create_conflict",
			requestBody: models.PurchaseOrderCreateDTO{
				OrderNumber:   "ORD123",
				OrderDate:     "2023-10-05",
				TrackingCode:  "TRK123",
				BuyerId:       1,
				CarrierId:     2,
				OrderStatusId: 3,
				WarehouseId:   4,
				OrderDetails: []models.OrderDetailsCreateDTO{
					{
						CleanlinessStatus: "good",
						Quantity:          5,
						Temperature:       4.5,
						ProductRecordId:   101,
					},
				},
			},
			setupMock: func(mockService *purchase_order.MockPurchaseOrderService) {
				conflictErr := errors.WrapErrConflict("purchase_order", "order_number", "ORD123")
				mockService.On("Create", mock.AnythingOfType("*models.PurchaseOrder")).
					Return((*models.PurchaseOrder)(nil), conflictErr)
			},
			expectedStatusCode: http.StatusConflict,
			validateResponse: func(t *testing.T, rr *httptest.ResponseRecorder) {
				var errorResponse map[string]interface{}
				err := json.Unmarshal(rr.Body.Bytes(), &errorResponse)
				assert.NoError(t, err)

				message, exists := errorResponse["message"]
				assert.True(t, exists)
				assert.Contains(t, message, "conflict")
				assert.Contains(t, message, "purchase_order")
				assert.Contains(t, message, "order_number")
				assert.Contains(t, message, "ORD123")
				assert.Contains(t, message, "already exists")
			},
		},
		{
			name:        "create_invalid_json",
			requestBody: "invalid json string",
			setupMock: func(mockService *purchase_order.MockPurchaseOrderService) {
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
			// Arrange:
			mockService := new(purchase_order.MockPurchaseOrderService)
			handler := NewpPurchaseOrderHandler(mockService)
			tt.setupMock(mockService)
			var requestBody []byte
			var err error
			if str, ok := tt.requestBody.(string); ok {
				requestBody = []byte(str)
			} else {
				requestBody, err = json.Marshal(tt.requestBody)
				require.NoError(t, err)
			}

			req := httptest.NewRequest(http.MethodPost, "/api/v1/purchaseOrders", bytes.NewBuffer(requestBody))
			req.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()

			// Act:
			handlerFunc := handler.Create()
			handlerFunc(rr, req)

			// Assert:
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			if tt.validateResponse != nil {
				tt.validateResponse(t, rr)
			}
			mockService.AssertExpectations(t)
		})
	}
}
