package employee

const (
	// EmployeeQueries contains all SQL queries for employee operations
	QueryGetAllEmployees = "SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees"

	QueryGetEmployeeById = "SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees WHERE id = ?"

	QueryCreateEmployee = "INSERT INTO employees (card_number_id, first_name, last_name, warehouse_id) VALUES (?, ?, ?, ?)"

	QueryExistsByCardNumberId = "SELECT EXISTS(SELECT 1 FROM employees WHERE card_number_id = ?)"

	QueryUpdateEmployee = "UPDATE employees SET card_number_id = ?, first_name = ?, last_name = ?, warehouse_id = ? WHERE id = ?"

	QueryDeleteEmployee = "DELETE FROM employees WHERE id = ?"
)
