package repository

const (
	QueryExists = "SELECT EXISTS (SELECT 1 FROM products WHERE id = ?)"
	QueryCreate = "INSERT INTO products_records (last_update_date, purchase_price, sale_price, product_id) VALUES (?, ?, ?, ?)"
	QueryReport = "SELECT p.id, p.description, COUNT(pr.id) AS records_count FROM products p LEFT JOIN products_records pr ON pr.product_id = p.id WHERE p.id = ? GROUP BY p.id, p.description"
	QueryGetAll = "SELECT p.id, p.description, COUNT(pr.id) as records_count FROM products p JOIN products_records pr ON pr.product_id = p.id GROUP BY p.id, p.description "
)
