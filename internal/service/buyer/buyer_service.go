package buyer

import "ProyectoFinal/pkg/models"

// Service defines the business logic contract for buyer operations.
type Service interface {
	GetById(id int) (*models.Buyer, error)
	GetAll() ([]*models.Buyer, error)
	Create(buyer *models.Buyer) (*models.Buyer, error)
	Update(id int, buyer *models.Buyer) (*models.Buyer, error)
	Delete(id int) error
	GetByIdWithOrderCount(id int) (*models.BuyerWithOrderCount, error)
	GetAllWithOrderCount() ([]*models.BuyerWithOrderCount, error)
}
