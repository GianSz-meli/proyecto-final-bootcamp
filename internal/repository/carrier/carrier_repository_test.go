package carrier

import (
	"errors"
	"testing"

	"ProyectoFinal/pkg/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestCreate_Success(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewSqlCarrierRepository(db)

	inputCarrier := &models.Carrier{
		Cid:         "C001",
		CompanyName: "Express Delivery",
		Address:     "123 Main St",
		Telephone:   "1234567890",
		LocalityId:  1,
	}

	expectedCarrier := &models.Carrier{
		Id:          1,
		Cid:         "C001",
		CompanyName: "Express Delivery",
		Address:     "123 Main St",
		Telephone:   "1234567890",
		LocalityId:  1,
	}

	mock.ExpectExec("^INSERT INTO carriers").
		WithArgs("C001", "Express Delivery", "123 Main St", "1234567890", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := repo.Create(inputCarrier)

	require.NoError(t, err)
	require.Equal(t, expectedCarrier, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestCreate_DatabaseError_ReturnsError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewSqlCarrierRepository(db)

	inputCarrier := &models.Carrier{
		Cid:         "C001",
		CompanyName: "Express Delivery",
		Address:     "123 Main St",
		Telephone:   "1234567890",
		LocalityId:  1,
	}

	expectedError := errors.New("connection with db was broken and failed")
	mock.ExpectExec("^INSERT INTO carriers").
		WithArgs("C001", "Express Delivery", "123 Main St", "1234567890", 1).
		WillReturnError(expectedError)

	result, err := repo.Create(inputCarrier)

	require.Error(t, err)
	require.Equal(t, expectedError, err)
	require.Nil(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestCreate_LastInsertIdError_ReturnsError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := NewSqlCarrierRepository(db)

	inputCarrier := &models.Carrier{
		Cid:         "C001",
		CompanyName: "Express Delivery",
		Address:     "123 Main St",
		Telephone:   "1234567890",
		LocalityId:  1,
	}

	expectedError := errors.New("last insert id failed")
	mock.ExpectExec("^INSERT INTO carriers").
		WithArgs("C001", "Express Delivery", "123 Main St", "1234567890", 1).
		WillReturnResult(sqlmock.NewErrorResult(expectedError))

	result, err := repo.Create(inputCarrier)

	require.Error(t, err)
	require.Equal(t, expectedError, err)
	require.Nil(t, result)
	require.NoError(t, mock.ExpectationsWereMet())
}
