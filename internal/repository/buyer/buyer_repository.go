package buyer

import (
	"ProyectoFinal/pkg/models"
)

type Repository interface {
	Save(buyer models.Buyer) (models.Buyer, error)
	GetById(id int) (models.Buyer, error)
	GetAll() ([]models.Buyer, error)
	Update(buyer models.Buyer) (models.Buyer, error)
	Delete(id int) error
	ExistsByCardNumberId(id string) bool
}
