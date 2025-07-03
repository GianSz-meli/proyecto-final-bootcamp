package buyer

import "ProyectoFinal/pkg/models"

type Service interface {
	Save(buyer models.Buyer) (models.Buyer, error)
	GetById(id int) (models.Buyer, error)
	GetAll() []models.Buyer
	Update(id int, buyerDTO models.BuyerUpdateDTO) (models.Buyer, error)
	Delete(id int) error
}
