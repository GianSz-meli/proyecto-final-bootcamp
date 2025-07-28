package purchase_order

import (
	"ProyectoFinal/pkg/models"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
)

func TestPurchaseOrderMySql_Create(t *testing.T) {
	tests := []struct {
		name           string
		purchaseOrder  *models.PurchaseOrder
		setupMock      func(mock sqlmock.Sqlmock)
		validateResult func(t *testing.T, result *models.PurchaseOrder, err error)
	}{
		{
			name: "create_success",
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
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(CreatePurchaseOrder)).
					WithArgs("ORD123", "2023-10-05", "TRK123", 1, 2, 3, 4).
					WillReturnResult(sqlmock.NewResult(1, 1))
				prep := mock.ExpectPrepare(regexp.QuoteMeta(CreateOrderDetail))
				prep.ExpectExec().
					WithArgs("good", 5, 4.5, 101, 1).
					WillReturnResult(sqlmock.NewResult(10, 1))
				mock.ExpectCommit()
			},
			validateResult: func(t *testing.T, result *models.PurchaseOrder, err error) {
				require.NoError(t, err)
				require.NotNil(t, result)
				assert.Equal(t, 1, result.Id)
				require.Len(t, result.OrderDetails, 1)
				od := result.OrderDetails[0]
				assert.Equal(t, 10, od.Id)
				assert.Equal(t, 1, od.PurchaseOrderId)
			},
		},
		{
			name: "create_begin_transaction_error",
			purchaseOrder: &models.PurchaseOrder{
				OrderNumber:   "ORD001",
				OrderDate:     "2023-10-05",
				TrackingCode:  "TRK001",
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
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin().WillReturnError(errors.New("begin transaction error"))
			},
			validateResult: func(t *testing.T, result *models.PurchaseOrder, err error) {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.EqualError(t, err, "begin transaction error")
			},
		},
		{
			name: "create_order_error",
			purchaseOrder: &models.PurchaseOrder{
				OrderNumber:   "ORD456",
				OrderDate:     "2023-10-05",
				TrackingCode:  "TRK456",
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
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(CreatePurchaseOrder)).
					WithArgs("ORD456", "2023-10-05", "TRK456", 1, 2, 3, 4).
					WillReturnError(errors.New("order error"))
			},
			validateResult: func(t *testing.T, result *models.PurchaseOrder, err error) {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.Contains(t, err.Error(), "order error")
			},
		},
		{
			name: "create_details_error",
			purchaseOrder: &models.PurchaseOrder{
				OrderNumber:   "ORD789",
				OrderDate:     "2023-12-01",
				TrackingCode:  "TRK789",
				BuyerId:       5,
				CarrierId:     6,
				OrderStatusId: 7,
				WarehouseId:   8,
				OrderDetails: []models.OrderDetails{
					{
						CleanlinessStatus: "excellent",
						Quantity:          3,
						Temperature:       2.5,
						ProductRecordId:   303,
					},
				},
			},
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(CreatePurchaseOrder)).
					WithArgs("ORD789", "2023-12-01", "TRK789", 5, 6, 7, 8).
					WillReturnResult(sqlmock.NewResult(2, 1))
				prep := mock.ExpectPrepare(regexp.QuoteMeta(CreateOrderDetail))
				prep.ExpectExec().
					WithArgs("excellent", 3, 2.5, 303, 2).
					WillReturnError(errors.New("detail error"))
			},
			validateResult: func(t *testing.T, result *models.PurchaseOrder, err error) {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.Contains(t, err.Error(), "detail error")
			},
		},
		{
			name: "create_commit_error",
			purchaseOrder: &models.PurchaseOrder{
				OrderNumber:   "ORDXYZ",
				OrderDate:     "2023-12-31",
				TrackingCode:  "TRKXYZ",
				BuyerId:       9,
				CarrierId:     10,
				OrderStatusId: 11,
				WarehouseId:   12,
				OrderDetails: []models.OrderDetails{
					{
						CleanlinessStatus: "excellent",
						Quantity:          7,
						Temperature:       3.2,
						ProductRecordId:   999,
					},
				},
			},
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(CreatePurchaseOrder)).
					WithArgs("ORDXYZ", "2023-12-31", "TRKXYZ", 9, 10, 11, 12).
					WillReturnResult(sqlmock.NewResult(3, 1))
				prep := mock.ExpectPrepare(regexp.QuoteMeta(CreateOrderDetail))
				prep.ExpectExec().
					WithArgs("excellent", 7, 3.2, 999, 3).
					WillReturnResult(sqlmock.NewResult(15, 1))
				mock.ExpectCommit().WillReturnError(errors.New("commit error"))
			},
			validateResult: func(t *testing.T, result *models.PurchaseOrder, err error) {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.EqualError(t, err, "commit error")
			},
		},
		{
			name: "create_order_last_insert_id_error",
			purchaseOrder: &models.PurchaseOrder{
				OrderNumber:   "ORD999",
				OrderDate:     "2023-11-11",
				TrackingCode:  "TRK999",
				BuyerId:       5,
				CarrierId:     6,
				OrderStatusId: 7,
				WarehouseId:   8,
				OrderDetails: []models.OrderDetails{
					{
						CleanlinessStatus: "good",
						Quantity:          10,
						Temperature:       3.3,
						ProductRecordId:   202,
					},
				},
			},
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(CreatePurchaseOrder)).
					WithArgs("ORD999", "2023-11-11", "TRK999", 5, 6, 7, 8).
					WillReturnResult(sqlmock.NewErrorResult(errors.New("last insert id error")))
			},
			validateResult: func(t *testing.T, result *models.PurchaseOrder, err error) {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.Contains(t, err.Error(), "last insert id error")
			},
		},
		{
			name: "create_no_order_details",
			purchaseOrder: &models.PurchaseOrder{
				OrderNumber:   "ORDEMPTY",
				OrderDate:     "2023-11-11",
				TrackingCode:  "TRKEMPTY",
				BuyerId:       5,
				CarrierId:     6,
				OrderStatusId: 7,
				WarehouseId:   8,
				OrderDetails:  []models.OrderDetails{},
			},
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(CreatePurchaseOrder)).
					WithArgs("ORDEMPTY", "2023-11-11", "TRKEMPTY", 5, 6, 7, 8).
					WillReturnResult(sqlmock.NewResult(4, 1))
				mock.ExpectCommit()
			},
			validateResult: func(t *testing.T, result *models.PurchaseOrder, err error) {
				assert.NoError(t, err)
				require.NotNil(t, result)
				assert.Equal(t, 4, result.Id)
				assert.Len(t, result.OrderDetails, 0)
			},
		},
		{
			name: "create_prepare_error",
			purchaseOrder: &models.PurchaseOrder{
				OrderNumber:   "ORDPREP",
				OrderDate:     "2023-11-11",
				TrackingCode:  "TRKPREP",
				BuyerId:       5,
				CarrierId:     6,
				OrderStatusId: 7,
				WarehouseId:   8,
				OrderDetails: []models.OrderDetails{
					{
						CleanlinessStatus: "good",
						Quantity:          10,
						Temperature:       3.3,
						ProductRecordId:   202,
					},
				},
			},
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(CreatePurchaseOrder)).
					WithArgs("ORDPREP", "2023-11-11", "TRKPREP", 5, 6, 7, 8).
					WillReturnResult(sqlmock.NewResult(5, 1))
				mock.ExpectPrepare(regexp.QuoteMeta(CreateOrderDetail)).
					WillReturnError(errors.New("prepare error"))
			},
			validateResult: func(t *testing.T, result *models.PurchaseOrder, err error) {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.Contains(t, err.Error(), "prepare error")
			},
		},
		{
			name: "create_details_last_insert_id_error",
			purchaseOrder: &models.PurchaseOrder{
				OrderNumber:   "ORD567",
				OrderDate:     "2023-12-15",
				TrackingCode:  "TRK567",
				BuyerId:       5,
				CarrierId:     6,
				OrderStatusId: 7,
				WarehouseId:   8,
				OrderDetails: []models.OrderDetails{
					{
						CleanlinessStatus: "excellent",
						Quantity:          2,
						Temperature:       3.5,
						ProductRecordId:   111,
					},
				},
			},
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(CreatePurchaseOrder)).
					WithArgs("ORD567", "2023-12-15", "TRK567", 5, 6, 7, 8).
					WillReturnResult(sqlmock.NewResult(4, 1))
				prep := mock.ExpectPrepare(regexp.QuoteMeta(CreateOrderDetail))
				prep.ExpectExec().
					WithArgs("excellent", 2, 3.5, 111, 4).
					WillReturnResult(sqlmock.NewErrorResult(errors.New("detail last insert id error")))
			},
			validateResult: func(t *testing.T, result *models.PurchaseOrder, err error) {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.Contains(t, err.Error(), "detail last insert id error")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			tt.setupMock(mock)
			repo := NewPurchaseOrderMySqlRepository(db)

			// Act
			result, err := repo.Create(tt.purchaseOrder)

			// Assert
			tt.validateResult(t, result, err)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
