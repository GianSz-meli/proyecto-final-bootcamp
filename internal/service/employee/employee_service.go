package employee

import (
	"ProyectoFinal/pkg/models"
)

type Service interface {
	GetAll() ([]models.Employee, error)
	GetById(id int) (models.Employee, error)
	Create(employee models.Employee) (models.Employee, error)
	PatchUpdate(id int, updateRequest *models.EmployeeUpdateRequest) (models.Employee, error)
	Delete(id int) error
}
