package buyer

import (
	"ProyectoFinal/pkg/models"
)

type Repository interface {
	Save(buyer models.Buyer) models.Buyer
	GetById(id int) (models.Buyer, bool)
	GetAll() []models.Buyer
	Update(buyer models.Buyer) models.Buyer
	Delete(id int)
	ExistsByCardNumberId(id string) bool
}
