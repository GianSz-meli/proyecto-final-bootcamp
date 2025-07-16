package employee

const (
	QueryGetAll  = "SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees"
	QueryGetById = "SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees WHERE id = ?"
	QueryCreate  = "INSERT INTO employees (card_number_id, first_name, last_name, warehouse_id) VALUES (?, ?, ?, ?)"
	QueryUpdate  = "UPDATE employees SET card_number_id = ?, first_name = ?, last_name = ?, warehouse_id = ? WHERE id = ?"
	QueryDelete  = "DELETE FROM employees WHERE id = ?"
)
