package buyer

import (
	pkgErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"database/sql"
	"errors"
)

type buyerMySql struct {
	db *sql.DB
}

func NewBuyerMySqlRepository(newDB *sql.DB) Repository {
	return &buyerMySql{
		db: newDB,
	}
}

func (r *buyerMySql) GetById(id int) (*models.Buyer, error) {
	var buyer models.Buyer

	err := r.db.QueryRow(GET_BUYER, id).Scan(
		&buyer.Id,
		&buyer.CardNumberId,
		&buyer.FirstName,
		&buyer.LastName,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, pkgErrors.WrapErrNotFound("buyer", "id", id)
	}

	if err != nil {
		return nil, err
	}

	return &buyer, nil
}

func (r *buyerMySql) GetAll() ([]*models.Buyer, error) {
	rows, err := r.db.Query(GET_ALL_BUYERS)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var buyers []*models.Buyer

	for rows.Next() {
		var buyer models.Buyer
		scanErr := rows.Scan(&buyer.Id, &buyer.CardNumberId, &buyer.FirstName, &buyer.LastName)
		if scanErr != nil {
			return nil, scanErr
		}
		buyers = append(buyers, &buyer)
	}

	if rowsErr := rows.Err(); rowsErr != nil {
		return nil, rowsErr
	}

	return buyers, nil
}

func (r *buyerMySql) Create(buyer *models.Buyer) (*models.Buyer, error) {
	result, execErr := r.db.Exec(CREATE_BUYER, buyer.CardNumberId, buyer.FirstName, buyer.LastName)
	if execErr != nil {
		return nil, execErr
	}

	id, idErr := result.LastInsertId()
	if idErr != nil {
		return nil, idErr
	}

	buyer.Id = int(id)
	return buyer, nil
}

func (r *buyerMySql) Update(buyer *models.Buyer) (*models.Buyer, error) {
	result, execErr := r.db.Exec(UPDATE_BUYER, buyer.CardNumberId, buyer.FirstName, buyer.LastName, buyer.Id)
	if execErr != nil {
		return nil, execErr
	}

	rowsAffected, rowsErr := result.RowsAffected()
	if rowsErr != nil {
		return nil, rowsErr
	}

	if rowsAffected == 0 {
		return nil, nil
	}

	return buyer, nil
}

func (r *buyerMySql) Delete(id int) error {
	result, execErr := r.db.Exec(DELETE_BUYER, id)
	if execErr != nil {
		return execErr
	}

	rowsAffected, rowsErr := result.RowsAffected()
	if rowsErr != nil {
		return rowsErr
	}

	if rowsAffected == 0 {
		return nil
	}

	return nil
}

func (r *buyerMySql) ExistsByCardNumberId(id string) (bool, error) {
	var exists int
	err := r.db.QueryRow(EXISTS_BY_CARD_NUMBER, id).Scan(&exists)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
