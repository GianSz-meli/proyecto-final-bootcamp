package products

import (
	"ProyectoFinal/pkg/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func ptrInt(v int) *int { return &v }

func TestMySQLRepository_CreateProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewProductSQL(db)

	product := models.Product{
		ID:             1,
		ProductCode:    "A1000",
		Description:    "Coca Cola 2L",
		Width:          5,
		Height:         20,
		Length:         12,
		NetWeight:      1.9,
		ExpirationRate: 15,
		Temperature:    10,
		FreezingRate:   10,
		ProductTypeID:  83,
		SellerID:       ptrInt(7),
	}

	mock.ExpectExec("INSERT INTO products").WithArgs(product.ProductCode, product.Description, product.Width, product.Height, product.Length, product.NetWeight, product.ExpirationRate, product.Temperature, product.FreezingRate, product.ProductTypeID, product.SellerID).WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := repo.CreateProduct(product)
	require.NoError(t, err)
	require.Equal(t, product, result)

	mock.ExpectationsWereMet()
}

func TestMySQLRepository_FindAllProducts(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewProductSQL(db)

	products := make(map[int]models.Product)
	products[1] = models.Product{
		ID:             1,
		ProductCode:    "A1000",
		Description:    "Coca Cola 2L",
		Width:          5,
		Height:         20,
		Length:         12,
		NetWeight:      1.9,
		ExpirationRate: 15,
		Temperature:    10,
		FreezingRate:   10,
		ProductTypeID:  83,
		SellerID:       ptrInt(7),
	}

	rows := sqlmock.NewRows([]string{"id", "product_code", "description", "width", "height", "length", "net_weight", "expiration_rate", "recommended_freezing_temperature", "freezing_rate", "product_type_id", "seller_id"}).
		AddRow(products[1].ID, products[1].ProductCode, products[1].Description, products[1].Width, products[1].Height, products[1].Length, products[1].NetWeight, products[1].ExpirationRate, products[1].Temperature, products[1].FreezingRate, products[1].ProductTypeID, products[1].SellerID)

	mock.ExpectQuery("SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id FROM products").WillReturnRows(rows)

	result, err := repo.FindAllProducts()
	require.NoError(t, err)
	require.Equal(t, products, result)

	mock.ExpectationsWereMet()
}

func TestMySQLRepository_FindProductsById(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewProductSQL(db)

	product := models.Product{
		ID:             1,
		ProductCode:    "A1000",
		Description:    "Coca Cola 2L",
		Width:          5,
		Height:         20,
		Length:         12,
		NetWeight:      1.9,
		ExpirationRate: 15,
		Temperature:    10,
		FreezingRate:   10,
		ProductTypeID:  83,
		SellerID:       ptrInt(7),
	}

	rows := sqlmock.NewRows([]string{"id", "product_code", "description", "width", "height", "length", "net_weight", "expiration_rate", "recommended_freezing_temperature", "freezing_rate", "product_type_id", "seller_id"}).
		AddRow(product.ID, product.ProductCode, product.Description, product.Width, product.Height, product.Length, product.NetWeight, product.ExpirationRate, product.Temperature, product.FreezingRate, product.ProductTypeID, product.SellerID)

	mock.ExpectQuery("SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id FROM products WHERE id = ?").WithArgs(1).WillReturnRows(rows)

	result, err := repo.FindProductsById(1)
	require.NoError(t, err)
	require.Equal(t, product, result)

	mock.ExpectationsWereMet()
}

func TestMySQLRepository_UpdateProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewProductSQL(db)

	product := models.Product{
		ID:             1,
		ProductCode:    "A1000",
		Description:    "Coca Cola 2L",
		Width:          5,
		Height:         20,
		Length:         12,
		NetWeight:      1.9,
		ExpirationRate: 15,
		Temperature:    10,
		FreezingRate:   10,
		ProductTypeID:  83,
		SellerID:       ptrInt(7),
	}

	mock.ExpectExec("UPDATE products").WithArgs(product.ProductCode, product.Description, product.Width, product.Height, product.Length, product.NetWeight, product.ExpirationRate, product.Temperature, product.FreezingRate, product.ProductTypeID, product.SellerID, product.ID).WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := repo.UpdateProduct(1, product)
	require.NoError(t, err)
	require.Equal(t, product, result)

	mock.ExpectationsWereMet()
}

func TestMySQLRepository_DeleteProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := NewProductSQL(db)
	idProduct := 1
	mock.ExpectExec("DELETE FROM products WHERE id = ?").WithArgs(idProduct).WillReturnResult(sqlmock.NewResult(1, 1))

	repo.DeleteProduct(idProduct)

	mock.ExpectationsWereMet()
}
