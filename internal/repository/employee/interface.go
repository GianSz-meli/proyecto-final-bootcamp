package employee

import (
	employeemodel "ProyectoFinal/pkg/models/employee"
)

type Repository interface {
	GetAll() ([]employeemodel.Employee, error)
	GetById(id int) (employeemodel.Employee, bool)
	Create(employee *employeemodel.Employee) error
	ExistsByCardNumberId(cardNumberId string) bool
	Update(id int, employee employeemodel.Employee) error
	Delete(id int) bool
}
