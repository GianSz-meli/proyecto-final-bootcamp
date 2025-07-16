package employee

import (
	pkgErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"database/sql"
	"errors"
)

type mysqlRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) Repository {
	return &mysqlRepository{
		db: db,
	}
}

// GetAll executes a query to retrieve all employee records from the database,
// scanning each row into Employee structs and returning them as a slice.
// Returns an error if the query fails or if there's an issue scanning the results.
func (r *mysqlRepository) GetAll() ([]models.Employee, error) {
	rows, err := r.db.Query(QueryGetAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		return nil, err
	}

	var employees []models.Employee
	for rows.Next() {
		var employee models.Employee

		err := rows.Scan(&employee.ID, &employee.CardNumberID, &employee.FirstName, &employee.LastName, &employee.WarehouseID)
		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return employees, nil
}

// GetById performs a database lookup to find an employee by their unique identifier.
// Returns the employee record if found, or a not found error if the ID doesn't exist.
func (r *mysqlRepository) GetById(id int) (models.Employee, error) {
	row := r.db.QueryRow(QueryGetById, id)

	var employee models.Employee

	err := row.Scan(&employee.ID, &employee.CardNumberID, &employee.FirstName, &employee.LastName, &employee.WarehouseID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Employee{}, pkgErrors.WrapErrNotFound("employee", "id", id)
		}
		return models.Employee{}, err
	}

	return employee, nil
}

// Create inserts a new employee record into the database and automatically assigns
// the generated primary key ID back to the employee object.
func (r *mysqlRepository) Create(employee *models.Employee) error {
	result, err := r.db.Exec(QueryCreate, employee.CardNumberID, employee.FirstName, employee.LastName, employee.WarehouseID)
	if err != nil {
		return err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	employee.ID = int(lastInsertID)
	return nil
}

// Update modifies an existing employee record in the database after verifying
// that the employee exists.
func (r *mysqlRepository) Update(id int, employee models.Employee) error {
	_, err := r.GetById(id)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(QueryUpdate, employee.CardNumberID, employee.FirstName, employee.LastName, employee.WarehouseID, id)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes an employee record from the database by ID and verifies that
// the deletion was successful by checking the affected rows count. Returns a
// not found error if no employee exists with the given ID.
func (r *mysqlRepository) Delete(id int) error {
	result, err := r.db.Exec(QueryDelete, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return pkgErrors.WrapErrNotFound("employee", "id", id)
	}

	return nil
}
