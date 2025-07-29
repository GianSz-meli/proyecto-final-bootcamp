package inbound_order

import (
	"ProyectoFinal/pkg/models"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupInboundOrderMySQLTest(t *testing.T) (*mysqlRepository, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	repository := NewMySQLRepository(db).(*mysqlRepository)

	cleanup := func() {
		db.Close()
	}

	return repository, mock, cleanup
}

func TestMySQLRepository_Create_Success(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupInboundOrderMySQLTest(t)
	defer cleanup()

	orderDate := time.Date(2023, 10, 15, 10, 30, 0, 0, time.UTC)
	inboundOrder := &models.InboundOrder{
		OrderDate:      orderDate,
		OrderNumber:    "ORD001",
		EmployeeID:     1,
		ProductBatchID: 1,
		WarehouseID:    1,
	}

	mock.ExpectExec("INSERT INTO inbound_orders \\(order_date, order_number, employee_id, product_batch_id, warehouse_id\\) VALUES \\(\\?, \\?, \\?, \\?, \\?\\)").
		WithArgs(orderDate, "ORD001", 1, 1, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Act
	err := repository.Create(inboundOrder)

	// Assert
	require.NoError(t, err)
	assert.Equal(t, 1, inboundOrder.ID)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestMySQLRepository_Create_Error(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupInboundOrderMySQLTest(t)
	defer cleanup()

	orderDate := time.Date(2023, 10, 15, 10, 30, 0, 0, time.UTC)
	inboundOrder := &models.InboundOrder{
		OrderDate:      orderDate,
		OrderNumber:    "ORD001",
		EmployeeID:     1,
		ProductBatchID: 1,
		WarehouseID:    1,
	}

	mock.ExpectExec("INSERT INTO inbound_orders \\(order_date, order_number, employee_id, product_batch_id, warehouse_id\\) VALUES \\(\\?, \\?, \\?, \\?, \\?\\)").
		WithArgs(orderDate, "ORD001", 1, 1, 1).
		WillReturnError(sql.ErrConnDone)

	// Act
	err := repository.Create(inboundOrder)

	// Assert
	require.Error(t, err)
	assert.Equal(t, sql.ErrConnDone, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestMySQLRepository_Create_LastInsertIdError(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupInboundOrderMySQLTest(t)
	defer cleanup()

	orderDate := time.Date(2023, 10, 15, 10, 30, 0, 0, time.UTC)
	inboundOrder := &models.InboundOrder{
		OrderDate:      orderDate,
		OrderNumber:    "ORD001",
		EmployeeID:     1,
		ProductBatchID: 1,
		WarehouseID:    1,
	}

	mock.ExpectExec("INSERT INTO inbound_orders \\(order_date, order_number, employee_id, product_batch_id, warehouse_id\\) VALUES \\(\\?, \\?, \\?, \\?, \\?\\)").
		WithArgs(orderDate, "ORD001", 1, 1, 1).
		WillReturnResult(sqlmock.NewErrorResult(sql.ErrNoRows))

	// Act
	err := repository.Create(inboundOrder)

	// Assert
	require.Error(t, err)
	assert.Equal(t, sql.ErrNoRows, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestMySQLRepository_GetEmployeeInboundOrdersReportByEmployeeId_Success(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupInboundOrderMySQLTest(t)
	defer cleanup()

	employeeId := 1
	expectedReport := models.EmployeeInboundOrdersReport{
		ID:                 1,
		CardNumberID:       "EMP001",
		FirstName:          "John",
		LastName:           "Doe",
		WarehouseID:        1,
		InboundOrdersCount: 5,
	}

	rows := sqlmock.NewRows([]string{
		"id", "card_number_id", "first_name", "last_name", "warehouse_id", "inbound_orders_count",
	}).AddRow(1, "EMP001", "John", "Doe", 1, 5)

	mock.ExpectQuery("SELECT e\\.id, e\\.card_number_id, e\\.first_name, e\\.last_name, e\\.warehouse_id, COUNT\\(io\\.id\\) as inbound_orders_count FROM employees e LEFT JOIN inbound_orders io ON e\\.id = io\\.employee_id WHERE e\\.id = \\? GROUP BY e\\.id, e\\.card_number_id, e\\.first_name, e\\.last_name, e\\.warehouse_id").
		WithArgs(employeeId).
		WillReturnRows(rows)

	// Act
	result, err := repository.GetEmployeeInboundOrdersReportByEmployeeId(employeeId)

	// Assert
	require.NoError(t, err)
	assert.Equal(t, expectedReport, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestMySQLRepository_GetEmployeeInboundOrdersReportByEmployeeId_NotFound(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupInboundOrderMySQLTest(t)
	defer cleanup()

	employeeId := 999

	mock.ExpectQuery("SELECT e\\.id, e\\.card_number_id, e\\.first_name, e\\.last_name, e\\.warehouse_id, COUNT\\(io\\.id\\) as inbound_orders_count FROM employees e LEFT JOIN inbound_orders io ON e\\.id = io\\.employee_id WHERE e\\.id = \\? GROUP BY e\\.id, e\\.card_number_id, e\\.first_name, e\\.last_name, e\\.warehouse_id").
		WithArgs(employeeId).
		WillReturnError(sql.ErrNoRows)

	// Act
	result, err := repository.GetEmployeeInboundOrdersReportByEmployeeId(employeeId)

	// Assert
	require.Error(t, err)
	assert.Contains(t, err.Error(), "employee")
	assert.Contains(t, err.Error(), "not found")
	assert.Equal(t, models.EmployeeInboundOrdersReport{}, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestMySQLRepository_GetEmployeeInboundOrdersReportByEmployeeId_DatabaseError(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupInboundOrderMySQLTest(t)
	defer cleanup()

	employeeId := 1

	mock.ExpectQuery("SELECT e\\.id, e\\.card_number_id, e\\.first_name, e\\.last_name, e\\.warehouse_id, COUNT\\(io\\.id\\) as inbound_orders_count FROM employees e LEFT JOIN inbound_orders io ON e\\.id = io\\.employee_id WHERE e\\.id = \\? GROUP BY e\\.id, e\\.card_number_id, e\\.first_name, e\\.last_name, e\\.warehouse_id").
		WithArgs(employeeId).
		WillReturnError(sql.ErrConnDone)

	// Act
	result, err := repository.GetEmployeeInboundOrdersReportByEmployeeId(employeeId)

	// Assert
	require.Error(t, err)
	assert.Equal(t, sql.ErrConnDone, err)
	assert.Equal(t, models.EmployeeInboundOrdersReport{}, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestMySQLRepository_GetEmployeeInboundOrdersReportAll_Success(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupInboundOrderMySQLTest(t)
	defer cleanup()

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

	rows := sqlmock.NewRows([]string{
		"id", "card_number_id", "first_name", "last_name", "warehouse_id", "inbound_orders_count",
	}).
		AddRow(1, "EMP001", "John", "Doe", 1, 5).
		AddRow(2, "EMP002", "Jane", "Smith", 2, 3)

	mock.ExpectQuery("SELECT e\\.id, e\\.card_number_id, e\\.first_name, e\\.last_name, e\\.warehouse_id, COUNT\\(io\\.id\\) as inbound_orders_count FROM employees e LEFT JOIN inbound_orders io ON e\\.id = io\\.employee_id GROUP BY e\\.id, e\\.card_number_id, e\\.first_name, e\\.last_name, e\\.warehouse_id ORDER BY e\\.id").
		WillReturnRows(rows)

	// Act
	result, err := repository.GetEmployeeInboundOrdersReportAll()

	// Assert
	require.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, expectedReports, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestMySQLRepository_GetEmployeeInboundOrdersReportAll_QueryError(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupInboundOrderMySQLTest(t)
	defer cleanup()

	mock.ExpectQuery("SELECT e\\.id, e\\.card_number_id, e\\.first_name, e\\.last_name, e\\.warehouse_id, COUNT\\(io\\.id\\) as inbound_orders_count FROM employees e LEFT JOIN inbound_orders io ON e\\.id = io\\.employee_id GROUP BY e\\.id, e\\.card_number_id, e\\.first_name, e\\.last_name, e\\.warehouse_id ORDER BY e\\.id").
		WillReturnError(sql.ErrConnDone)

	// Act
	result, err := repository.GetEmployeeInboundOrdersReportAll()

	// Assert
	require.Error(t, err)
	assert.Equal(t, sql.ErrConnDone, err)
	assert.Nil(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestMySQLRepository_GetEmployeeInboundOrdersReportAll_ScanError(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupInboundOrderMySQLTest(t)
	defer cleanup()

	rows := sqlmock.NewRows([]string{
		"id", "card_number_id", "first_name", "last_name", "warehouse_id", "inbound_orders_count",
	}).AddRow("invalid_id", "EMP001", "John", "Doe", 1, 5) // invalid_id should cause scan error

	mock.ExpectQuery("SELECT e\\.id, e\\.card_number_id, e\\.first_name, e\\.last_name, e\\.warehouse_id, COUNT\\(io\\.id\\) as inbound_orders_count FROM employees e LEFT JOIN inbound_orders io ON e\\.id = io\\.employee_id GROUP BY e\\.id, e\\.card_number_id, e\\.first_name, e\\.last_name, e\\.warehouse_id ORDER BY e\\.id").
		WillReturnRows(rows)

	// Act
	result, err := repository.GetEmployeeInboundOrdersReportAll()

	// Assert
	require.Error(t, err)
	assert.Nil(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestMySQLRepository_GetEmployeeInboundOrdersReportAll_EmptyResult(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupInboundOrderMySQLTest(t)
	defer cleanup()

	rows := sqlmock.NewRows([]string{
		"id", "card_number_id", "first_name", "last_name", "warehouse_id", "inbound_orders_count",
	})

	mock.ExpectQuery("SELECT e\\.id, e\\.card_number_id, e\\.first_name, e\\.last_name, e\\.warehouse_id, COUNT\\(io\\.id\\) as inbound_orders_count FROM employees e LEFT JOIN inbound_orders io ON e\\.id = io\\.employee_id GROUP BY e\\.id, e\\.card_number_id, e\\.first_name, e\\.last_name, e\\.warehouse_id ORDER BY e\\.id").
		WillReturnRows(rows)

	// Act
	result, err := repository.GetEmployeeInboundOrdersReportAll()

	// Assert
	require.NoError(t, err)
	assert.Empty(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}
