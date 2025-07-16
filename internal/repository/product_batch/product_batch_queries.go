package repository

const (
	// ProductBatch queries
	// ProductBatchInsert creates a new product batch record with all required fields
	ProductBatchInsert = `INSERT INTO product_batches (batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	// Validation queries
	// ProductBatchExistsByNumber checks if a product batch exists with the given batch number
	ProductBatchExistsByNumber = `SELECT EXISTS(SELECT 1 FROM product_batches WHERE batch_number = ?)`
	// ProductExistsById checks if a product exists with the given product ID
	ProductExistsById = `SELECT EXISTS(SELECT 1 FROM products WHERE id = ?)`
	// SectionExistsById checks if a section exists with the given section ID
	SectionExistsById = `SELECT EXISTS(SELECT 1 FROM sections WHERE id = ?)`

	// Report queries
	// ProductBatchCountBySection retrieves a report of product counts for all sections
	// Returns section ID, section number, and count of product batches in each section
	ProductBatchCountBySection = `SELECT s.id as section_id, s.section_number, COUNT(pb.id) as products_count 
		FROM sections s 
		LEFT JOIN product_batches pb ON s.id = pb.section_id 
		GROUP BY s.id, s.section_number`

	// ProductBatchCountBySectionId retrieves a report of product counts for a specific section
	// Returns section ID, section number, and count of product batches in the specified section
	ProductBatchCountBySectionId = `SELECT s.id as section_id, s.section_number, COUNT(pb.id) as products_count 
		FROM sections s 
		LEFT JOIN product_batches pb ON s.id = pb.section_id 
		WHERE s.id = ?
		GROUP BY s.id, s.section_number`
)
