package repository

import (
	"ProyectoFinal/pkg/models"
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
		ID: 1,
		LastUpdateDate: "2024-01-01",
		PurchasePrice: 100,
		SalePrice: 150,
		ProductID: 1,
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

	productRecord := models.ProductRecord{
		ID: 1,
		LastUpdateDate: "2024-01-01",
		PurchasePrice: 100,
		SalePrice: 150,
		ProductID: 1,
	}

	mock.ExpectQuery("SELECT p.id, p.description, COUNT(pr.id) AS records_count FROM products p LEFT JOIN products_records pr ON pr.product_id = p.id WHERE p.id = ?").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "description", "records_count"}).AddRow(productRecord.ID, productRecord.Description, productRecord.RecordsCount))

	result, err := repo.GetRecordsProduct(1)
	require.NoError(t, err)
	require.Equal(t, productRecord, result)

	mock.ExpectationsWereMet()
}
