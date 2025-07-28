package employee

import (
	"ProyectoFinal/pkg/models"
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestMySQLRepository_GetAll(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewMySQLRepository(db)

	warehouseID1 := 1
	warehouseID2 := 2
	rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id"}).
		AddRow(1, "12345", "John", "Doe", &warehouseID1).
		AddRow(2, "67890", "Jane", "Smith", &warehouseID2)

	mock.ExpectQuery("SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees").
		WillReturnRows(rows)

	// Act
	employees, err := repo.GetAll()

	// Assert
	assert.NoError(t, err)
	assert.Len(t, employees, 2)
	assert.Equal(t, 1, employees[0].ID)
	assert.Equal(t, "12345", employees[0].CardNumberID)
	assert.Equal(t, "John", employees[0].FirstName)
	assert.Equal(t, "Doe", employees[0].LastName)
	assert.Equal(t, &warehouseID1, employees[0].WarehouseID)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestMySQLRepository_GetById_Existent(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewMySQLRepository(db)
	employeeID := 1
	warehouseID := 1

	rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id"}).
		AddRow(1, "12345", "John", "Doe", &warehouseID)

	mock.ExpectQuery("SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees WHERE id = \\?").
		WithArgs(employeeID).
		WillReturnRows(rows)

	// Act
	employee, err := repo.GetById(employeeID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 1, employee.ID)
	assert.Equal(t, "12345", employee.CardNumberID)
	assert.Equal(t, "John", employee.FirstName)
	assert.Equal(t, "Doe", employee.LastName)
	assert.Equal(t, &warehouseID, employee.WarehouseID)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestMySQLRepository_GetById_NonExistent(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewMySQLRepository(db)
	employeeID := 124

	mock.ExpectQuery("SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees WHERE id = \\?").
		WithArgs(employeeID).
		WillReturnError(sql.ErrNoRows)

	// Act
	employee, err := repo.GetById(employeeID)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, models.Employee{}, employee)
	assert.Contains(t, err.Error(), "not found")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestMySQLRepository_Create_OK(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewMySQLRepository(db)
	warehouseID := 1
	employee := &models.Employee{
		CardNumberID: "12345",
		FirstName:    "John",
		LastName:     "Doe",
		WarehouseID:  &warehouseID,
	}

	mock.ExpectExec("INSERT INTO employees \\(card_number_id, first_name, last_name, warehouse_id\\) VALUES \\(\\?, \\?, \\?, \\?\\)").
		WithArgs("12345", "John", "Doe", &warehouseID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Act
	err = repo.Create(employee)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 1, employee.ID)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestMySQLRepository_Create_Conflict(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewMySQLRepository(db)
	warehouseID := 1
	employee := &models.Employee{
		CardNumberID: "12345",
		FirstName:    "John",
		LastName:     "Doe",
		WarehouseID:  &warehouseID,
	}

	mock.ExpectExec("INSERT INTO employees \\(card_number_id, first_name, last_name, warehouse_id\\) VALUES \\(\\?, \\?, \\?, \\?\\)").
		WithArgs("12345", "John", "Doe", &warehouseID).
		WillReturnError(fmt.Errorf("duplicate entry"))

	// Act
	err = repo.Create(employee)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "duplicate entry")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestMySQLRepository_Update_OK(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewMySQLRepository(db)
	employeeID := 1
	warehouseID := 1
	employee := models.Employee{
		CardNumberID: "12345",
		FirstName:    "Johnny",
		LastName:     "Doe",
		WarehouseID:  &warehouseID,
	}

	// Mock GetById call first
	rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id"}).
		AddRow(1, "12345", "John", "Doe", &warehouseID)

	mock.ExpectQuery("SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees WHERE id = \\?").
		WithArgs(employeeID).
		WillReturnRows(rows)

	mock.ExpectExec("UPDATE employees SET card_number_id = \\?, first_name = \\?, last_name = \\?, warehouse_id = \\? WHERE id = \\?").
		WithArgs("12345", "Johnny", "Doe", &warehouseID, employeeID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// Act
	err = repo.Update(employeeID, employee)

	// Assert
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestMySQLRepository_Update_NonExistent(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewMySQLRepository(db)
	employeeID := 124
	warehouseID := 1
	employee := models.Employee{
		CardNumberID: "12345",
		FirstName:    "Johnny",
		LastName:     "Doe",
		WarehouseID:  &warehouseID,
	}

	mock.ExpectQuery("SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees WHERE id = \\?").
		WithArgs(employeeID).
		WillReturnError(sql.ErrNoRows)

	// Act
	err = repo.Update(employeeID, employee)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestMySQLRepository_Delete_OK(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewMySQLRepository(db)
	employeeID := 1

	mock.ExpectExec("DELETE FROM employees WHERE id = \\?").
		WithArgs(employeeID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	// Act
	err = repo.Delete(employeeID)

	// Assert
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestMySQLRepository_Delete_NonExistent(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewMySQLRepository(db)
	employeeID := 124

	mock.ExpectExec("DELETE FROM employees WHERE id = \\?").
		WithArgs(employeeID).
		WillReturnResult(sqlmock.NewResult(0, 0))

	// Act
	err = repo.Delete(employeeID)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")

	assert.NoError(t, mock.ExpectationsWereMet())
}
