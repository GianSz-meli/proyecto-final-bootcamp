package buyer

import "ProyectoFinal/pkg/models"

type Service interface {
	Save(buyerDto models.BuyerCreateDto) error
	GetById(id int) (models.Buyer, error)
	GetAll() []models.Buyer
	Update(buyerUpdateDto models.BuyerUpdateDto) error
	Delete(id int) error
}
