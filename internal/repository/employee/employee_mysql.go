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

func (r *mysqlRepository) GetAll() ([]models.Employee, error) {
	rows, err := r.db.Query(QueryGetAllEmployees)
	if err != nil {
		return nil, pkgErrors.HandleMysqlError(err)
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		return nil, pkgErrors.HandleMysqlError(err)
	}

	var employees []models.Employee
	for rows.Next() {
		var employee models.Employee

		err := rows.Scan(&employee.ID, &employee.CardNumberID, &employee.FirstName, &employee.LastName, &employee.WarehouseID)
		if err != nil {
			return nil, pkgErrors.HandleMysqlError(err)
		}

		employees = append(employees, employee)
	}

	return employees, nil
}

func (r *mysqlRepository) GetById(id int) (models.Employee, error) {
	row := r.db.QueryRow(QueryGetEmployeeById, id)

	var employee models.Employee

	err := row.Scan(&employee.ID, &employee.CardNumberID, &employee.FirstName, &employee.LastName, &employee.WarehouseID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Employee{}, pkgErrors.WrapErrNotFound("employee", "id", id)
		}
		return models.Employee{}, pkgErrors.HandleMysqlError(err)
	}

	return employee, nil
}

func (r *mysqlRepository) Create(employee *models.Employee) error {
	result, err := r.db.Exec(QueryCreateEmployee, employee.CardNumberID, employee.FirstName, employee.LastName, employee.WarehouseID)
	if err != nil {
		return pkgErrors.HandleMysqlError(err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return pkgErrors.HandleMysqlError(err)
	}

	employee.ID = int(lastInsertID)
	return nil
}

func (r *mysqlRepository) ExistsByCardNumberId(cardNumberId string) (bool, error) {
	var exists bool

	err := r.db.QueryRow(QueryExistsByCardNumberId, cardNumberId).Scan(&exists)
	if err != nil {
		return false, pkgErrors.HandleMysqlError(err)
	}

	return exists, nil
}

func (r *mysqlRepository) Update(id int, employee models.Employee) error {
	_, err := r.GetById(id)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(QueryUpdateEmployee, employee.CardNumberID, employee.FirstName, employee.LastName, employee.WarehouseID, id)
	if err != nil {
		return pkgErrors.HandleMysqlError(err)
	}

	return nil
}

func (r *mysqlRepository) Delete(id int) error {
	result, err := r.db.Exec(QueryDeleteEmployee, id)
	if err != nil {
		return pkgErrors.HandleMysqlError(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return pkgErrors.HandleMysqlError(err)
	}

	if rowsAffected == 0 {
		return pkgErrors.WrapErrNotFound("employee", "id", id)
	}

	return nil
}
