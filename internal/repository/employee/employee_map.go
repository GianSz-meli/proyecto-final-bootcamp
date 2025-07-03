package employee

import (
	"ProyectoFinal/internal/repository/utils"
	"ProyectoFinal/pkg/models"
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

func (r *repository) GetById(id int) (models.Employee, bool) {
	employee, ok := r.employees[id]
	return employee, ok
}

func (r *repository) Create(employee *models.Employee) error {
	r.lastID++
	employee.ID = r.lastID
	r.employees[r.lastID] = *employee
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

func (r *repository) Update(id int, employee models.Employee) error {
	employee.ID = id
	r.employees[id] = employee
	return nil
}

func (r *repository) Delete(id int) {
	delete(r.employees, id)
}
