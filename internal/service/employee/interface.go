package employee

import (
	employeemodel "ProyectoFinal/pkg/models/employee"
)

type Service interface {
	GetAll() ([]employeemodel.Employee, error)
	GetById(id int) (employeemodel.Employee, error)
	Create(employee employeemodel.Employee) (employeemodel.Employee, error)
	Update(id int, employee employeemodel.Employee) (employeemodel.Employee, error)
	Delete(id int) error
}
