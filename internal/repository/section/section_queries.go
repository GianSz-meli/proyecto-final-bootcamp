package repository

const (
	SectionSelectAll      = `SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, product_type_id, warehouse_id, maximum_capacity FROM sections`
	SectionSelectById     = `SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, product_type_id, warehouse_id, maximum_capacity FROM sections WHERE id = ?`
	SectionInsert         = `INSERT INTO sections (section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, product_type_id, warehouse_id, maximum_capacity) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	SectionUpdate         = `UPDATE sections SET section_number=?, current_temperature=?, minimum_temperature=?, current_capacity=?, minimum_capacity=?, product_type_id=?, warehouse_id=?, maximum_capacity=? WHERE id=?`
	SectionDelete         = `DELETE FROM sections WHERE id=?`
	SectionExistsByNumber = `SELECT EXISTS(SELECT 1 FROM sections WHERE section_number = ?)`
)
