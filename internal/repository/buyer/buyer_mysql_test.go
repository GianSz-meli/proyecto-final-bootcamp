package buyer

import (
	"ProyectoFinal/pkg/models"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
)

func TestBuyerMySql_GetById(t *testing.T) {
	tests := []struct {
		name           string
		id             int
		setupMock      func(sqlmock.Sqlmock)
		validateResult func(*testing.T, *models.Buyer, error)
	}{
		{
			name: "get_buyer_by_id_success",
			id:   1,
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name"}).
					AddRow(1, "CARD001", "Pepito", "Perez")

				mock.ExpectQuery(regexp.QuoteMeta(GetBuyer)).
					WithArgs(1).
					WillReturnRows(rows)
			},
			validateResult: func(t *testing.T, buyer *models.Buyer, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, buyer)
				assert.Equal(t, 1, buyer.Id)
				assert.Equal(t, "CARD001", buyer.CardNumberId)
				assert.Equal(t, "Pepito", buyer.FirstName)
				assert.Equal(t, "Perez", buyer.LastName)
			},
		},
		{
			name: "get_buyer_by_id_not_found",
			id:   999,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(GetBuyer)).
					WithArgs(999).
					WillReturnError(sql.ErrNoRows)
			},
			validateResult: func(t *testing.T, buyer *models.Buyer, err error) {
				assert.Error(t, err)
				assert.Nil(t, buyer)
				assert.Contains(t, err.Error(), "not found")
				assert.Contains(t, err.Error(), "buyer")
				assert.Contains(t, err.Error(), "id")
				assert.Contains(t, err.Error(), "999")
			},
		},
		{
			name: "get_buyer_by_id_database_connection_error",
			id:   1,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(GetBuyer)).
					WithArgs(1).
					WillReturnError(sql.ErrConnDone)
			},
			validateResult: func(t *testing.T, buyer *models.Buyer, err error) {
				assert.Error(t, err)
				assert.Nil(t, buyer)
				assert.Equal(t, sql.ErrConnDone, err)
			},
		},
		{
			name: "get_buyer_by_id_scan_error",
			id:   1,
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name"}).
					AddRow("invalid_id", "CARD001", "Pepito", "Perez")

				mock.ExpectQuery(regexp.QuoteMeta(GetBuyer)).
					WithArgs(1).
					WillReturnRows(rows)
			},
			validateResult: func(t *testing.T, buyer *models.Buyer, err error) {
				assert.Error(t, err)
				assert.Nil(t, buyer)
				assert.Contains(t, err.Error(), "converting")
			},
		},
		{
			name: "get_buyer_by_id_query_error",
			id:   1,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(GetBuyer)).
					WithArgs(1).
					WillReturnError(errors.New("syntax error in query"))
			},
			validateResult: func(t *testing.T, buyer *models.Buyer, err error) {
				assert.Error(t, err)
				assert.Nil(t, buyer)
				assert.Contains(t, err.Error(), "syntax error")
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

			repo := NewBuyerMySqlRepository(db)

			// Act
			result, err := repo.GetById(tt.id)

			// Assert
			tt.validateResult(t, result, err)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestBuyerMySql_GetAll(t *testing.T) {
	tests := []struct {
		name           string
		setupMock      func(sqlmock.Sqlmock)
		validateResult func(*testing.T, []*models.Buyer, error)
	}{
		{
			name: "get_all_buyers_success_multiple",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name"}).
					AddRow(1, "CARD001", "Pepito", "Perez").
					AddRow(2, "CARD002", "Juan", "Lopez").
					AddRow(3, "CARD003", "Ana", "Garcia")

				mock.ExpectQuery(regexp.QuoteMeta(GetAllBuyers)).
					WillReturnRows(rows)
			},
			validateResult: func(t *testing.T, buyers []*models.Buyer, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, buyers)
				assert.Len(t, buyers, 3)
				assert.Equal(t, 1, buyers[0].Id)
				assert.Equal(t, "CARD001", buyers[0].CardNumberId)
				assert.Equal(t, "Pepito", buyers[0].FirstName)
				assert.Equal(t, "Perez", buyers[0].LastName)
				assert.Equal(t, 2, buyers[1].Id)
				assert.Equal(t, "CARD002", buyers[1].CardNumberId)
				assert.Equal(t, "Juan", buyers[1].FirstName)
				assert.Equal(t, "Lopez", buyers[1].LastName)
				assert.Equal(t, 3, buyers[2].Id)
				assert.Equal(t, "CARD003", buyers[2].CardNumberId)
				assert.Equal(t, "Ana", buyers[2].FirstName)
				assert.Equal(t, "Garcia", buyers[2].LastName)
			},
		},
		{
			name: "get_all_buyers_success_single",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name"}).
					AddRow(1, "CARD001", "Pepito", "Perez")

				mock.ExpectQuery(regexp.QuoteMeta(GetAllBuyers)).
					WillReturnRows(rows)
			},
			validateResult: func(t *testing.T, buyers []*models.Buyer, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, buyers)
				assert.Len(t, buyers, 1)

				assert.Equal(t, 1, buyers[0].Id)
				assert.Equal(t, "CARD001", buyers[0].CardNumberId)
				assert.Equal(t, "Pepito", buyers[0].FirstName)
				assert.Equal(t, "Perez", buyers[0].LastName)
			},
		},
		{
			name: "get_all_buyers_empty_result",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name"})
				mock.ExpectQuery(regexp.QuoteMeta(GetAllBuyers)).
					WillReturnRows(rows)
			},
			validateResult: func(t *testing.T, buyers []*models.Buyer, err error) {
				assert.NoError(t, err)
				assert.Len(t, buyers, 0)
			},
		},
		{
			name: "get_all_buyers_query_error",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(GetAllBuyers)).
					WillReturnError(sql.ErrConnDone)
			},
			validateResult: func(t *testing.T, buyers []*models.Buyer, err error) {
				assert.Error(t, err)
				assert.Nil(t, buyers)
				assert.Equal(t, sql.ErrConnDone, err)
			},
		},
		{
			name: "get_all_buyers_scan_error",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name"}).
					AddRow(1, "CARD001", "Pepito", "Perez").
					AddRow("invalid_id", "CARD002", "Juan", "Lopez")

				mock.ExpectQuery(regexp.QuoteMeta(GetAllBuyers)).
					WillReturnRows(rows)
			},
			validateResult: func(t *testing.T, buyers []*models.Buyer, err error) {
				assert.Error(t, err)
				assert.Nil(t, buyers)
				assert.Contains(t, err.Error(), "converting")
			},
		},
		{
			name: "get_all_buyers_rows_error",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name"}).
					AddRow(1, "CARD001", "Pepito", "Perez").
					AddRow(2, "CARD002", "Juan", "Lopez").
					RowError(1, errors.New("row processing error"))

				mock.ExpectQuery(regexp.QuoteMeta(GetAllBuyers)).
					WillReturnRows(rows)
			},
			validateResult: func(t *testing.T, buyers []*models.Buyer, err error) {
				assert.Error(t, err)
				assert.Nil(t, buyers)
				assert.Contains(t, err.Error(), "row processing error")
			},
		},
		{
			name: "get_all_buyers_database_connection_error",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(GetAllBuyers)).
					WillReturnError(errors.New("database connection lost"))
			},
			validateResult: func(t *testing.T, buyers []*models.Buyer, err error) {
				assert.Error(t, err)
				assert.Nil(t, buyers)
				assert.Contains(t, err.Error(), "database connection lost")
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

			repo := NewBuyerMySqlRepository(db)

			// Act
			result, err := repo.GetAll()

			// Assert
			tt.validateResult(t, result, err)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestBuyerMySql_Create(t *testing.T) {
	tests := []struct {
		name           string
		inputBuyer     *models.Buyer
		setupMock      func(mock sqlmock.Sqlmock)
		validateResult func(t *testing.T, result *models.Buyer, err error)
	}{
		{
			name: "create_buyer_success",
			inputBuyer: &models.Buyer{
				CardNumberId: "CARD001",
				FirstName:    "Pepito",
				LastName:     "Perez",
			},
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(CreateBuyer)).
					WithArgs("CARD001", "Pepito", "Perez").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			validateResult: func(t *testing.T, result *models.Buyer, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, 1, result.Id)
				assert.Equal(t, "CARD001", result.CardNumberId)
				assert.Equal(t, "Pepito", result.FirstName)
				assert.Equal(t, "Perez", result.LastName)
			},
		},
		{
			name: "create_buyer_exec_error",
			inputBuyer: &models.Buyer{
				CardNumberId: "CARD002",
				FirstName:    "Juan",
				LastName:     "Lopez",
			},
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(CreateBuyer)).
					WithArgs("CARD002", "Juan", "Lopez").
					WillReturnError(errors.New("syntax error in query"))
			},
			validateResult: func(t *testing.T, result *models.Buyer, err error) {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.Contains(t, err.Error(), "syntax error")
			},
		},
		{
			name: "create_buyer_last_insert_id_error",
			inputBuyer: &models.Buyer{
				CardNumberId: "CARD003",
				FirstName:    "Ana",
				LastName:     "Garcia",
			},
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(CreateBuyer)).
					WithArgs("CARD003", "Ana", "Garcia").
					WillReturnResult(sqlmock.NewErrorResult(errors.New("last insert id error")))
			},
			validateResult: func(t *testing.T, result *models.Buyer, err error) {
				assert.Error(t, err)
				assert.Nil(t, result)
				assert.Contains(t, err.Error(), "last insert id error")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			tt.setupMock(mock)

			repo := NewBuyerMySqlRepository(db)

			// Act
			result, err := repo.Create(tt.inputBuyer)

			// Assert
			tt.validateResult(t, result, err)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestBuyerMySql_Update(t *testing.T) {
	tests := []struct {
		name           string
		inputBuyer     *models.Buyer
		setupMock      func(mock sqlmock.Sqlmock)
		validateResult func(t *testing.T, buyer *models.Buyer, err error)
	}{
		{
			name: "update_buyer_success",
			inputBuyer: &models.Buyer{
				Id:           1,
				CardNumberId: "CARD001",
				FirstName:    "Pepito",
				LastName:     "Perez",
			},
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(UpdateBuyer)).
					WithArgs("CARD001", "Pepito", "Perez", 1).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			validateResult: func(t *testing.T, buyer *models.Buyer, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, buyer)
				assert.Equal(t, 1, buyer.Id)
				assert.Equal(t, "CARD001", buyer.CardNumberId)
				assert.Equal(t, "Pepito", buyer.FirstName)
				assert.Equal(t, "Perez", buyer.LastName)
			},
		},
		{
			name: "update_buyer_exec_error",
			inputBuyer: &models.Buyer{
				Id:           2,
				CardNumberId: "CARD002",
				FirstName:    "Juan",
				LastName:     "Lopez",
			},
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(UpdateBuyer)).
					WithArgs("CARD002", "Juan", "Lopez", 2).
					WillReturnError(errors.New("update query error"))
			},
			validateResult: func(t *testing.T, buyer *models.Buyer, err error) {
				assert.Error(t, err)
				assert.Nil(t, buyer)
				assert.Contains(t, err.Error(), "update query error")
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

			repo := NewBuyerMySqlRepository(db)

			// Act
			result, err := repo.Update(tt.inputBuyer)

			// Assert
			tt.validateResult(t, result, err)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestBuyerMySql_Delete(t *testing.T) {
	tests := []struct {
		name           string
		id             int
		setupMock      func(mock sqlmock.Sqlmock)
		validateResult func(t *testing.T, err error)
	}{
		{
			name: "delete_buyer_success",
			id:   1,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(DeleteBuyer)).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			validateResult: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "delete_buyer_exec_error",
			id:   2,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(DeleteBuyer)).
					WithArgs(2).
					WillReturnError(errors.New("delete exec error"))
			},
			validateResult: func(t *testing.T, err error) {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "delete exec error")
			},
		},
		{
			name: "delete_buyer_no_rows_affected",
			id:   3,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(DeleteBuyer)).
					WithArgs(3).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			validateResult: func(t *testing.T, err error) {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "buyer")
				assert.Contains(t, err.Error(), "id")
				assert.Contains(t, err.Error(), "3")
			},
		},
		{
			name: "delete_buyer_rows_affected_error",
			id:   4,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(regexp.QuoteMeta(DeleteBuyer)).
					WithArgs(4).
					WillReturnResult(sqlmock.NewErrorResult(errors.New("rows affected error")))
			},
			validateResult: func(t *testing.T, err error) {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "rows affected error")
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
			repo := NewBuyerMySqlRepository(db)

			// Act
			err = repo.Delete(tt.id)

			// Assert
			tt.validateResult(t, err)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestBuyerMySql_GetByIdWithOrderCount(t *testing.T) {
	tests := []struct {
		name           string
		id             int
		setupMock      func(mock sqlmock.Sqlmock)
		validateResult func(t *testing.T, buyer *models.BuyerWithOrderCount, err error)
	}{
		{
			name: "get_buyer_with_order_count_success",
			id:   1,
			setupMock: func(mock sqlmock.Sqlmock) {
				columns := []string{"id", "card_number_id", "first_name", "last_name", "purchase_orders_count"}
				row := sqlmock.NewRows(columns).
					AddRow(1, "CARD001", "Pepito", "Perez", 5)
				mock.ExpectQuery(regexp.QuoteMeta(GetBuyerWithPurchaseOrders)).
					WithArgs(1).
					WillReturnRows(row)
			},
			validateResult: func(t *testing.T, buyer *models.BuyerWithOrderCount, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, buyer)
				assert.Equal(t, 1, buyer.Id)
				assert.Equal(t, "CARD001", buyer.CardNumberId)
				assert.Equal(t, "Pepito", buyer.FirstName)
				assert.Equal(t, "Perez", buyer.LastName)
				assert.Equal(t, 5, buyer.PurchaseOrdersCount)
			},
		},
		{
			name: "get_buyer_with_order_count_not_found",
			id:   999,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(GetBuyerWithPurchaseOrders)).
					WithArgs(999).
					WillReturnError(sql.ErrNoRows)
			},
			validateResult: func(t *testing.T, buyer *models.BuyerWithOrderCount, err error) {
				assert.Error(t, err)
				assert.Nil(t, buyer)
				assert.Contains(t, err.Error(), "buyer")
				assert.Contains(t, err.Error(), "id")
				assert.Contains(t, err.Error(), "999")
			},
		},
		{
			name: "get_buyer_with_order_count_query_error",
			id:   1,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(GetBuyerWithPurchaseOrders)).
					WithArgs(1).
					WillReturnError(errors.New("database error"))
			},
			validateResult: func(t *testing.T, buyer *models.BuyerWithOrderCount, err error) {
				assert.Error(t, err)
				assert.Nil(t, buyer)
				assert.Contains(t, err.Error(), "database error")
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

			repo := NewBuyerMySqlRepository(db)

			// Act: se invoca la función
			result, err := repo.GetByIdWithOrderCount(tt.id)

			// Assert: se valida el resultado
			tt.validateResult(t, result, err)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestBuyerMySql_GetAllWithOrderCount(t *testing.T) {
	tests := []struct {
		name           string
		setupMock      func(mock sqlmock.Sqlmock)
		validateResult func(t *testing.T, buyers []*models.BuyerWithOrderCount, err error)
	}{
		{
			name: "get_all_with_order_count_success_multiple",
			setupMock: func(mock sqlmock.Sqlmock) {
				columns := []string{"id", "card_number_id", "first_name", "last_name", "purchase_orders_count"}
				rows := sqlmock.NewRows(columns).
					AddRow(1, "CARD001", "Pepito", "Perez", 3).
					AddRow(2, "CARD002", "Juan", "Lopez", 0).
					AddRow(3, "CARD003", "Ana", "Garcia", 7)
				mock.ExpectQuery(regexp.QuoteMeta(GetAllBuyersWithPurchaseOrders)).
					WillReturnRows(rows)
			},
			validateResult: func(t *testing.T, buyers []*models.BuyerWithOrderCount, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, buyers)
				assert.Len(t, buyers, 3)
				assert.Equal(t, 1, buyers[0].Id)
				assert.Equal(t, "CARD001", buyers[0].CardNumberId)
				assert.Equal(t, "Pepito", buyers[0].FirstName)
				assert.Equal(t, "Perez", buyers[0].LastName)
				assert.Equal(t, 3, buyers[0].PurchaseOrdersCount)

				assert.Equal(t, 2, buyers[1].Id)
				assert.Equal(t, "CARD002", buyers[1].CardNumberId)
				assert.Equal(t, "Juan", buyers[1].FirstName)
				assert.Equal(t, "Lopez", buyers[1].LastName)
				assert.Equal(t, 0, buyers[1].PurchaseOrdersCount)

				assert.Equal(t, 3, buyers[2].Id)
				assert.Equal(t, "CARD003", buyers[2].CardNumberId)
				assert.Equal(t, "Ana", buyers[2].FirstName)
				assert.Equal(t, "Garcia", buyers[2].LastName)
				assert.Equal(t, 7, buyers[2].PurchaseOrdersCount)
			},
		},
		{
			name: "get_all_with_order_count_success_single",
			setupMock: func(mock sqlmock.Sqlmock) {
				columns := []string{"id", "card_number_id", "first_name", "last_name", "purchase_orders_count"}
				rows := sqlmock.NewRows(columns).AddRow(1, "CARD001", "Pepito", "Perez", 5)
				mock.ExpectQuery(regexp.QuoteMeta(GetAllBuyersWithPurchaseOrders)).
					WillReturnRows(rows)
			},
			validateResult: func(t *testing.T, buyers []*models.BuyerWithOrderCount, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, buyers)
				assert.Len(t, buyers, 1)

				assert.Equal(t, 1, buyers[0].Id)
				assert.Equal(t, "CARD001", buyers[0].CardNumberId)
				assert.Equal(t, "Pepito", buyers[0].FirstName)
				assert.Equal(t, "Perez", buyers[0].LastName)
				assert.Equal(t, 5, buyers[0].PurchaseOrdersCount)
			},
		},
		{
			name: "get_all_with_order_count_empty_result",
			setupMock: func(mock sqlmock.Sqlmock) {
				columns := []string{"id", "card_number_id", "first_name", "last_name", "purchase_orders_count"}
				rows := sqlmock.NewRows(columns)
				mock.ExpectQuery(regexp.QuoteMeta(GetAllBuyersWithPurchaseOrders)).
					WillReturnRows(rows)
			},
			validateResult: func(t *testing.T, buyers []*models.BuyerWithOrderCount, err error) {
				assert.NoError(t, err)
				assert.NotNil(t, buyers)
				assert.Len(t, buyers, 0)
			},
		},
		{
			name: "get_all_with_order_count_query_error",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(GetAllBuyersWithPurchaseOrders)).
					WillReturnError(sql.ErrConnDone)
			},
			validateResult: func(t *testing.T, buyers []*models.BuyerWithOrderCount, err error) {
				assert.Error(t, err)
				assert.Nil(t, buyers)
				assert.Equal(t, sql.ErrConnDone, err)
			},
		},
		{
			name: "get_all_with_order_count_scan_error",
			setupMock: func(mock sqlmock.Sqlmock) {
				columns := []string{"id", "card_number_id", "first_name", "last_name", "purchase_orders_count"}
				rows := sqlmock.NewRows(columns).
					AddRow(1, "CARD001", "Pepito", "Perez", "invalid_count")
				mock.ExpectQuery(regexp.QuoteMeta(GetAllBuyersWithPurchaseOrders)).
					WillReturnRows(rows)
			},
			validateResult: func(t *testing.T, buyers []*models.BuyerWithOrderCount, err error) {
				assert.Error(t, err)
				assert.Nil(t, buyers)
				assert.Contains(t, err.Error(), "converting")
			},
		},
		{
			name: "get_all_with_order_count_rows_error",
			setupMock: func(mock sqlmock.Sqlmock) {
				columns := []string{"id", "card_number_id", "first_name", "last_name", "purchase_orders_count"}
				rows := sqlmock.NewRows(columns).
					AddRow(1, "CARD001", "Pepito", "Perez", 3).
					AddRow(2, "CARD002", "Juan", "Lopez", 0).
					RowError(1, errors.New("row processing error"))
				mock.ExpectQuery(regexp.QuoteMeta(GetAllBuyersWithPurchaseOrders)).
					WillReturnRows(rows)
			},
			validateResult: func(t *testing.T, buyers []*models.BuyerWithOrderCount, err error) {
				assert.Error(t, err)
				assert.Nil(t, buyers)
				assert.Contains(t, err.Error(), "row processing error")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			tt.setupMock(mock)

			repo := NewBuyerMySqlRepository(db)
			result, err := repo.GetAllWithOrderCount()
			tt.validateResult(t, result, err)

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
