package buyer

import (
	"ProyectoFinal/pkg/models"
)

// Repository defines the contract for buyer data access operations using database connections.
type Repository interface {
	GetById(id int) (*models.Buyer, error)
	GetAll() ([]*models.Buyer, error)
	Create(buyer *models.Buyer) (*models.Buyer, error)
	Update(buyer *models.Buyer) (*models.Buyer, error)
	Delete(id int) error
	GetByIdWithOrderCount(id int) (*models.BuyerWithOrderCount, error)
	GetAllWithOrderCount() ([]*models.BuyerWithOrderCount, error)
}

// RepositoryMap defines the contract for buyer data access operations using in-memory storage.
type RepositoryMap interface {
	Create(buyer models.Buyer) models.Buyer
	GetById(id int) (models.Buyer, bool)
	GetAll() []models.Buyer
	Update(buyer models.Buyer) models.Buyer
	Delete(id int)
	ExistsByCardNumberId(id string) bool
}
