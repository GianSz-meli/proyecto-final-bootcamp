package repository

const (
	QueryExists = "SELECT EXISTS (SELECT 1 FROM products WHERE id = ?)"
	QueryCreate = "INSERT INTO product_records (last_update_date, purchase_price, sale_price, product_id) VALUES (?, ?, ?, ?)"
	QueryReport = "SELECT p.id, p.description, COUNT(pr.id) AS records_count FROM products p LEFT JOIN product_records pr ON pr.product_id = p.i WHERE p.id = ? GROUP BY p.id, p.description"
)
