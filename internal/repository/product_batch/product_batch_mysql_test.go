package repository

import (
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func setupProductBatchMySQLTest(t *testing.T) (*ProductBatchMySQL, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	repository := NewProductBatchMySQL(db)

	cleanup := func() {
		db.Close()
	}

	return repository, mock, cleanup
}

func TestProductBatchMySQL_Create_Success(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupProductBatchMySQLTest(t)
	defer cleanup()

	productBatch := models.ProductBatch{
		BatchNumber:        "BATCH001",
		CurrentQuantity:    100,
		CurrentTemperature: 15.5,
		DueDate:            time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
		InitialQuantity:    100,
		ManufacturingDate:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		ManufacturingHour:  10,
		MinimumTemperature: 10.0,
		ProductID:          1,
		SectionID:          1,
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO product_batches \\(batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?\\)").WithArgs(
		"BATCH001", 100, 15.5, time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC), 100, time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), "10:00:00", 10.0, 1, 1,
	).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Act
	result, err := repository.Create(productBatch)

	// Assert
	require.NoError(t, err)
	require.Equal(t, 1, result.ID)
	require.Equal(t, "BATCH001", result.BatchNumber)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestProductBatchMySQL_Create_BeginTransactionError(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupProductBatchMySQLTest(t)
	defer cleanup()

	productBatch := models.ProductBatch{
		BatchNumber:        "BATCH001",
		CurrentQuantity:    100,
		CurrentTemperature: 15.5,
		DueDate:            time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
		InitialQuantity:    100,
		ManufacturingDate:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		ManufacturingHour:  10,
		MinimumTemperature: 10.0,
		ProductID:          1,
		SectionID:          1,
	}

	mock.ExpectBegin().WillReturnError(sql.ErrConnDone)

	// Act
	result, err := repository.Create(productBatch)

	// Assert
	require.Error(t, err)
	require.Equal(t, productBatch, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestProductBatchMySQL_Create_ExecError(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupProductBatchMySQLTest(t)
	defer cleanup()

	productBatch := models.ProductBatch{
		BatchNumber:        "BATCH001",
		CurrentQuantity:    100,
		CurrentTemperature: 15.5,
		DueDate:            time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
		InitialQuantity:    100,
		ManufacturingDate:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		ManufacturingHour:  10,
		MinimumTemperature: 10.0,
		ProductID:          1,
		SectionID:          1,
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO product_batches \\(batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?\\)").WithArgs(
		"BATCH001", 100, 15.5, time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC), 100, time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), "10:00:00", 10.0, 1, 1,
	).WillReturnError(sql.ErrConnDone)
	mock.ExpectRollback()

	// Act
	result, err := repository.Create(productBatch)

	// Assert
	require.Error(t, err)
	require.Equal(t, productBatch, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestProductBatchMySQL_Create_LastInsertIdError(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupProductBatchMySQLTest(t)
	defer cleanup()

	productBatch := models.ProductBatch{
		BatchNumber:        "BATCH001",
		CurrentQuantity:    100,
		CurrentTemperature: 15.5,
		DueDate:            time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
		InitialQuantity:    100,
		ManufacturingDate:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		ManufacturingHour:  10,
		MinimumTemperature: 10.0,
		ProductID:          1,
		SectionID:          1,
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO product_batches \\(batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?\\)").WithArgs(
		"BATCH001", 100, 15.5, time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC), 100, time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), "10:00:00", 10.0, 1, 1,
	).WillReturnResult(sqlmock.NewErrorResult(sql.ErrConnDone))
	mock.ExpectRollback()

	// Act
	result, err := repository.Create(productBatch)

	// Assert
	require.Error(t, err)
	require.Equal(t, productBatch, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestProductBatchMySQL_Create_CommitError(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupProductBatchMySQLTest(t)
	defer cleanup()

	productBatch := models.ProductBatch{
		BatchNumber:        "BATCH001",
		CurrentQuantity:    100,
		CurrentTemperature: 15.5,
		DueDate:            time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
		InitialQuantity:    100,
		ManufacturingDate:  time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		ManufacturingHour:  10,
		MinimumTemperature: 10.0,
		ProductID:          1,
		SectionID:          1,
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO product_batches \\(batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?\\)").WithArgs(
		"BATCH001", 100, 15.5, time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC), 100, time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), "10:00:00", 10.0, 1, 1,
	).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit().WillReturnError(sql.ErrConnDone)

	// Act
	result, err := repository.Create(productBatch)

	// Assert
	require.Error(t, err)
	// En caso de error en commit, el ID ya se asignó antes del commit
	require.Equal(t, 1, result.ID)
	require.Equal(t, "BATCH001", result.BatchNumber)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestProductBatchMySQL_ExistsByBatchNumber_Exists(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupProductBatchMySQLTest(t)
	defer cleanup()

	rows := sqlmock.NewRows([]string{"EXISTS(SELECT 1 FROM product_batches WHERE batch_number = ?)"}).AddRow(1)

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM product_batches WHERE batch_number = \\?\\)").WithArgs("BATCH001").WillReturnRows(rows)

	// Act
	result := repository.ExistsByBatchNumber("BATCH001")

	// Assert
	require.True(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestProductBatchMySQL_ExistsByBatchNumber_NotExists(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupProductBatchMySQLTest(t)
	defer cleanup()

	rows := sqlmock.NewRows([]string{"EXISTS(SELECT 1 FROM product_batches WHERE batch_number = ?)"}).AddRow(0)

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM product_batches WHERE batch_number = \\?\\)").WithArgs("BATCH999").WillReturnRows(rows)

	// Act
	result := repository.ExistsByBatchNumber("BATCH999")

	// Assert
	require.False(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestProductBatchMySQL_ExistsByBatchNumber_Error(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupProductBatchMySQLTest(t)
	defer cleanup()

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM product_batches WHERE batch_number = \\?\\)").WithArgs("BATCH001").WillReturnError(sql.ErrConnDone)

	// Act
	result := repository.ExistsByBatchNumber("BATCH001")

	// Assert
	require.False(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestProductBatchMySQL_ProductExists_Exists(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupProductBatchMySQLTest(t)
	defer cleanup()

	rows := sqlmock.NewRows([]string{"EXISTS(SELECT 1 FROM products WHERE id = ?)"}).AddRow(1)

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM products WHERE id = \\?\\)").WithArgs(1).WillReturnRows(rows)

	// Act
	result := repository.ProductExists(1)

	// Assert
	require.True(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestProductBatchMySQL_ProductExists_NotExists(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupProductBatchMySQLTest(t)
	defer cleanup()

	rows := sqlmock.NewRows([]string{"EXISTS(SELECT 1 FROM products WHERE id = ?)"}).AddRow(0)

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM products WHERE id = \\?\\)").WithArgs(999).WillReturnRows(rows)

	// Act
	result := repository.ProductExists(999)

	// Assert
	require.False(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestProductBatchMySQL_ProductExists_Error(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupProductBatchMySQLTest(t)
	defer cleanup()

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM products WHERE id = \\?\\)").WithArgs(1).WillReturnError(sql.ErrConnDone)

	// Act
	result := repository.ProductExists(1)

	// Assert
	require.False(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestProductBatchMySQL_SectionExists_Exists(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupProductBatchMySQLTest(t)
	defer cleanup()

	rows := sqlmock.NewRows([]string{"EXISTS(SELECT 1 FROM sections WHERE id = ?)"}).AddRow(1)

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM sections WHERE id = \\?\\)").WithArgs(1).WillReturnRows(rows)

	// Act
	result := repository.SectionExists(1)

	// Assert
	require.True(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestProductBatchMySQL_SectionExists_NotExists(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupProductBatchMySQLTest(t)
	defer cleanup()

	rows := sqlmock.NewRows([]string{"EXISTS(SELECT 1 FROM sections WHERE id = ?)"}).AddRow(0)

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM sections WHERE id = \\?\\)").WithArgs(999).WillReturnRows(rows)

	// Act
	result := repository.SectionExists(999)

	// Assert
	require.False(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestProductBatchMySQL_SectionExists_Error(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupProductBatchMySQLTest(t)
	defer cleanup()

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM sections WHERE id = \\?\\)").WithArgs(1).WillReturnError(sql.ErrConnDone)

	// Act
	result := repository.SectionExists(1)

	// Assert
	require.False(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestProductBatchMySQL_GetProductCountBySection_AllSections_Success(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupProductBatchMySQLTest(t)
	defer cleanup()

	rows := sqlmock.NewRows([]string{"section_id", "section_number", "products_count"}).
		AddRow(1, "SEC001", 5).
		AddRow(2, "SEC002", 3).
		AddRow(3, "SEC003", 0)

	mock.ExpectQuery("SELECT s\\.id as section_id, s\\.section_number, COUNT\\(pb\\.id\\) as products_count FROM sections s LEFT JOIN product_batches pb ON s\\.id = pb\\.section_id GROUP BY s\\.id, s\\.section_number").WillReturnRows(rows)

	// Act
	result, err := repository.GetProductCountBySection(nil)

	// Assert
	require.NoError(t, err)
	require.Len(t, result, 3)
	require.Equal(t, 1, result[0].SectionID)
	require.Equal(t, "SEC001", result[0].SectionNumber)
	require.Equal(t, 5, result[0].ProductsCount)
	require.Equal(t, 2, result[1].SectionID)
	require.Equal(t, "SEC002", result[1].SectionNumber)
	require.Equal(t, 3, result[1].ProductsCount)
	require.Equal(t, 3, result[2].SectionID)
	require.Equal(t, "SEC003", result[2].SectionNumber)
	require.Equal(t, 0, result[2].ProductsCount)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestProductBatchMySQL_GetProductCountBySection_SpecificSection_Success(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupProductBatchMySQLTest(t)
	defer cleanup()

	sectionID := 1

	// Mock para verificar si la sección existe
	existsRows := sqlmock.NewRows([]string{"EXISTS(SELECT 1 FROM sections WHERE id = ?)"}).AddRow(1)
	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM sections WHERE id = \\?\\)").WithArgs(1).WillReturnRows(existsRows)

	// Mock para obtener el reporte de la sección específica
	reportRows := sqlmock.NewRows([]string{"section_id", "section_number", "products_count"}).
		AddRow(1, "SEC001", 5)

	mock.ExpectQuery("SELECT s\\.id as section_id, s\\.section_number, COUNT\\(pb\\.id\\) as products_count FROM sections s LEFT JOIN product_batches pb ON s\\.id = pb\\.section_id WHERE s\\.id = \\? GROUP BY s\\.id, s\\.section_number").WithArgs(1).WillReturnRows(reportRows)

	// Act
	result, err := repository.GetProductCountBySection(&sectionID)

	// Assert
	require.NoError(t, err)
	require.Len(t, result, 1)
	require.Equal(t, 1, result[0].SectionID)
	require.Equal(t, "SEC001", result[0].SectionNumber)
	require.Equal(t, 5, result[0].ProductsCount)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestProductBatchMySQL_GetProductCountBySection_SectionNotFound_Error(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupProductBatchMySQLTest(t)
	defer cleanup()

	sectionID := 999

	// Mock para verificar si la sección existe (retorna false)
	existsRows := sqlmock.NewRows([]string{"EXISTS(SELECT 1 FROM sections WHERE id = ?)"}).AddRow(0)
	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM sections WHERE id = \\?\\)").WithArgs(999).WillReturnRows(existsRows)

	// Act
	result, err := repository.GetProductCountBySection(&sectionID)

	// Assert
	require.Error(t, err)
	require.Nil(t, result)
	require.ErrorIs(t, err, errors.ErrNotFound)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestProductBatchMySQL_GetProductCountBySection_AllSections_QueryError(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupProductBatchMySQLTest(t)
	defer cleanup()

	mock.ExpectQuery("SELECT s\\.id as section_id, s\\.section_number, COUNT\\(pb\\.id\\) as products_count FROM sections s LEFT JOIN product_batches pb ON s\\.id = pb\\.section_id GROUP BY s\\.id, s\\.section_number").WillReturnError(sql.ErrConnDone)

	// Act
	result, err := repository.GetProductCountBySection(nil)

	// Assert
	require.Error(t, err)
	require.Nil(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestProductBatchMySQL_GetProductCountBySection_SpecificSection_QueryError(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupProductBatchMySQLTest(t)
	defer cleanup()

	sectionID := 1

	// Mock para verificar si la sección existe
	existsRows := sqlmock.NewRows([]string{"EXISTS(SELECT 1 FROM sections WHERE id = ?)"}).AddRow(1)
	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM sections WHERE id = \\?\\)").WithArgs(1).WillReturnRows(existsRows)

	// Mock para obtener el reporte de la sección específica (error)
	mock.ExpectQuery("SELECT s\\.id as section_id, s\\.section_number, COUNT\\(pb\\.id\\) as products_count FROM sections s LEFT JOIN product_batches pb ON s\\.id = pb\\.section_id WHERE s\\.id = \\? GROUP BY s\\.id, s\\.section_number").WithArgs(1).WillReturnError(sql.ErrConnDone)

	// Act
	result, err := repository.GetProductCountBySection(&sectionID)

	// Assert
	require.Error(t, err)
	require.Nil(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestProductBatchMySQL_GetProductCountBySection_ScanError(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupProductBatchMySQLTest(t)
	defer cleanup()

	// Simular un error de scan usando una columna incorrecta
	rows := sqlmock.NewRows([]string{"section_id", "section_number", "products_count"}).
		AddRow("invalid_id", "SEC001", 5) // section_id como string en lugar de int

	mock.ExpectQuery("SELECT s\\.id as section_id, s\\.section_number, COUNT\\(pb\\.id\\) as products_count FROM sections s LEFT JOIN product_batches pb ON s\\.id = pb\\.section_id GROUP BY s\\.id, s\\.section_number").WillReturnRows(rows)

	// Act
	result, err := repository.GetProductCountBySection(nil)

	// Assert
	require.Error(t, err)
	require.Nil(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestProductBatchMySQL_GetProductCountBySection_EmptyResult_Success(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupProductBatchMySQLTest(t)
	defer cleanup()

	// No hay resultados
	rows := sqlmock.NewRows([]string{"section_id", "section_number", "products_count"})

	mock.ExpectQuery("SELECT s\\.id as section_id, s\\.section_number, COUNT\\(pb\\.id\\) as products_count FROM sections s LEFT JOIN product_batches pb ON s\\.id = pb\\.section_id GROUP BY s\\.id, s\\.section_number").WillReturnRows(rows)

	// Act
	result, err := repository.GetProductCountBySection(nil)

	// Assert
	require.NoError(t, err)
	require.Len(t, result, 0)
	require.NoError(t, mock.ExpectationsWereMet())
}
