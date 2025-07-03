package employee

import (
	"ProyectoFinal/pkg/models"
)

type Service interface {
	GetAll() ([]models.Employee, error)
	GetById(id int) (models.Employee, error)
	Create(employee models.Employee) (models.Employee, error)
	Update(id int, employee models.Employee) (models.Employee, error)
	Delete(id int) error
}
