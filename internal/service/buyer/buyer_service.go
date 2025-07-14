package buyer

import "ProyectoFinal/pkg/models"

type Service interface {
	GetById(id int) (*models.Buyer, error)
	GetAll() ([]*models.Buyer, error)
	Create(buyer *models.Buyer) (*models.Buyer, error)
	Update(id int, buyer *models.Buyer) (*models.Buyer, error)
	Delete(id int) error
}
