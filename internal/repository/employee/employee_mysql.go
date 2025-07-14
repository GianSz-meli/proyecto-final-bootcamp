package employee

import (
	"ProyectoFinal/pkg/models"
	"database/sql"
	"fmt"
)

type mysqlRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) Repository {
	return &mysqlRepository{
		db: db,
	}
}

func (r *mysqlRepository) GetAll() ([]models.Employee, error) {
	rows, err := r.db.Query(QueryGetAllEmployees)
	if err != nil {
		return nil, fmt.Errorf("error querying employees: %w", err)
	}
	defer rows.Close()

	var employees []models.Employee
	for rows.Next() {
		var employee models.Employee
		var warehouseID sql.NullInt64

		err := rows.Scan(&employee.ID, &employee.CardNumberID, &employee.FirstName, &employee.LastName, &warehouseID)
		if err != nil {
			return nil, fmt.Errorf("error scanning employee row: %w", err)
		}

		if warehouseID.Valid {
			employee.WarehouseID = int(warehouseID.Int64)
		}

		employees = append(employees, employee)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating employee rows: %w", err)
	}

	return employees, nil
}

func (r *mysqlRepository) GetById(id int) (models.Employee, error) {
	row := r.db.QueryRow(QueryGetEmployeeById, id)

	var employee models.Employee
	var warehouseID sql.NullInt64

	err := row.Scan(&employee.ID, &employee.CardNumberID, &employee.FirstName, &employee.LastName, &warehouseID)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Employee{}, fmt.Errorf("employee with id %d not found", id)
		}
		return models.Employee{}, fmt.Errorf("error scanning employee row: %w", err)
	}

	if warehouseID.Valid {
		employee.WarehouseID = int(warehouseID.Int64)
	}

	return employee, nil
}

func (r *mysqlRepository) Create(employee *models.Employee) error {
	var warehouseID sql.NullInt64
	if employee.WarehouseID != 0 {
		warehouseID = sql.NullInt64{Int64: int64(employee.WarehouseID), Valid: true}
	}

	result, err := r.db.Exec(QueryCreateEmployee, employee.CardNumberID, employee.FirstName, employee.LastName, warehouseID)
	if err != nil {
		return fmt.Errorf("error creating employee: %w", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error getting last insert ID: %w", err)
	}

	employee.ID = int(lastInsertID)
	return nil
}

func (r *mysqlRepository) ExistsByCardNumberId(cardNumberId string) (bool, error) {
	var exists bool

	err := r.db.QueryRow(QueryExistsByCardNumberId, cardNumberId).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking if employee exists with card_number_id %s: %w", cardNumberId, err)
	}

	return exists, nil
}

func (r *mysqlRepository) Update(id int, employee models.Employee) error {
	var warehouseID sql.NullInt64
	if employee.WarehouseID != 0 {
		warehouseID = sql.NullInt64{Int64: int64(employee.WarehouseID), Valid: true}
	}

	result, err := r.db.Exec(QueryUpdateEmployee, employee.CardNumberID, employee.FirstName, employee.LastName, warehouseID, id)
	if err != nil {
		return fmt.Errorf("error updating employee: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no employee found with id %d", id)
	}

	return nil
}

func (r *mysqlRepository) Delete(id int) error {
	result, err := r.db.Exec(QueryDeleteEmployee, id)
	if err != nil {
		return fmt.Errorf("error deleting employee with id %d: %w", id, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected for employee id %d: %w", id, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no employee found with id %d to delete", id)
	}

	return nil
}
