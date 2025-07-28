package locality

import (
	pkgErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestLocalityMysql_Create_QueryError(t *testing.T) {
	//Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := NewLocalityMysqlRepository(db)
	locality := models.Locality{
		LocalityName: "test",
		Province: models.Province{
			ProvinceName: "Lombardy",
			Country: models.Country{
				CountryName: "Italy",
			},
		},
	}
	expectedError := errors.New("error sql query")
	mock.ExpectExec("^INSERT INTO localities").
		WithArgs(
			locality.Id,
			locality.LocalityName,
			locality.Province.ProvinceName,
			locality.Province.Country.CountryName,
		).
		WillReturnError(expectedError)
	//Act
	result, err := repo.Create(locality)
	//Assert
	require.Equal(t, models.Locality{}, result)
	require.Error(t, err)
	require.EqualError(t, expectedError, err.Error())
	require.NoError(t, mock.ExpectationsWereMet())
}
func TestLocalityMysql_Create_GetLastInsertError(t *testing.T) {
	//Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := NewLocalityMysqlRepository(db)
	locality := models.Locality{
		LocalityName: "test",
		Province: models.Province{
			ProvinceName: "Lombardy",
			Country: models.Country{
				CountryName: "Italy",
			},
		},
	}
	expectedError := errors.New("error")
	mock.ExpectExec("^INSERT INTO localities").
		WithArgs(
			locality.Id,
			locality.LocalityName,
			locality.Province.ProvinceName,
			locality.Province.Country.CountryName,
		).WillReturnResult(sqlmock.NewErrorResult(expectedError))
	//Act
	result, err := repo.Create(locality)
	//Assert
	require.Equal(t, models.Locality{}, result)
	require.Error(t, err)
	require.EqualError(t, expectedError, err.Error())
	require.NoError(t, mock.ExpectationsWereMet())
}
func TestLocalityMysql_Create_Success(t *testing.T) {
	//Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := NewLocalityMysqlRepository(db)
	locality := models.Locality{
		LocalityName: "test",
		Province: models.Province{
			ProvinceName: "Lombardy",
			Country: models.Country{
				CountryName: "Italy",
			},
		},
	}
	newLocalityId := 1
	newLocality := locality
	newLocality.Id = newLocalityId
	mock.ExpectExec("^INSERT INTO localities").
		WithArgs(
			locality.Id,
			locality.LocalityName,
			locality.Province.ProvinceName,
			locality.Province.Country.CountryName,
		).WillReturnResult(sqlmock.NewResult(int64(newLocalityId), 1))
	//Act
	result, err := repo.Create(locality)
	//Assert
	require.Equal(t, newLocality, result)
	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestLocalityMysql_GetSellersByIdLocality_QueryError(t *testing.T) {
	//Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := NewLocalityMysqlRepository(db)
	localityId := 1
	expectedError := errors.New("error sql query")

	mock.ExpectQuery(`^SELECT COUNT\(\*\) as sellers_count, l.id, l.locality_name FROM sellers s JOIN localities l ON s.locality_id = l.id WHERE l.id = \? GROUP BY l.id`).
		WithArgs(localityId).
		WillReturnError(expectedError)
	//Act
	result, err := repo.GetSellersByIdLocality(localityId)
	//Assert
	require.Equal(t, models.SellersByLocalityReport{}, result)
	require.Error(t, err)
	require.EqualError(t, err, expectedError.Error())
	require.NoError(t, mock.ExpectationsWereMet())
}
func TestLocalityMysql_GetSellersByIdLocality_NotFoundError(t *testing.T) {
	//Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := NewLocalityMysqlRepository(db)
	localityId := 1
	rows := sqlmock.NewRows([]string{
		"sellers_count", "locality_id", "locality_name",
	})
	mock.ExpectQuery(`^SELECT COUNT\(\*\) as sellers_count, l.id, l.locality_name FROM sellers s JOIN localities l ON s.locality_id = l.id WHERE l.id = \? GROUP BY l.id`).
		WithArgs(localityId).WillReturnRows(rows)
	//Act
	result, err := repo.GetSellersByIdLocality(localityId)
	//Assert
	require.Equal(t, models.SellersByLocalityReport{}, result)
	require.Error(t, err)
	require.ErrorIs(t, err, pkgErrors.ErrNotFound)
	require.NoError(t, mock.ExpectationsWereMet())
}
func TestLocalityMysql_GetSellersByIdLocality_ScannerError(t *testing.T) {
	//Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := NewLocalityMysqlRepository(db)
	localityId := 1
	rows := sqlmock.NewRows([]string{
		"sellers_count", "locality_id", "locality_name",
	}).AddRow("NOT INT", 1, "test")

	mock.ExpectQuery(`^SELECT COUNT\(\*\) as sellers_count, l.id, l.locality_name FROM sellers s JOIN localities l ON s.locality_id = l.id WHERE l.id = \? GROUP BY l.id`).
		WithArgs(localityId).WillReturnRows(rows)
	//Act
	result, err := repo.GetSellersByIdLocality(localityId)
	//Assert
	require.Equal(t, models.SellersByLocalityReport{}, result)
	require.Error(t, err)
	require.Contains(t, err.Error(), "Scan error")
	require.NoError(t, mock.ExpectationsWereMet())
}
func TestLocalityMysql_GetSellersByIdLocality_Success(t *testing.T) {
	//Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := NewLocalityMysqlRepository(db)
	localityId := 1
	expectedLocalityReport := models.SellersByLocalityReport{
		LocalityId:   localityId,
		LocalityName: "La Plata",
		SellersCount: 3,
	}
	rows := sqlmock.NewRows([]string{
		"sellers_count", "locality_id", "locality_name",
	}).AddRow(expectedLocalityReport.SellersCount, expectedLocalityReport.LocalityId, expectedLocalityReport.LocalityName)

	mock.ExpectQuery(`^SELECT COUNT\(\*\) as sellers_count, l.id, l.locality_name FROM sellers s JOIN localities l ON s.locality_id = l.id WHERE l.id = \? GROUP BY l.id`).
		WithArgs(localityId).
		WillReturnRows(rows)
	//Act
	result, err := repo.GetSellersByIdLocality(localityId)
	//Assert
	require.Equal(t, expectedLocalityReport, result)
	require.Nil(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestLocalityMysql_GetSellersByLocalities_QueryError(t *testing.T) {
	//Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := NewLocalityMysqlRepository(db)
	expectedError := errors.New("error sql query")

	mock.ExpectQuery(`^SELECT COUNT\(\*\) as sellers_count, l.id, l.locality_name FROM sellers s JOIN localities l ON s.locality_id = l.id GROUP BY l.id`).
		WillReturnError(expectedError)
	//Act
	result, err := repo.GetSellersByLocalities()
	//Assert
	require.Nil(t, result)
	require.Error(t, err)
	require.EqualError(t, err, expectedError.Error())
	require.NoError(t, mock.ExpectationsWereMet())
}
func TestLocalityMysql_GetSellersByLocalities_RowsError(t *testing.T) {
	//Arrange
	expectedLocalitiesReport := []models.SellersByLocalityReport{
		{
			LocalityId:   1,
			LocalityName: "La Plata",
			SellersCount: 3,
		},
		{
			LocalityId:   2,
			LocalityName: "Buenos Aires",
			SellersCount: 6,
		},
	}
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := NewLocalityMysqlRepository(db)
	expectedError := errors.New("error scanning rows")

	rows := sqlmock.NewRows([]string{
		"sellers_count", "locality_id", "locality_name",
	}).
		AddRow(expectedLocalitiesReport[0].SellersCount, expectedLocalitiesReport[0].LocalityId, expectedLocalitiesReport[0].LocalityName).
		AddRow(expectedLocalitiesReport[1].SellersCount, expectedLocalitiesReport[1].LocalityId, expectedLocalitiesReport[1].LocalityName).
		RowError(0, expectedError)
	mock.ExpectQuery(`^SELECT COUNT\(\*\) as sellers_count, l.id, l.locality_name FROM sellers s JOIN localities l ON s.locality_id = l.id GROUP BY l.id`).
		WillReturnRows(rows)
	//Act
	result, err := repo.GetSellersByLocalities()
	//Assert
	require.Nil(t, result)
	require.Error(t, err)
	require.EqualError(t, expectedError, err.Error())
	require.NoError(t, mock.ExpectationsWereMet())
}
func TestLocalityMysql_GetSellersByLocalities_ScannerError(t *testing.T) {
	//Arrange
	expectedLocalitiesReport := []models.SellersByLocalityReport{
		{
			LocalityId:   1,
			LocalityName: "La Plata",
			SellersCount: 3,
		},
		{
			LocalityId:   2,
			LocalityName: "Buenos Aires",
			SellersCount: 6,
		},
	}
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := NewLocalityMysqlRepository(db)

	rows := sqlmock.NewRows([]string{
		"sellers_count", "locality_id", "locality_name",
	}).
		AddRow(expectedLocalitiesReport[0].SellersCount, expectedLocalitiesReport[0].LocalityId, expectedLocalitiesReport[0].LocalityName).
		AddRow("NOT INT", expectedLocalitiesReport[1].LocalityId, expectedLocalitiesReport[1].LocalityName)
	mock.ExpectQuery(`^SELECT COUNT\(\*\) as sellers_count, l.id, l.locality_name FROM sellers s JOIN localities l ON s.locality_id = l.id GROUP BY l.id`).
		WillReturnRows(rows)
	//Act
	result, err := repo.GetSellersByLocalities()
	//Assert

	require.Nil(t, result)
	require.Error(t, err)
	require.Contains(t, err.Error(), "Scan error")
	require.NoError(t, mock.ExpectationsWereMet())
}
func TestLocalityMysql_GetSellersByLocalities_Success(t *testing.T) {
	//Arrange
	expectedLocalitiesReport := []models.SellersByLocalityReport{
		{
			LocalityId:   1,
			LocalityName: "La Plata",
			SellersCount: 3,
		},
		{
			LocalityId:   2,
			LocalityName: "Buenos Aires",
			SellersCount: 6,
		},
	}
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := NewLocalityMysqlRepository(db)

	rows := sqlmock.NewRows([]string{
		"sellers_count", "locality_id", "locality_name",
	}).
		AddRow(expectedLocalitiesReport[0].SellersCount, expectedLocalitiesReport[0].LocalityId, expectedLocalitiesReport[0].LocalityName).
		AddRow(expectedLocalitiesReport[1].SellersCount, expectedLocalitiesReport[1].LocalityId, expectedLocalitiesReport[1].LocalityName)

	mock.ExpectQuery(`^SELECT COUNT\(\*\) as sellers_count, l.id, l.locality_name FROM sellers s JOIN localities l ON s.locality_id = l.id GROUP BY l.id`).
		WillReturnRows(rows)
	//Act
	result, err := repo.GetSellersByLocalities()
	//Assert

	require.Nil(t, err)
	require.Equal(t, expectedLocalitiesReport, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestReportCarriersByLocality_Success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	repo := NewLocalityMysqlRepository(db)

	localityId := 1
	expectedReports := []models.CarrierReport{
		{
			LocalityId:    1,
			LocalityName:  "Buenos Aires",
			CarriersCount: 3,
		},
	}
	rows := sqlmock.NewRows([]string{
		"id", "locality_name", "carriers_count",
	}).
		AddRow(1, "Buenos Aires", 3)

	mock.ExpectQuery("^SELECT (.+) FROM localities l (.+)").
		WithArgs(localityId, localityId).
		WillReturnRows(rows)

	result, err := repo.ReportCarriersByLocality(&localityId)

	require.NoError(t, err)
	require.Equal(t, expectedReports, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestReportCarriersByLocality_LocalityIdIsNil_ReturnsAllLocalities(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	repo := NewLocalityMysqlRepository(db)

	expectedReports := []models.CarrierReport{
		{
			LocalityId:    1,
			LocalityName:  "Buenos Aires",
			CarriersCount: 3,
		},
		{
			LocalityId:    2,
			LocalityName:  "Córdoba",
			CarriersCount: 1,
		},
		{
			LocalityId:    3,
			LocalityName:  "Rosario",
			CarriersCount: 2,
		},
	}
	rows := sqlmock.NewRows([]string{
		"id", "locality_name", "carriers_count",
	}).
		AddRow(1, "Buenos Aires", 3).
		AddRow(2, "Córdoba", 1).
		AddRow(3, "Rosario", 2)

	mock.ExpectQuery("^SELECT (.+) FROM localities l (.+)").
		WithArgs(nil, nil).
		WillReturnRows(rows)

	result, err := repo.ReportCarriersByLocality(nil)

	require.NoError(t, err)
	require.Equal(t, expectedReports, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestReportCarriersByLocality_QueryError_ReturnsError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	repo := NewLocalityMysqlRepository(db)

	localityId := 1
	expectedError := errors.New("connection with db was broken and failed")
	mock.ExpectQuery("^SELECT (.+) FROM localities l (.+)").
		WithArgs(localityId, localityId).
		WillReturnError(expectedError)

	result, err := repo.ReportCarriersByLocality(&localityId)

	require.Error(t, err)
	require.Equal(t, expectedError, err)
	require.Nil(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestReportCarriersByLocality_ScanError_ReturnsError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	repo := NewLocalityMysqlRepository(db)

	localityId := 1
	rows := sqlmock.NewRows([]string{
		"id", "locality_name", "carriers_count",
	}).
		AddRow("invalid_id", "Buenos Aires", 3) // invalid_id is not an int

	mock.ExpectQuery("^SELECT (.+) FROM localities l (.+)").
		WithArgs(localityId, localityId).
		WillReturnRows(rows)

	result, err := repo.ReportCarriersByLocality(&localityId)

	require.Error(t, err)
	require.Nil(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestReportCarriersByLocality_RowsErrError_ReturnsError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	repo := NewLocalityMysqlRepository(db)

	localityId := 1
	expectedError := errors.New("rows iteration error")
	rows := sqlmock.NewRows([]string{
		"id", "locality_name", "carriers_count",
	}).
		AddRow(1, "Buenos Aires", 3).
		CloseError(expectedError)

	mock.ExpectQuery("^SELECT (.+) FROM localities l (.+)").
		WithArgs(localityId, localityId).
		WillReturnRows(rows)

	result, err := repo.ReportCarriersByLocality(&localityId)

	require.Error(t, err)
	require.Equal(t, expectedError, err)
	require.Nil(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestReportCarriersByLocality_NoDataFound_ReturnsErrNotFound(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	repo := NewLocalityMysqlRepository(db)

	localityId := 999

	rows := sqlmock.NewRows([]string{
		"id", "locality_name", "carriers_count",
	})
	mock.ExpectQuery("^SELECT (.+) FROM localities l (.+)").
		WithArgs(localityId, localityId).
		WillReturnRows(rows)

	result, err := repo.ReportCarriersByLocality(&localityId)

	require.Error(t, err)
	require.ErrorIs(t, err, pkgErrors.ErrNotFound)
	require.Nil(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}
