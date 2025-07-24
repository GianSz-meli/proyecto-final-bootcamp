package warehouse

import (
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"database/sql"
)

type SqlWarehouseRepository struct {
	db *sql.DB
}

func NewSqlWarehouseRepository(db *sql.DB) WarehouseRepository {
	return &SqlWarehouseRepository{
		db: db,
	}
}

func (r *SqlWarehouseRepository) GetAll() ([]models.Warehouse, error) {
	rows, err := r.db.Query(GetAllWarehousesQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var warehouses []models.Warehouse
	for rows.Next() {
		var warehouse models.Warehouse

		var ( 
			localityId *int
			localityName *string
			provinceId *int
			provinceName *string
			countryId *int
			countryName *string
		)

		err := rows.Scan(
			&warehouse.ID,
			&warehouse.WarehouseCode,
			&warehouse.Address,
			&warehouse.Telephone,
			&warehouse.MinimumCapacity,
			&warehouse.MinimumTemperature,
			&warehouse.LocalityId,
			&localityId,
			&localityName,
			&provinceId,
			&provinceName,
			&countryId,
			&countryName,
		)
		if err != nil {
			return nil, err
		}

		if localityId != nil {
			country := models.Country{
				Id: *countryId,
				CountryName: *countryName,
			}

			province := models.Province{
				Id: *countryId,
				ProvinceName: *provinceName,
				Country: country,
			}
			locality := models.Locality{
				Id: *provinceId,
				LocalityName: *localityName,
				Province: province,
			}
			warehouse.Locality = &locality
		}

		warehouses = append(warehouses, warehouse)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return warehouses, nil
}

func (r *SqlWarehouseRepository) GetById(id int) (*models.Warehouse, error) {
	var warehouse models.Warehouse
	var ( 
		localityId *int
		localityName *string
		provinceId *int
		provinceName *string
		countryId *int
		countryName *string
	)

	err := r.db.QueryRow(GetWarehouseByIdQuery, id).Scan(
		&warehouse.ID,
		&warehouse.WarehouseCode,
		&warehouse.Address,
		&warehouse.Telephone,
		&warehouse.MinimumCapacity,
		&warehouse.MinimumTemperature,
		&warehouse.LocalityId,
		&localityId,
		&localityName,
		&provinceId,
		&provinceName,
		&countryId,
		&countryName,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.WrapErrNotFound("warehouse", "id", id)
		}
		return nil, err
	}

	if localityId != nil {
		country := models.Country{
			Id: *countryId,
			CountryName: *countryName,
		}
		province := models.Province{
			Id: *provinceId,
			ProvinceName: *provinceName,
			Country: country,
		}
		locality := models.Locality{
			Id: *localityId,
			LocalityName: *localityName,
			Province: province,
		}
		warehouse.Locality = &locality
	}

	return &warehouse, nil
}

func (r *SqlWarehouseRepository) Create(warehouse models.Warehouse) (models.Warehouse, error) {
	result, err := r.db.Exec(CreateWarehouseQuery,
		warehouse.WarehouseCode,
		warehouse.Address,
		warehouse.Telephone,
		warehouse.MinimumCapacity,
		warehouse.MinimumTemperature,
		warehouse.LocalityId,
	)

	if err != nil {
		return models.Warehouse{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Warehouse{}, err
	}

	warehouse.ID = int(id)
	return warehouse, nil
}

func (r *SqlWarehouseRepository) Update(id int, warehouse models.Warehouse) (models.Warehouse, error) {
	_, err := r.db.Exec(UpdateWarehouseQuery,
		warehouse.WarehouseCode,
		warehouse.Address,
		warehouse.Telephone,
		warehouse.MinimumCapacity,
		warehouse.MinimumTemperature,
		warehouse.LocalityId,
		id,
	)
	if err != nil {
		return models.Warehouse{}, err
	}

	return warehouse, nil
}

func (r *SqlWarehouseRepository) Delete(id int) error {
	result, err := r.db.Exec(DeleteWarehouseByIdQuery, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.WrapErrNotFound("warehouse", "id", id)
	}

	return nil
}
