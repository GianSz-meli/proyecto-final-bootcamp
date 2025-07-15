package warehouse

const (
	GetAllWarehousesQuery = `
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
	GetWarehouseByIdQuery = `
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
	ExistsByCodeQuery = `SELECT EXISTS(SELECT 1 FROM warehouses WHERE warehouse_code = ?)`
	CreateWarehouseQuery = `
		INSERT INTO warehouses (warehouse_code, address, telephone, minimum_capacity, minimum_temperature, locality_id) 
		VALUES (?, ?, ?, ?, ?, ?)
	`
	UpdateWarehouseQuery = `
		UPDATE warehouses 
		SET warehouse_code = ?, address = ?, telephone = ?, minimum_capacity = ?, minimum_temperature = ?, locality_id = ?
		WHERE id = ?
	`
	DeleteWarehouseByIdQuery = `DELETE FROM warehouses WHERE id = ?`
)
