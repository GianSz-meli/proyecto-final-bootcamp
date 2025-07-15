package repository

const (
	// ProductBatch queries
	ProductBatchInsert = `INSERT INTO product_batches (batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, product_id, section_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	// Validation queries
	ProductBatchExistsByNumber = `SELECT EXISTS(SELECT 1 FROM product_batches WHERE batch_number = ?)`
	ProductExistsById          = `SELECT EXISTS(SELECT 1 FROM products WHERE id = ?)`
	SectionExistsById          = `SELECT EXISTS(SELECT 1 FROM sections WHERE id = ?)`

	// Report queries
	ProductBatchCountBySection = `SELECT s.id as section_id, s.section_number, COUNT(pb.id) as products_count 
		FROM sections s 
		LEFT JOIN product_batches pb ON s.id = pb.section_id 
		GROUP BY s.id, s.section_number`

	ProductBatchCountBySectionId = `SELECT s.id as section_id, s.section_number, COUNT(pb.id) as products_count 
		FROM sections s 
		LEFT JOIN product_batches pb ON s.id = pb.section_id 
		WHERE s.id = ?
		GROUP BY s.id, s.section_number`
)
