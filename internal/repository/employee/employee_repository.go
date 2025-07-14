package employee

import (
	"ProyectoFinal/pkg/models"
)

type Repository interface {
	GetAll() ([]models.Employee, error)
	GetById(id int) (models.Employee, error)
	Create(employee *models.Employee) error
	ExistsByCardNumberId(cardNumberId string) (bool, error)
	Update(id int, employee models.Employee) error
	Delete(id int) error
}
