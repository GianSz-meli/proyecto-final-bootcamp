package purchase_order

import (
	"ProyectoFinal/mocks/purchase_order"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPurchaseOrderService_Create(t *testing.T) {
	tests := []struct {
		name           string
		purchaseOrder  *models.PurchaseOrder
		setupMock      func(*purchase_order.MockPurchaseOrderRepository)
		validateResult func(*testing.T, *models.PurchaseOrder, error)
	}{
		{
			name: "create_ok",
			purchaseOrder: &models.PurchaseOrder{
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
			},
			setupMock: func(mockRepo *purchase_order.MockPurchaseOrderRepository) {
				expectedPO := &models.PurchaseOrder{
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
							Id:                1,
							CleanlinessStatus: "good",
							Quantity:          5,
							Temperature:       4.5,
							ProductRecordId:   101,
						},
					},
				}
				mockRepo.On("Create", mock.MatchedBy(func(po *models.PurchaseOrder) bool {
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
				})).Return(expectedPO, nil)
			},
			validateResult: func(t *testing.T, result *models.PurchaseOrder, err error) {
				require.NoError(t, err)
				require.NotNil(t, result)
				assert.Equal(t, 1, result.Id)
				assert.Equal(t, "ORD123", result.OrderNumber)
				assert.Equal(t, "2023-10-05", result.OrderDate)
				assert.Equal(t, "TRK123", result.TrackingCode)
				assert.Equal(t, 1, result.BuyerId)
				assert.Equal(t, 2, result.CarrierId)
				assert.Equal(t, 3, result.OrderStatusId)
				assert.Equal(t, 4, result.WarehouseId)
				require.Len(t, result.OrderDetails, 1)
				od := result.OrderDetails[0]
				assert.Equal(t, "good", od.CleanlinessStatus)
				assert.Equal(t, 5, od.Quantity)
				assert.Equal(t, 4.5, od.Temperature)
				assert.Equal(t, 101, od.ProductRecordId)
			},
		},
		{
			name: "create_conflict",
			purchaseOrder: &models.PurchaseOrder{
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
			},
			setupMock: func(mockRepo *purchase_order.MockPurchaseOrderRepository) {
				conflictErr := errors.WrapErrConflict("purchase_order", "order_number", "ORD123")
				mockRepo.On("Create", mock.MatchedBy(func(po *models.PurchaseOrder) bool {
					return po.OrderNumber == "ORD123" &&
						po.OrderDate == "2023-10-05" &&
						po.TrackingCode == "TRK123" &&
						po.BuyerId == 1 &&
						po.CarrierId == 2 &&
						po.OrderStatusId == 3 &&
						po.WarehouseId == 4
				})).Return((*models.PurchaseOrder)(nil), conflictErr)
			},
			validateResult: func(t *testing.T, result *models.PurchaseOrder, err error) {
				require.Error(t, err)
				assert.Nil(t, result)
				assert.Contains(t, err.Error(), "conflict")
				assert.Contains(t, err.Error(), "purchase_order")
				assert.Contains(t, err.Error(), "order_number")
				assert.Contains(t, err.Error(), "ORD123")
				assert.Contains(t, err.Error(), "already exists")
			},
		},
		{
			name: "create_repository_error",
			purchaseOrder: &models.PurchaseOrder{
				OrderNumber:   "ORD456",
				OrderDate:     "2023-11-05",
				TrackingCode:  "TRK456",
				BuyerId:       2,
				CarrierId:     3,
				OrderStatusId: 4,
				WarehouseId:   5,
				OrderDetails: []models.OrderDetails{
					{
						CleanlinessStatus: "bad",
						Quantity:          10,
						Temperature:       7.2,
						ProductRecordId:   202,
					},
				},
			},
			setupMock: func(mockRepo *purchase_order.MockPurchaseOrderRepository) {
				mockRepo.On("Create", mock.AnythingOfType("*models.PurchaseOrder")).
					Return((*models.PurchaseOrder)(nil), errors.ErrGeneral)
			},
			validateResult: func(t *testing.T, result *models.PurchaseOrder, err error) {
				require.Error(t, err)
				assert.Equal(t, errors.ErrGeneral, err)
				assert.Nil(t, result)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			mockRepo := new(purchase_order.MockPurchaseOrderRepository)
			service := NewPurchaseOrderService(mockRepo)
			tt.setupMock(mockRepo)

			// Act
			result, err := service.Create(tt.purchaseOrder)

			// Assert
			tt.validateResult(t, result, err)
			mockRepo.AssertExpectations(t)
		})
	}
}
