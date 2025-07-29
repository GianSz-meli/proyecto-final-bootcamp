package repository

import (
	"ProyectoFinal/pkg/models"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestMySQLRepository_CreateProductRecord(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewProductRecordSQL(db)

	productRecord := models.ProductRecord{
		ID:             1,
		LastUpdateDate: "2024-01-01",
		PurchasePrice:  100,
		SalePrice:      150,
		ProductID:      1,
	}

	mock.ExpectExec("INSERT INTO products_records").WithArgs(productRecord.LastUpdateDate, productRecord.PurchasePrice, productRecord.SalePrice, productRecord.ProductID).WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := repo.CreateProductRecord(productRecord)
	require.NoError(t, err)
	require.Equal(t, productRecord, result)

	mock.ExpectationsWereMet()
}

func TestMySQLRepository_GetRecordsProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewProductRecordSQL(db)

	product := models.Product{
		ID:          1,
		Description: "Coca Cola 2L",
	}

	mock.ExpectQuery("SELECT p.id, p.description, COUNT\\(pr.id\\) AS records_count FROM products p LEFT JOIN products_records pr ON pr.product_id = p.id WHERE p.id = \\? GROUP BY p.id, p.description").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "description", "records_count"}).AddRow(product.ID, product.Description, 1))

	expected := models.ReportProductData{
		ProductID:    1,
		Description:  "Coca Cola 2L",
		RecordsCount: 1,
	}

	result, err := repo.GetRecordsProduct(1)
	require.NoError(t, err)
	require.Equal(t, expected, result)

	mock.ExpectationsWereMet()
}

func TestMySQLRepository_GetRecordsProduct_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewProductRecordSQL(db)

	mock.ExpectQuery("SELECT p.id, p.description, COUNT\\(pr.id\\) AS records_count FROM products p LEFT JOIN products_records pr ON pr.product_id = p.id WHERE p.id = \\? GROUP BY p.id, p.description").
		WithArgs(999).
		WillReturnError(sql.ErrNoRows)

	result, err := repo.GetRecordsProduct(999)

	require.Error(t, err)
	require.Equal(t, models.ReportProductData{}, result)

	mock.ExpectationsWereMet()
}

func TestMySQLRepository_GetRecordsProduct_DatabaseError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewProductRecordSQL(db)

	expectedError := errors.New("database connection failed")
	mock.ExpectQuery("SELECT p.id, p.description, COUNT\\(pr.id\\) AS records_count FROM products p LEFT JOIN products_records pr ON pr.product_id = p.id WHERE p.id = \\? GROUP BY p.id, p.description").
		WithArgs(1).
		WillReturnError(expectedError)

	result, err := repo.GetRecordsProduct(1)

	require.Equal(t, models.ReportProductData{}, result)
	require.Equal(t, expectedError, err)

	mock.ExpectationsWereMet()
}

func TestMySQLRepository_GetRecordsProductAll(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	require.NoError(t, err)
	defer db.Close()

	repo := NewProductRecordSQL(db)

	mock.ExpectQuery("SELECT p\\.id, p\\.description, COUNT\\(pr\\.id\\) as records_count FROM products p JOIN products_records pr ON pr\\.product_id = p\\.id GROUP BY p\\.id, p\\.description").WillReturnRows(sqlmock.NewRows([]string{"id", "description", "records_count"}).AddRow(1, "Coca Cola 2L", 1).AddRow(2, "Coca Cola 1L", 2))

	expected := []models.ReportProductData{
		{ProductID: 1, Description: "Coca Cola 2L", RecordsCount: 1},
		{ProductID: 2, Description: "Coca Cola 1L", RecordsCount: 2},
	}

	result, err := repo.GetRecordsProductAll()
	require.NoError(t, err)
	require.Equal(t, expected, result)

	mock.ExpectationsWereMet()
}
