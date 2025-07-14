package warehouse

import (
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
	query := `
		SELECT 
			w.id, w.warehouse_code, w.address, w.telephone, w.minimum_capacity, w.minimum_temperature, w.locality_id,
			l.id, l.locality_name,
			p.id, p.province_name,
			c.id, c.country_name
		FROM warehouses w
		LEFT JOIN localities l ON w.locality_id = l.id
		LEFT JOIN provinces p ON l.province_id = p.id
		LEFT JOIN countries c ON p.country_id = c.id
		ORDER BY w.id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var warehouses []models.Warehouse
	for rows.Next() {
		var warehouse models.Warehouse
		var locality models.Locality
		var province models.Province
		var country models.Country

		err := rows.Scan(
			&warehouse.ID,
			&warehouse.WarehouseCode,
			&warehouse.Address,
			&warehouse.Telephone,
			&warehouse.MinimumCapacity,
			&warehouse.MinimumTemperature,
			&warehouse.LocalityId,
			&locality.Id,
			&locality.LocalityName,
			&province.Id,
			&province.ProvinceName,
			&country.Id,
			&country.CountryName,
		)
		if err != nil {
			return nil, err
		}

		if locality.Id != nil {
			province.Country = &country
			locality.Province = &province
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
	query := `
		SELECT 
			w.id, w.warehouse_code, w.address, w.telephone, w.minimum_capacity, w.minimum_temperature, w.locality_id,
			l.id, l.locality_name,
			p.id, p.province_name,
			c.id, c.country_name
		FROM warehouses w
		LEFT JOIN localities l ON w.locality_id = l.id
		LEFT JOIN provinces p ON l.province_id = p.id
		LEFT JOIN countries c ON p.country_id = c.id
		WHERE w.id = ?
	`

	var warehouse models.Warehouse
	var locality models.Locality
	var province models.Province
	var country models.Country

	err := r.db.QueryRow(query, id).Scan(
		&warehouse.ID,
		&warehouse.WarehouseCode,
		&warehouse.Address,
		&warehouse.Telephone,
		&warehouse.MinimumCapacity,
		&warehouse.MinimumTemperature,
		&warehouse.LocalityId,
		&locality.Id,
		&locality.LocalityName,
		&province.Id,
		&province.ProvinceName,
		&country.Id,
		&country.CountryName,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if locality.Id != nil {
		province.Country = &country
		locality.Province = &province
		warehouse.Locality = &locality
	}

	return &warehouse, nil
}

func (r *SqlWarehouseRepository) ExistsByCode(code string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM warehouses WHERE warehouse_code = ?)`
	var exists bool
	err := r.db.QueryRow(query, code).Scan(&exists)
	return exists, err
}

func (r *SqlWarehouseRepository) Create(warehouse models.Warehouse) (models.Warehouse, error) {
	query := `INSERT INTO warehouses (warehouse_code, address, telephone, minimum_capacity, minimum_temperature, locality_id) 
			  VALUES (?, ?, ?, ?, ?, ?)`

	result, err := r.db.Exec(query,
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
	query := `UPDATE warehouses 
			  SET warehouse_code = ?, address = ?, telephone = ?, minimum_capacity = ?, minimum_temperature = ?, locality_id = ? 
			  WHERE id = ?`

	_, err := r.db.Exec(query,
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

	warehouse.ID = id
	return warehouse, nil
}

func (r *SqlWarehouseRepository) Delete(id int) error {
	query := `DELETE FROM warehouses WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}
