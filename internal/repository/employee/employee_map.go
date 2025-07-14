package employee

import (
	"ProyectoFinal/internal/repository/utils"
	"ProyectoFinal/pkg/models"
	"fmt"
)

type repository struct {
	employees map[int]models.Employee
	lastID    int
}

func NewRepository(employees map[int]models.Employee) Repository {
	repo := &repository{
		employees: employees,
	}
	repo.lastID = utils.GetLastId[models.Employee](employees)
	return repo
}

func (r *repository) GetAll() ([]models.Employee, error) {
	employees := make([]models.Employee, 0, len(r.employees))

	for _, employee := range r.employees {
		employees = append(employees, employee)
	}

	return employees, nil
}

func (r *repository) GetById(id int) (models.Employee, error) {
	if employee, exists := r.employees[id]; exists {
		return employee, nil
	}
	return models.Employee{}, fmt.Errorf("employee with id %d not found", id)
}

func (r *repository) Create(employee *models.Employee) error {
	r.lastID++
	employee.ID = r.lastID
	r.employees[r.lastID] = *employee
	return nil
}

func (r *repository) ExistsByCardNumberId(cardNumberId string) (bool, error) {
	for _, employee := range r.employees {
		if employee.CardNumberID == cardNumberId {
			return true, nil
		}
	}
	return false, nil
}

func (r *repository) Update(id int, employee models.Employee) error {
	if _, exists := r.employees[id]; !exists {
		return fmt.Errorf("employee with id %d not found", id)
	}
	employee.ID = id
	r.employees[id] = employee
	return nil
}

func (r *repository) Delete(id int) error {
	if _, exists := r.employees[id]; !exists {
		return fmt.Errorf("employee with id %d not found", id)
	}
	delete(r.employees, id)
	return nil
}
