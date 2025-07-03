package employee

import (
	"ProyectoFinal/pkg/models"
)

type Repository interface {
	GetAll() ([]models.Employee, error)
	GetById(id int) (models.Employee, bool)
	Create(employee *models.Employee) error
	ExistsByCardNumberId(cardNumberId string) bool
	ExistsByCardNumberIdExcludingID(cardNumberId string, excludeID int) bool
	Update(id int, employee models.Employee) error
	Delete(id int)
}
