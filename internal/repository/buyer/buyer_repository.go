package buyer

import "ProyectoFinal/pkg/models"

type Repository interface {
	Save(buyer models.Buyer) error
	GetById(id int) (models.Buyer, error)
	GetAll() []models.Buyer
	Update(buyer models.Buyer) error
	Delete(id int) error
}
