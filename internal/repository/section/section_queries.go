package repository

const (
	// SectionSelectAll retrieves all sections from the sections table
	SectionSelectAll = `SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, product_type_id, warehouse_id, maximum_capacity FROM sections`
	// SectionSelectById retrieves a specific section by its ID
	SectionSelectById = `SELECT id, section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, product_type_id, warehouse_id, maximum_capacity FROM sections WHERE id = ?`
	// SectionInsert creates a new section record
	SectionInsert = `INSERT INTO sections (section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, product_type_id, warehouse_id, maximum_capacity) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	// SectionUpdate modifies an existing section record by ID
	SectionUpdate = `UPDATE sections SET section_number=?, current_temperature=?, minimum_temperature=?, current_capacity=?, minimum_capacity=?, product_type_id=?, warehouse_id=?, maximum_capacity=? WHERE id=?`
	// SectionDelete removes a section record by ID
	SectionDelete = `DELETE FROM sections WHERE id=?`
	// SectionExistsByNumber checks if a section exists with the given section number
	SectionExistsByNumber = `SELECT EXISTS(SELECT 1 FROM sections WHERE section_number = ?)`
)
