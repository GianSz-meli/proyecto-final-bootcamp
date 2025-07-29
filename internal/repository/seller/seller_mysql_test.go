package seller

import (
	pkgErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSellerMysql_Create_QueryError(t *testing.T) {
	//Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := NewSellerMysqlRepository(db)
	seller := models.Seller{
		Cid:         "GDJ2SJ3",
		CompanyName: "Farm to Table Produce Hub",
		Address:     "812 Cypress Way, Denver, CO 80201",
		Telephone:   "+1-555-1901",
		LocalityId:  1,
	}
	expectedError := errors.New("error sql query")
	mock.ExpectExec("^INSERT INTO sellers").
		WithArgs(
			seller.Cid,
			seller.CompanyName,
			seller.Address,
			seller.Telephone,
			seller.LocalityId,
		).
		WillReturnError(expectedError)
	//Act
	result, err := repo.Create(seller)
	//Assert
	require.Equal(t, models.Seller{}, result)
	require.Error(t, err)
	require.EqualError(t, expectedError, err.Error())
	require.NoError(t, mock.ExpectationsWereMet())
}
func TestSellerMysql_Create_GetLastInsertError(t *testing.T) {
	//Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := NewSellerMysqlRepository(db)
	seller := models.Seller{
		Cid:         "GDJ2SJ3",
		CompanyName: "Farm to Table Produce Hub",
		Address:     "812 Cypress Way, Denver, CO 80201",
		Telephone:   "+1-555-1901",
		LocalityId:  1,
	}
	expectedError := errors.New("error")
	mock.ExpectExec("^INSERT INTO sellers").
		WithArgs(
			seller.Cid,
			seller.CompanyName,
			seller.Address,
			seller.Telephone,
			seller.LocalityId,
		).WillReturnResult(sqlmock.NewErrorResult(expectedError))
	//Act
	result, err := repo.Create(seller)
	//Assert
	require.Equal(t, models.Seller{}, result)
	require.Error(t, err)
	require.EqualError(t, expectedError, err.Error())
	require.NoError(t, mock.ExpectationsWereMet())
}
func TestSellerMysql_Create_Success(t *testing.T) {
	//Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := NewSellerMysqlRepository(db)
	seller := models.Seller{
		Cid:         "GDJ2SJ3",
		CompanyName: "Farm to Table Produce Hub",
		Address:     "812 Cypress Way, Denver, CO 80201",
		Telephone:   "+1-555-1901",
		LocalityId:  1,
	}
	newSellerId := 1
	newSeller := seller
	newSeller.Id = newSellerId
	mock.ExpectExec("^INSERT INTO sellers").
		WithArgs(
			seller.Cid,
			seller.CompanyName,
			seller.Address,
			seller.Telephone,
			seller.LocalityId,
		).WillReturnResult(sqlmock.NewResult(int64(newSellerId), 1))
	//Act
	result, err := repo.Create(seller)
	//Assert
	require.Equal(t, newSeller, result)
	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSellerMysql_GetById_QueryError(t *testing.T) {
	//Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := NewSellerMysqlRepository(db)
	sellerId := 1
	expectedError := errors.New("error sql query")

	mock.ExpectQuery("^SELECT (.+) FROM sellers WHERE id = \\?").
		WithArgs(sellerId).
		WillReturnError(expectedError)
	//Act
	result, err := repo.GetById(sellerId)
	//Assert
	require.Nil(t, result)
	require.Error(t, err)
	require.EqualError(t, expectedError, err.Error())
	require.NoError(t, mock.ExpectationsWereMet())
}
func TestSellerMysql_GetById_NotFoundError(t *testing.T) {
	//Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := NewSellerMysqlRepository(db)
	sellerId := 1
	rows := sqlmock.NewRows([]string{
		"id", "cid", "company_name", "address", "telephone", "locality_id",
	})
	mock.ExpectQuery("^SELECT (.+) FROM sellers WHERE id = \\?").
		WithArgs(sellerId).WillReturnRows(rows)
	//Act
	result, err := repo.GetById(sellerId)
	//Assert
	require.Nil(t, result)
	require.Error(t, err)
	require.ErrorIs(t, err, pkgErrors.ErrNotFound)
	require.NoError(t, mock.ExpectationsWereMet())
}
func TestSellerMysql_GetById_ScannerError(t *testing.T) {
	//Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := NewSellerMysqlRepository(db)
	sellerId := 1
	rows := sqlmock.NewRows([]string{
		"id", "cid", "company_name", "address", "telephone", "locality_id",
	}).AddRow("NOT INT", "GDJ2SJ3", "Farm to Table Produce Hub", "812 Cypress Way, Denver, CO 80201", "+1-555-1901", 1)

	mock.ExpectQuery("^SELECT (.+) FROM sellers WHERE id = \\?").
		WithArgs(sellerId).WillReturnRows(rows)
	//Act
	result, err := repo.GetById(sellerId)
	//Assert
	require.Nil(t, result)
	require.Error(t, err)
	require.Contains(t, err.Error(), "Scan error")
	require.NoError(t, mock.ExpectationsWereMet())
}
func TestSellerMysql_GetById_Success(t *testing.T) {
	//Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := NewSellerMysqlRepository(db)
	sellerId := 1
	expectedSeller := models.Seller{
		Id:          1,
		Cid:         "GDJ2SJ3",
		CompanyName: "Farm to Table Produce Hub",
		Address:     "812 Cypress Way, Denver, CO 80201",
		Telephone:   "+1-555-1901",
		LocalityId:  1,
	}
	rows := sqlmock.NewRows([]string{
		"id", "cid", "company_name", "address", "telephone", "locality_id",
	}).AddRow(expectedSeller.Id, expectedSeller.Cid, expectedSeller.CompanyName, expectedSeller.Address, expectedSeller.Telephone, expectedSeller.LocalityId)

	mock.ExpectQuery("^SELECT (.+) FROM sellers WHERE id = \\?").
		WithArgs(sellerId).
		WillReturnRows(rows)
	//Act
	result, err := repo.GetById(expectedSeller.Id)
	//Assert
	require.Equal(t, &expectedSeller, result)
	require.Nil(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSellerMysql_GetAll_QueryError(t *testing.T) {
	//Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := NewSellerMysqlRepository(db)
	expectedError := errors.New("error sql query")

	mock.ExpectQuery("^SELECT (.+) FROM sellers").
		WillReturnError(expectedError)
	//Act
	result, err := repo.GetAll()
	//Assert
	require.Nil(t, result)
	require.Error(t, err)
	require.EqualError(t, expectedError, err.Error())
	require.NoError(t, mock.ExpectationsWereMet())
}
func TestSellerMysql_GetAll_RowsError(t *testing.T) {
	//Arrange
	expectedSellers := []models.Seller{
		{
			Id:          1,
			Cid:         "GDJ2SJ3",
			CompanyName: "Farm to Table Produce Hub",
			Address:     "812 Cypress Way, Denver, CO 80201",
			Telephone:   "+1-555-1901",
			LocalityId:  1,
		},
		{
			Id:          2,
			Cid:         "GDJ2SJ3",
			CompanyName: "Farm to Table Produce Hub",
			Address:     "812 Cypress Way, Denver, CO 80201",
			Telephone:   "+1-555-1901",
			LocalityId:  2,
		},
	}
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := NewSellerMysqlRepository(db)
	expectedError := errors.New("error scanning rows")
	rows := sqlmock.NewRows([]string{
		"id", "cid", "company_name", "address", "telephone", "locality_id",
	}).
		AddRow(expectedSellers[0].Id, expectedSellers[0].Cid, expectedSellers[0].CompanyName, expectedSellers[0].Address, expectedSellers[0].Telephone, expectedSellers[0].LocalityId).
		AddRow(expectedSellers[1].Id, expectedSellers[1].Cid, expectedSellers[1].CompanyName, expectedSellers[1].Address, expectedSellers[1].Telephone, expectedSellers[1].LocalityId).
		RowError(0, expectedError)

	mock.ExpectQuery("^SELECT (.+) FROM sellers").
		WillReturnRows(rows)
	//Act
	result, err := repo.GetAll()
	//Assert

	require.Nil(t, result)
	require.Error(t, err)
	require.EqualError(t, expectedError, err.Error())
	require.NoError(t, mock.ExpectationsWereMet())
}
func TestSellerMysql_GetAll_ScannerError(t *testing.T) {
	//Arrange
	expectedSellers := []models.Seller{
		{
			Id:          1,
			Cid:         "GDJ2SJ3",
			CompanyName: "Farm to Table Produce Hub",
			Address:     "812 Cypress Way, Denver, CO 80201",
			Telephone:   "+1-555-1901",
			LocalityId:  1,
		},
		{
			Id:          2,
			Cid:         "GDJ2SJ3",
			CompanyName: "Farm to Table Produce Hub",
			Address:     "812 Cypress Way, Denver, CO 80201",
			Telephone:   "+1-555-1901",
			LocalityId:  2,
		},
	}
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := NewSellerMysqlRepository(db)

	rows := sqlmock.NewRows([]string{
		"id", "cid", "company_name", "address", "telephone", "locality_id",
	}).
		AddRow(expectedSellers[0].Id, expectedSellers[0].Cid, expectedSellers[0].CompanyName, expectedSellers[0].Address, expectedSellers[0].Telephone, expectedSellers[0].LocalityId).
		AddRow("NOT INT", expectedSellers[1].Cid, expectedSellers[1].CompanyName, expectedSellers[1].Address, expectedSellers[1].Telephone, expectedSellers[1].LocalityId)

	mock.ExpectQuery("^SELECT (.+) FROM sellers").
		WillReturnRows(rows)
	//Act
	result, err := repo.GetAll()
	//Assert

	require.Nil(t, result)
	require.Error(t, err)
	require.Contains(t, err.Error(), "Scan error")
	require.NoError(t, mock.ExpectationsWereMet())
}
func TestSellerMysql_GetAll_Success(t *testing.T) {
	//Arrange
	expectedSellers := []models.Seller{
		{
			Id:          1,
			Cid:         "GDJ2SJ3",
			CompanyName: "Farm to Table Produce Hub",
			Address:     "812 Cypress Way, Denver, CO 80201",
			Telephone:   "+1-555-1901",
			LocalityId:  1,
		},
		{
			Id:          2,
			Cid:         "GDJ2SJ3",
			CompanyName: "Farm to Table Produce Hub",
			Address:     "812 Cypress Way, Denver, CO 80201",
			Telephone:   "+1-555-1901",
			LocalityId:  2,
		},
	}
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := NewSellerMysqlRepository(db)

	rows := sqlmock.NewRows([]string{
		"id", "cid", "company_name", "address", "telephone", "locality_id",
	}).
		AddRow(expectedSellers[0].Id, expectedSellers[0].Cid, expectedSellers[0].CompanyName, expectedSellers[0].Address, expectedSellers[0].Telephone, expectedSellers[0].LocalityId).
		AddRow(expectedSellers[1].Id, expectedSellers[1].Cid, expectedSellers[1].CompanyName, expectedSellers[1].Address, expectedSellers[1].Telephone, expectedSellers[1].LocalityId)

	mock.ExpectQuery("^SELECT (.+) FROM sellers").
		WillReturnRows(rows)
	//Act
	result, err := repo.GetAll()
	//Assert

	require.Nil(t, err)
	require.Equal(t, expectedSellers, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSellerMysql_Delete_QueryError(t *testing.T) {
	//Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	sellerId := 1
	repo := NewSellerMysqlRepository(db)
	expectedError := errors.New("error sql query")
	mock.ExpectExec("^DELETE FROM sellers").
		WithArgs(sellerId).
		WillReturnError(expectedError)
	//Act
	err = repo.Delete(sellerId)
	//Assert
	require.Error(t, err)
	require.EqualError(t, expectedError, err.Error())
	require.NoError(t, mock.ExpectationsWereMet())
}
func TestSellerMysql_Delete_RowsAffectedError(t *testing.T) {
	//Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	sellerId := 1
	repo := NewSellerMysqlRepository(db)
	expectedError := errors.New("rows affected error")
	mock.ExpectExec("^DELETE FROM sellers").
		WithArgs(sellerId).
		WillReturnResult(sqlmock.NewErrorResult(expectedError))
	//Act
	err = repo.Delete(sellerId)
	//Assert
	require.Error(t, err)
	require.EqualError(t, expectedError, err.Error())
	require.NoError(t, mock.ExpectationsWereMet())
}
func TestSellerMysql_Delete_NotFoundError(t *testing.T) {
	//Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	sellerId := 1
	repo := NewSellerMysqlRepository(db)
	expectedError := pkgErrors.WrapErrNotFound("Seller", "id", sellerId)
	mock.ExpectExec("^DELETE FROM sellers").
		WithArgs(sellerId).
		WillReturnResult(sqlmock.NewResult(0, 0))
	//Act
	err = repo.Delete(sellerId)
	//Assert
	require.Error(t, err)
	require.EqualError(t, expectedError, err.Error())
	require.ErrorIs(t, err, pkgErrors.ErrNotFound)
	require.NoError(t, mock.ExpectationsWereMet())
}
func TestSellerMysql_Delete_Success(t *testing.T) {
	//Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	sellerId := 1
	repo := NewSellerMysqlRepository(db)
	mock.ExpectExec("^DELETE FROM sellers").
		WithArgs(sellerId).
		WillReturnResult(sqlmock.NewResult(int64(sellerId), 1))
	//Act
	err = repo.Delete(sellerId)
	//Assert
	require.Nil(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestSellerMysql_Update_QueryError(t *testing.T) {
	//Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := NewSellerMysqlRepository(db)
	seller := models.Seller{
		Id:          1,
		Cid:         "GDJ2SJ3",
		CompanyName: "Farm to Table Produce Hub",
		Address:     "812 Cypress Way, Denver, CO 80201",
		Telephone:   "+1-555-1901",
		LocalityId:  1,
	}
	expectedError := errors.New("error sql query")
	mock.ExpectExec("^UPDATE sellers SET").
		WithArgs(
			seller.Cid,
			seller.CompanyName,
			seller.Address,
			seller.Telephone,
			seller.LocalityId,
			seller.Id,
		).
		WillReturnError(expectedError)
	//Act
	result, err := repo.Update(&seller)
	//Assert
	require.Equal(t, models.Seller{}, result)
	require.Error(t, err)
	require.EqualError(t, expectedError, err.Error())
	require.NoError(t, mock.ExpectationsWereMet())
}
func TestSellerMysql_Update_Success(t *testing.T) {
	//Arrange
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()
	repo := NewSellerMysqlRepository(db)
	seller := models.Seller{
		Id:          1,
		Cid:         "GDJ2SJ3",
		CompanyName: "Farm to Table Produce Hub",
		Address:     "812 Cypress Way, Denver, CO 80201",
		Telephone:   "+1-555-1901",
		LocalityId:  1,
	}
	mock.ExpectExec("^UPDATE sellers SET").
		WithArgs(
			seller.Cid,
			seller.CompanyName,
			seller.Address,
			seller.Telephone,
			seller.LocalityId,
			seller.Id,
		).WillReturnResult(sqlmock.NewResult(1, 1))
	//Act
	result, err := repo.Update(&seller)
	//Assert
	require.Equal(t, seller, result)
	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}
