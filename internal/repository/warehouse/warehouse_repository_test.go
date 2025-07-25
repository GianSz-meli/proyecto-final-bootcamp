package warehouse

import (
	"database/sql"
	"errors"
	"testing"

	customErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestGetAll_Success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewSqlWarehouseRepository(db)

	expectedWarehouses := []models.Warehouse{
		{
			ID:                 1,
			WarehouseCode:      "WH001",
			Address:            "Address 1",
			Telephone:          "1234567890",
			MinimumCapacity:    100,
			MinimumTemperature: -10.5,
			LocalityId:         &[]int{1}[0],
			Locality: &models.Locality{
				Id:           1,
				LocalityName: "Buenos Aires",
				Province: models.Province{
					Id:           1,
					ProvinceName: "Buenos Aires Province",
					Country: models.Country{
						Id:          1,
						CountryName: "Argentina",
					},
				},
			},
		},
		{
			ID:                 2,
			WarehouseCode:      "WH002",
			Address:            "Address 2",
			Telephone:          "0987654321",
			MinimumCapacity:    200,
			MinimumTemperature: -5.0,
			LocalityId:         nil,
			Locality:           nil,
		},
	}

	rows := sqlmock.NewRows([]string{
		"id", "warehouse_code", "address", "telephone", "minimum_capacity", "minimum_temperature", "locality_id",
		"locality_id", "locality_name", "province_id", "province_name", "country_id", "country_name",
	}).
		AddRow(1, "WH001", "Address 1", "1234567890", 100, -10.5, 1, 1, "Buenos Aires", 1, "Buenos Aires Province", 1, "Argentina").
		AddRow(2, "WH002", "Address 2", "0987654321", 200, -5.0, nil, nil, nil, nil, nil, nil, nil)

	mock.ExpectQuery("^SELECT (.+) FROM warehouses").
		WillReturnRows(rows)

	result, err := repo.GetAll()

	require.NoError(t, err)
	require.Equal(t, expectedWarehouses, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAll_ErrorWhenQuery_ReturnsError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewSqlWarehouseRepository(db)

	expectedError := errors.New("connection with db was broken and failed")
	mock.ExpectQuery("^SELECT (.+) FROM warehouses").
		WillReturnError(expectedError)

	result, err := repo.GetAll()

	require.Error(t, err)
	require.Equal(t, expectedError, err)
	require.Nil(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAll_RowsScanError_ReturnsError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewSqlWarehouseRepository(db)

	// Create rows with incompatible data type that will cause Scan to fail (invalid_id is not an int)
	rows := sqlmock.NewRows([]string{
		"id", "warehouse_code", "address", "telephone", "minimum_capacity", "minimum_temperature", "locality_id",
		"locality_id", "locality_name", "province_id", "province_name", "country_id", "country_name",
	}).
		AddRow("invalid_id", "WH001", "Address 1", "1234567890", 100, -10.5, 1, 1, "Buenos Aires", 1, "Buenos Aires Province", 1, "Argentina")

	mock.ExpectQuery("^SELECT (.+) FROM warehouses").
		WillReturnRows(rows)

	result, err := repo.GetAll()

	require.Error(t, err)
	require.Nil(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAll_RowsErrError_ReturnsError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewSqlWarehouseRepository(db)

	expectedError := errors.New("rows iteration error")

	// Create valid rows but set an error that will be returned by rows.Err()
	rows := sqlmock.NewRows([]string{
		"id", "warehouse_code", "address", "telephone", "minimum_capacity", "minimum_temperature", "locality_id",
		"locality_id", "locality_name", "province_id", "province_name", "country_id", "country_name",
	}).
		AddRow(1, "WH001", "Address 1", "1234567890", 100, -10.5, 1, 1, "Buenos Aires", 1, "Buenos Aires Province", 1, "Argentina").
		CloseError(expectedError)

	mock.ExpectQuery("^SELECT (.+) FROM warehouses").
		WillReturnRows(rows)

	result, err := repo.GetAll()

	require.Error(t, err)
	require.Equal(t, expectedError, err)
	require.Nil(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetById_Success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewSqlWarehouseRepository(db)

	expectedWarehouse := models.Warehouse{
		ID:                 1,
		WarehouseCode:      "WH001",
		Address:            "Address 1",
		Telephone:          "1234567890",
		MinimumCapacity:    100,
		MinimumTemperature: -10.5,
		LocalityId:         &[]int{1}[0],
		Locality: &models.Locality{
			Id:           1,
			LocalityName: "Buenos Aires",
			Province: models.Province{
				Id:           1,
				ProvinceName: "Buenos Aires Province",
				Country: models.Country{
					Id:          1,
					CountryName: "Argentina",
				},
			},
		},
	}

	rows := sqlmock.NewRows([]string{
		"id", "warehouse_code", "address", "telephone", "minimum_capacity", "minimum_temperature", "locality_id",
		"locality_id", "locality_name", "province_id", "province_name", "country_id", "country_name",
	}).
		AddRow(1, "WH001", "Address 1", "1234567890", 100, -10.5, 1, 1, "Buenos Aires", 1, "Buenos Aires Province", 1, "Argentina")

	mock.ExpectQuery("^SELECT (.+) FROM warehouses").
		WithArgs(1).
		WillReturnRows(rows)

	result, err := repo.GetById(1)

	require.NoError(t, err)
	require.Equal(t, &expectedWarehouse, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetById_ErrorWhenQuery_ReturnsError(t *testing.T) {
	genericError := errors.New("connection with db was broken and failed")
	testCases := []struct {
		name          string
		mockRows      *sqlmock.Rows
		mockError     error
		expectedError error
	}{
		{
			name:          "Returns error when query fails because of unexpected error",
			mockRows:      sqlmock.NewRows([]string{}),
			mockError:     genericError,
			expectedError: genericError,
		},
		{
			name:          "Should return ErrNotFound when query fails because of no rows were found",
			mockRows:      sqlmock.NewRows([]string{}),
			mockError:     sql.ErrNoRows,
			expectedError: customErrors.ErrNotFound,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New()
			defer db.Close()

			repo := NewSqlWarehouseRepository(db)

			mock.ExpectQuery("^SELECT (.+) FROM warehouses").
				WithArgs(1).
				WillReturnRows(testCase.mockRows).
				WillReturnError(testCase.mockError)

			result, err := repo.GetById(1)

			require.Error(t, err)
			require.ErrorIs(t, err, testCase.expectedError)
			require.Nil(t, result)
			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestCreate_Success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewSqlWarehouseRepository(db)

	inputWarehouse := models.Warehouse{
		WarehouseCode:      "WH001",
		Address:            "Address 1",
		Telephone:          "1234567890",
		MinimumCapacity:    100,
		MinimumTemperature: -10.5,
		LocalityId:         &[]int{1}[0],
	}

	expectedWarehouse := inputWarehouse
	expectedWarehouse.ID = 1

	mock.ExpectExec("^INSERT INTO warehouses").
		WithArgs("WH001", "Address 1", "1234567890", 100, -10.5, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := repo.Create(inputWarehouse)

	require.NoError(t, err)
	require.Equal(t, expectedWarehouse, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestCreate_ErrorWhenQuery_ReturnsError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewSqlWarehouseRepository(db)

	inputWarehouse := models.Warehouse{
		WarehouseCode:      "WH001",
		Address:            "Address 1",
		Telephone:          "1234567890",
		MinimumCapacity:    100,
		MinimumTemperature: -10.5,
		LocalityId:         &[]int{1}[0],
	}

	expectedError := errors.New("connection with db was broken and failed")
	mock.ExpectExec("^INSERT INTO warehouses").
		WithArgs("WH001", "Address 1", "1234567890", 100, -10.5, 1).
		WillReturnError(expectedError)

	result, err := repo.Create(inputWarehouse)

	require.Error(t, err)
	require.Equal(t, expectedError, err)
	require.Equal(t, models.Warehouse{}, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestCreate_LastInsertIdError_ReturnsError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewSqlWarehouseRepository(db)

	inputWarehouse := models.Warehouse{
		WarehouseCode:      "WH001",
		Address:            "Address 1",
		Telephone:          "1234567890",
		MinimumCapacity:    100,
		MinimumTemperature: -10.5,
		LocalityId:         &[]int{1}[0],
	}
	expectedError := errors.New("last insert id failed")

	mock.ExpectExec("^INSERT INTO warehouses").
		WithArgs("WH001", "Address 1", "1234567890", 100, -10.5, 1).
		WillReturnResult(sqlmock.NewErrorResult(expectedError))

	result, err := repo.Create(inputWarehouse)

	require.Error(t, err)
	require.Equal(t, expectedError, err)
	require.Equal(t, models.Warehouse{}, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdate_Success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewSqlWarehouseRepository(db)

	expectedWarehouse := models.Warehouse{
		ID:                 1,
		WarehouseCode:      "WH001-UPDATED",
		Address:            "Address Updated",
		Telephone:          "9876543210",
		MinimumCapacity:    200,
		MinimumTemperature: -15.0,
		LocalityId:         &[]int{2}[0],
	}

	mock.ExpectExec("UPDATE warehouses SET (.+) WHERE id = (.+)").
		WithArgs("WH001-UPDATED", "Address Updated", "9876543210", 200, -15.0, 2, 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	result, err := repo.Update(1, expectedWarehouse)

	require.NoError(t, err)
	require.Equal(t, expectedWarehouse, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdate_ErrorWhenQuery_ReturnsError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewSqlWarehouseRepository(db)

	inputWarehouse := models.Warehouse{
		ID:                 1,
		WarehouseCode:      "WH001-UPDATED",
		Address:            "Address Updated",
		Telephone:          "9876543210",
		MinimumCapacity:    200,
		MinimumTemperature: -15.0,
		LocalityId:         &[]int{2}[0],
	}

	expectedError := errors.New("connection with db was broken and failed")
	mock.ExpectExec("UPDATE warehouses SET (.+) WHERE id = (.+)").
		WithArgs("WH001-UPDATED", "Address Updated", "9876543210", 200, -15.0, 2, 1).
		WillReturnError(expectedError)

	result, err := repo.Update(1, inputWarehouse)

	require.Error(t, err)
	require.Equal(t, expectedError, err)
	require.Equal(t, models.Warehouse{}, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDelete_Success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewSqlWarehouseRepository(db)

	mock.ExpectExec("DELETE FROM warehouses WHERE id = ?").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Delete(1)
	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDelete_ErrorWhenQuery_ReturnsError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewSqlWarehouseRepository(db)

	expectedError := errors.New("connection with db was broken and failed")
	mock.ExpectExec("DELETE FROM warehouses WHERE id = ?").
		WithArgs(1).
		WillReturnError(expectedError)

	err := repo.Delete(1)

	require.Error(t, err)
	require.Equal(t, expectedError, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDelete_RowsAffectedError_ReturnsError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewSqlWarehouseRepository(db)

	expectedError := errors.New("rows affected failed")
	mock.ExpectExec("DELETE FROM warehouses WHERE id = ?").
		WithArgs(1).
		WillReturnResult(sqlmock.NewErrorResult(expectedError))

	err := repo.Delete(1)

	require.Error(t, err)
	require.Equal(t, expectedError, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDelete_NoRowsAffected_ReturnsErrNotFound(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewSqlWarehouseRepository(db)

	mock.ExpectExec("DELETE FROM warehouses WHERE id = ?").
		WithArgs(999).
		WillReturnResult(sqlmock.NewResult(0, 0)) // 0 rows affected

	err := repo.Delete(999)

	require.Error(t, err)
	require.ErrorIs(t, err, customErrors.ErrNotFound)
	require.NoError(t, mock.ExpectationsWereMet())
}
