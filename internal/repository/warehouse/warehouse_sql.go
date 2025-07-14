package warehouse

import (
	"ProyectoFinal/pkg/models"
	"database/sql"
)

type SqlWarehouseRepository struct {
	db *sql.DB
}

func NewSqlWarehouseRepository(db *sql.DB) *SqlWarehouseRepository {
	return &SqlWarehouseRepository{
		db: db,
	}
}

func (r *SqlWarehouseRepository) GetAll() ([]models.Warehouse, error) {
	query := `SELECT id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature, locality_id 
			  FROM warehouses ORDER BY id`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var warehouses []models.Warehouse
	for rows.Next() {
		var warehouse models.Warehouse
		err := rows.Scan(
			&warehouse.ID,
			&warehouse.WarehouseCode,
			&warehouse.Address,
			&warehouse.Telephone,
			&warehouse.MinimumCapacity,
			&warehouse.MinimumTemperature,
			&warehouse.LocalityId,
		)
		if err != nil {
			return nil, err
		}
		warehouses = append(warehouses, warehouse)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return warehouses, nil
}

func (r *SqlWarehouseRepository) GetById(id int) (*models.Warehouse, error) {
	query := `SELECT id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature, locality_id 
			  FROM warehouses WHERE id = ?`

	var warehouse models.Warehouse
	err := r.db.QueryRow(query, id).Scan(
		&warehouse.ID,
		&warehouse.WarehouseCode,
		&warehouse.Address,
		&warehouse.Telephone,
		&warehouse.MinimumCapacity,
		&warehouse.MinimumTemperature,
		&warehouse.LocalityId,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &warehouse, nil
}

func (r *SqlWarehouseRepository) ExistsByCode(code string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM warehouses WHERE warehouse_code = ?)`

	var exists bool
	err := r.db.QueryRow(query, code).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
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
	if err != nil {
		return err
	}
	return nil
}
