package repository

import (
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func setupSectionMySQLTest(t *testing.T) (*SectionMySQL, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	repository := NewSectionMySQL(db)

	cleanup := func() {
		db.Close()
	}

	return repository, mock, cleanup
}

func TestSectionMySQL_GetAll_Success(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupSectionMySQLTest(t)
	defer cleanup()

	rows := sqlmock.NewRows([]string{
		"id", "section_number", "current_temperature", "minimum_temperature",
		"current_capacity", "minimum_capacity", "product_type_id", "warehouse_id", "maximum_capacity",
	}).AddRow(
		1, "SEC001", 15.5, 10.0, 50, 20, 1, 1, 100,
	).AddRow(
		2, "SEC002", 20.0, 15.0, 75, 30, 2, 1, 150,
	)

	mock.ExpectQuery("SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, product_type_id, warehouse_id, maximum_capacity FROM sections").WillReturnRows(rows)

	// Act
	result, err := repository.GetAll()

	// Assert
	require.NoError(t, err)
	require.Len(t, result, 2)
	require.Equal(t, 1, result[0].ID)
	require.Equal(t, "SEC001", result[0].SectionNumber)
	require.Equal(t, 2, result[1].ID)
	require.Equal(t, "SEC002", result[1].SectionNumber)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSectionMySQL_GetAll_Error(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupSectionMySQLTest(t)
	defer cleanup()

	mock.ExpectQuery("SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, product_type_id, warehouse_id, maximum_capacity FROM sections").WillReturnError(sql.ErrConnDone)

	// Act
	result, err := repository.GetAll()

	// Assert
	require.Error(t, err)
	require.Nil(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSectionMySQL_GetAll_ScanError(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupSectionMySQLTest(t)
	defer cleanup()

	// Simular un error de scan usando una columna incorrecta
	rows := sqlmock.NewRows([]string{
		"id", "section_number", "current_temperature", "minimum_temperature",
		"current_capacity", "minimum_capacity", "product_type_id", "warehouse_id", "maximum_capacity",
	}).AddRow(
		"invalid_id", "SEC001", 15.5, 10.0, 50, 20, 1, 1, 100, // ID como string en lugar de int
	)

	mock.ExpectQuery("SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, product_type_id, warehouse_id, maximum_capacity FROM sections").WillReturnRows(rows)

	// Act
	result, err := repository.GetAll()

	// Assert
	require.Error(t, err)
	require.Nil(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSectionMySQL_GetById_Success(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupSectionMySQLTest(t)
	defer cleanup()

	rows := sqlmock.NewRows([]string{
		"id", "section_number", "current_temperature", "minimum_temperature",
		"current_capacity", "minimum_capacity", "product_type_id", "warehouse_id", "maximum_capacity",
	}).AddRow(
		1, "SEC001", 15.5, 10.0, 50, 20, 1, 1, 100,
	)

	mock.ExpectQuery("SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, product_type_id, warehouse_id, maximum_capacity FROM sections WHERE id = \\?").WithArgs(1).WillReturnRows(rows)

	// Act
	result, err := repository.GetById(1)

	// Assert
	require.NoError(t, err)
	require.Equal(t, 1, result.ID)
	require.Equal(t, "SEC001", result.SectionNumber)
	require.Equal(t, 15.5, result.CurrentTemperature)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSectionMySQL_GetById_NotFound(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupSectionMySQLTest(t)
	defer cleanup()

	mock.ExpectQuery("SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, product_type_id, warehouse_id, maximum_capacity FROM sections WHERE id = \\?").WithArgs(999).WillReturnError(sql.ErrNoRows)

	// Act
	result, err := repository.GetById(999)

	// Assert
	require.Error(t, err)
	require.Equal(t, models.Section{}, result)
	require.ErrorIs(t, err, errors.ErrNotFound)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSectionMySQL_GetById_DatabaseError(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupSectionMySQLTest(t)
	defer cleanup()

	mock.ExpectQuery("SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, product_type_id, warehouse_id, maximum_capacity FROM sections WHERE id = \\?").WithArgs(1).WillReturnError(sql.ErrConnDone)

	// Act
	result, err := repository.GetById(1)

	// Assert
	require.Error(t, err)
	require.Equal(t, models.Section{}, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSectionMySQL_Create_Success(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupSectionMySQLTest(t)
	defer cleanup()

	section := models.Section{
		SectionAttributes: models.SectionAttributes{
			SectionNumber:      "SEC001",
			CurrentTemperature: 15.5,
			MinimumTemperature: 10.0,
			CurrentCapacity:    50,
			MinimumCapacity:    20,
			MaximumCapacity:    100,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO sections \\(section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, product_type_id, warehouse_id, maximum_capacity\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?\\)").WithArgs(
		"SEC001", 15.5, 10.0, 50, 20, 1, 1, 100,
	).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Act
	result, err := repository.Create(section)

	// Assert
	require.NoError(t, err)
	require.Equal(t, 1, result.ID)
	require.Equal(t, "SEC001", result.SectionNumber)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSectionMySQL_Create_BeginTransactionError(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupSectionMySQLTest(t)
	defer cleanup()

	section := models.Section{
		SectionAttributes: models.SectionAttributes{
			SectionNumber:      "SEC001",
			CurrentTemperature: 15.5,
			MinimumTemperature: 10.0,
			CurrentCapacity:    50,
			MinimumCapacity:    20,
			MaximumCapacity:    100,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
	}

	mock.ExpectBegin().WillReturnError(sql.ErrConnDone)

	// Act
	result, err := repository.Create(section)

	// Assert
	require.Error(t, err)
	require.Equal(t, section, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSectionMySQL_Create_ExecError(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupSectionMySQLTest(t)
	defer cleanup()

	section := models.Section{
		SectionAttributes: models.SectionAttributes{
			SectionNumber:      "SEC001",
			CurrentTemperature: 15.5,
			MinimumTemperature: 10.0,
			CurrentCapacity:    50,
			MinimumCapacity:    20,
			MaximumCapacity:    100,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO sections \\(section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, product_type_id, warehouse_id, maximum_capacity\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?\\)").WithArgs(
		"SEC001", 15.5, 10.0, 50, 20, 1, 1, 100,
	).WillReturnError(sql.ErrConnDone)
	mock.ExpectRollback()

	// Act
	result, err := repository.Create(section)

	// Assert
	require.Error(t, err)
	require.Equal(t, section, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSectionMySQL_Create_LastInsertIdError(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupSectionMySQLTest(t)
	defer cleanup()

	section := models.Section{
		SectionAttributes: models.SectionAttributes{
			SectionNumber:      "SEC001",
			CurrentTemperature: 15.5,
			MinimumTemperature: 10.0,
			CurrentCapacity:    50,
			MinimumCapacity:    20,
			MaximumCapacity:    100,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO sections \\(section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, product_type_id, warehouse_id, maximum_capacity\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?\\)").WithArgs(
		"SEC001", 15.5, 10.0, 50, 20, 1, 1, 100,
	).WillReturnResult(sqlmock.NewErrorResult(sql.ErrConnDone))
	mock.ExpectRollback()

	// Act
	result, err := repository.Create(section)

	// Assert
	require.Error(t, err)
	require.Equal(t, section, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSectionMySQL_Create_CommitError(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupSectionMySQLTest(t)
	defer cleanup()

	section := models.Section{
		SectionAttributes: models.SectionAttributes{
			SectionNumber:      "SEC001",
			CurrentTemperature: 15.5,
			MinimumTemperature: 10.0,
			CurrentCapacity:    50,
			MinimumCapacity:    20,
			MaximumCapacity:    100,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
	}

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO sections \\(section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, product_type_id, warehouse_id, maximum_capacity\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?, \\?, \\?\\)").WithArgs(
		"SEC001", 15.5, 10.0, 50, 20, 1, 1, 100,
	).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit().WillReturnError(sql.ErrConnDone)

	// Act
	result, err := repository.Create(section)

	// Assert
	require.Error(t, err)
	// En caso de error en commit, el ID ya se asignó antes del commit
	require.Equal(t, 1, result.ID)
	require.Equal(t, "SEC001", result.SectionNumber)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSectionMySQL_Update_Success(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupSectionMySQLTest(t)
	defer cleanup()

	section := models.Section{
		SectionAttributes: models.SectionAttributes{
			SectionNumber:      "SEC001-UPDATED",
			CurrentTemperature: 20.0,
			MinimumTemperature: 10.0,
			CurrentCapacity:    50,
			MinimumCapacity:    20,
			MaximumCapacity:    150,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
	}

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE sections SET section_number=\\?, current_temperature=\\?, minimum_temperature=\\?, current_capacity=\\?, minimum_capacity=\\?, product_type_id=\\?, warehouse_id=\\?, maximum_capacity=\\? WHERE id=\\?").WithArgs(
		"SEC001-UPDATED", 20.0, 10.0, 50, 20, 1, 1, 150, 1,
	).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	// Act
	result, err := repository.Update(1, section)

	// Assert
	require.NoError(t, err)
	require.Equal(t, 1, result.ID)
	require.Equal(t, "SEC001-UPDATED", result.SectionNumber)
	require.Equal(t, 20.0, result.CurrentTemperature)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSectionMySQL_Update_BeginTransactionError(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupSectionMySQLTest(t)
	defer cleanup()

	section := models.Section{
		SectionAttributes: models.SectionAttributes{
			SectionNumber:      "SEC001-UPDATED",
			CurrentTemperature: 20.0,
			MinimumTemperature: 10.0,
			CurrentCapacity:    50,
			MinimumCapacity:    20,
			MaximumCapacity:    150,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
	}

	mock.ExpectBegin().WillReturnError(sql.ErrConnDone)

	// Act
	result, err := repository.Update(1, section)

	// Assert
	require.Error(t, err)
	require.Equal(t, section, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSectionMySQL_Update_ExecError(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupSectionMySQLTest(t)
	defer cleanup()

	section := models.Section{
		SectionAttributes: models.SectionAttributes{
			SectionNumber:      "SEC001-UPDATED",
			CurrentTemperature: 20.0,
			MinimumTemperature: 10.0,
			CurrentCapacity:    50,
			MinimumCapacity:    20,
			MaximumCapacity:    150,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
	}

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE sections SET section_number=\\?, current_temperature=\\?, minimum_temperature=\\?, current_capacity=\\?, minimum_capacity=\\?, product_type_id=\\?, warehouse_id=\\?, maximum_capacity=\\? WHERE id=\\?").WithArgs(
		"SEC001-UPDATED", 20.0, 10.0, 50, 20, 1, 1, 150, 1,
	).WillReturnError(sql.ErrConnDone)
	mock.ExpectRollback()

	// Act
	result, err := repository.Update(1, section)

	// Assert
	require.Error(t, err)
	require.Equal(t, section, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSectionMySQL_Update_CommitError(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupSectionMySQLTest(t)
	defer cleanup()

	section := models.Section{
		SectionAttributes: models.SectionAttributes{
			SectionNumber:      "SEC001-UPDATED",
			CurrentTemperature: 20.0,
			MinimumTemperature: 10.0,
			CurrentCapacity:    50,
			MinimumCapacity:    20,
			MaximumCapacity:    150,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
	}

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE sections SET section_number=\\?, current_temperature=\\?, minimum_temperature=\\?, current_capacity=\\?, minimum_capacity=\\?, product_type_id=\\?, warehouse_id=\\?, maximum_capacity=\\? WHERE id=\\?").WithArgs(
		"SEC001-UPDATED", 20.0, 10.0, 50, 20, 1, 1, 150, 1,
	).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit().WillReturnError(sql.ErrConnDone)

	// Act
	result, err := repository.Update(1, section)

	// Assert
	require.Error(t, err)
	require.Equal(t, section, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSectionMySQL_Delete_Success(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupSectionMySQLTest(t)
	defer cleanup()

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM sections WHERE id=\\?").WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	// Act
	err := repository.Delete(1)

	// Assert
	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSectionMySQL_Delete_BeginTransactionError(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupSectionMySQLTest(t)
	defer cleanup()

	mock.ExpectBegin().WillReturnError(sql.ErrConnDone)

	// Act
	err := repository.Delete(1)

	// Assert
	require.Error(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSectionMySQL_Delete_ExecError(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupSectionMySQLTest(t)
	defer cleanup()

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM sections WHERE id=\\?").WithArgs(1).WillReturnError(sql.ErrConnDone)
	mock.ExpectRollback()

	// Act
	err := repository.Delete(1)

	// Assert
	require.Error(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSectionMySQL_Delete_CommitError(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupSectionMySQLTest(t)
	defer cleanup()

	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM sections WHERE id=\\?").WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit().WillReturnError(sql.ErrConnDone)

	// Act
	err := repository.Delete(1)

	// Assert
	require.Error(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSectionMySQL_ExistBySectionNumber_Exists(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupSectionMySQLTest(t)
	defer cleanup()

	rows := sqlmock.NewRows([]string{"EXISTS(SELECT 1 FROM sections WHERE section_number = ?)"}).AddRow(1)

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM sections WHERE section_number = \\?\\)").WithArgs("SEC001").WillReturnRows(rows)

	// Act
	result := repository.ExistBySectionNumber("SEC001")

	// Assert
	require.True(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSectionMySQL_ExistBySectionNumber_NotExists(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupSectionMySQLTest(t)
	defer cleanup()

	rows := sqlmock.NewRows([]string{"EXISTS(SELECT 1 FROM sections WHERE section_number = ?)"}).AddRow(0)

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM sections WHERE section_number = \\?\\)").WithArgs("SEC999").WillReturnRows(rows)

	// Act
	result := repository.ExistBySectionNumber("SEC999")

	// Assert
	require.False(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSectionMySQL_ExistBySectionNumber_Error(t *testing.T) {
	// Arrange
	repository, mock, cleanup := setupSectionMySQLTest(t)
	defer cleanup()

	mock.ExpectQuery("SELECT EXISTS\\(SELECT 1 FROM sections WHERE section_number = \\?\\)").WithArgs("SEC001").WillReturnError(sql.ErrConnDone)

	// Act
	result := repository.ExistBySectionNumber("SEC001")

	// Assert
	require.False(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}
