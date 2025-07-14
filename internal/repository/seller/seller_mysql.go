package seller

import (
	"ProyectoFinal/internal/repository/utils"
	pkgErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"database/sql"
	"errors"
	"fmt"
)

type SellerMysql struct {
	db *sql.DB
}

func NewSellerRepository(db *sql.DB) SellerRepository {
	return &SellerMysql{
		db: db,
	}
}

func (r *SellerMysql) Create(seller models.Seller) (models.Seller, error) {
	result, err := r.db.Exec(SQL_CREATE, seller.Cid, seller.CompanyName, seller.Address, seller.Telephone, seller.LocalityId)
	if err != nil {
		return models.Seller{}, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return models.Seller{}, err
	}
	seller.Id = int(lastId)
	return seller, nil
}

func (r *SellerMysql) GetById(id int) (*models.Seller, error) {
	row := r.db.QueryRow(SQL_GET_BY_ID, id)
	if err := row.Err(); err != nil {
		return nil, err
	}
	var seller models.Seller

	if err := utils.SellerScan(row, &seller); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			newError := pkgErrors.WrapErrNotFound("Seller", "id", id)
			return nil, newError
		}
		return nil, err
	}

	return &seller, nil
}

func (r *SellerMysql) ExistsByCid(cid string) (bool, error) {
	var exists bool

	if err := r.db.QueryRow(SQL_EXIST_BY_CID, cid).Scan(&exists); err != nil {
		return false, err
	}
	return exists, nil
}

func (r *SellerMysql) GetAll() ([]models.Seller, error) {
	rows, err := r.db.Query(SQL_GET_ALL)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		return nil, err
	}
	var sellers []models.Seller

	for rows.Next() {
		var seller models.Seller
		if err = utils.SellerScan(rows, &seller); err != nil {
			return nil, err
		}
		sellers = append(sellers, seller)
	}

	return sellers, nil
}

func (r *SellerMysql) Delete(id int) error {
	_, err := r.db.Exec(SQL_DELETE, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *SellerMysql) Update(seller *models.Seller) (models.Seller, error) {

	_, err := r.db.Exec(SQL_UPDATE, seller.Cid, seller.CompanyName, seller.Address, seller.Telephone, seller.LocalityId, seller.Id)
	if err != nil {
		return models.Seller{}, err
	}
	return *seller, nil
}
