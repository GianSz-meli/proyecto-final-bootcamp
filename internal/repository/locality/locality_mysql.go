package locality

import (
	pkgErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"database/sql"
	"errors"
)

type LocalityMysql struct {
	db *sql.DB
}

func NewLocalityMysqlRepository(db *sql.DB) LocalityRepository {
	return &LocalityMysql{
		db: db,
	}
}

func (r *LocalityMysql) Create(locality models.Locality) (models.Locality, error) {
	result, err := r.db.Exec(SQL_CREATE, locality.Id, locality.LocalityName, locality.Province.ProvinceName, locality.Province.Country.CountryName)
	if err != nil {
		return models.Locality{}, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return models.Locality{}, err
	}
	locality.Id = int(lastId)
	return locality, nil
}

func (r *LocalityMysql) GetById(id int) (*models.Locality, error) {
	row := r.db.QueryRow(SQL_GET_BY_ID, id)
	if err := row.Err(); err != nil {
		return nil, err
	}
	var locality models.Locality

	if err := LocalityScan(row, &locality); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			newError := pkgErrors.WrapErrNotFound("Locality", "id", id)
			return nil, newError
		}
		return nil, err
	}

	return &locality, nil
}

func (r *LocalityMysql) GetSellersByIdLocality(idLocality int) (models.SellersByLocalityReport, error) {
	row := r.db.QueryRow(SQL_SELLERS_BY_ID_LOCALITY, idLocality)
	if err := row.Err(); err != nil {
		return models.SellersByLocalityReport{}, err
	}
	var sellerByLocality models.SellersByLocalityReport

	if err := SellersByLocalityScan(row, &sellerByLocality); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			newError := pkgErrors.WrapErrNotFound("locality", "id", idLocality)
			return models.SellersByLocalityReport{}, newError
		}
		return models.SellersByLocalityReport{}, err
	}

	return sellerByLocality, nil
}

func (r *LocalityMysql) GetSellersByLocalities() ([]models.SellersByLocalityReport, error) {
	rows, err := r.db.Query(SQL_SELLERS_BY_LOCALITY)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		return nil, err
	}
	var sellersByLocality []models.SellersByLocalityReport

	for rows.Next() {
		var sellerByLocality models.SellersByLocalityReport
		if err = SellersByLocalityScan(rows, &sellerByLocality); err != nil {
			return nil, err
		}
		sellersByLocality = append(sellersByLocality, sellerByLocality)
	}

	return sellersByLocality, nil
}
