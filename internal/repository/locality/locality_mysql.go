package locality

import (
	"ProyectoFinal/internal/repository/utils"
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

	if err := utils.LocalityScan(row, &locality); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			newError := pkgErrors.WrapErrNotFound("Locality", "id", id)
			return nil, newError
		}
		return nil, err
	}

	return &locality, nil
}
