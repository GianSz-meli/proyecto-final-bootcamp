package employee

import (
	employeemodel "ProyectoFinal/pkg/models/employee"
)

type repository struct {
	employees map[int]employeemodel.Employee
}

func NewRepository(employees map[int]employeemodel.Employee) Repository {
	return &repository{
		employees: employees,
	}
}

func (r *repository) GetAll() ([]employeemodel.Employee, error) {
	employees := make([]employeemodel.Employee, 0, len(r.employees))

	for _, employee := range r.employees {
		employees = append(employees, employee)
	}

	return employees, nil
}

func (r *repository) GetById(id int) (employeemodel.Employee, bool) {
	employee, ok := r.employees[id]
	return employee, ok
}

func (r *repository) Create(employee *employeemodel.Employee) error {
	id := len(r.employees) + 1
	employee.ID = id
	r.employees[id] = *employee
	return nil
}

func (r *repository) ExistsByCardNumberId(cardNumberId string) bool {
	for _, employee := range r.employees {
		if employee.CardNumberID == cardNumberId {
			return true
		}
	}
	return false
}

func (r *repository) Update(id int, employee employeemodel.Employee) error {
	employee.ID = id
	r.employees[id] = employee
	return nil
}

func (r *repository) Delete(id int) bool {
	_, exists := r.employees[id]
	if exists {
		delete(r.employees, id)
	}
	return exists
}
