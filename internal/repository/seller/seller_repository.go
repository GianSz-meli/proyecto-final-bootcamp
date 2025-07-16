package seller

import "ProyectoFinal/pkg/models"

// SellerRepository defines methods for CRUD operations on the seller model with persistent storage.
type SellerRepository interface {
	Create(seller models.Seller) (models.Seller, error)
	GetById(id int) (*models.Seller, error)
	Update(seller *models.Seller) (models.Seller, error)
	GetAll() ([]models.Seller, error)
	Delete(id int) error
}

// SellerRepositoryMap defines in-memory operations for managing sellers using a map-based repository.
type SellerRepositoryMap interface {
	Create(seller *models.Seller)
	GetById(id int) (models.Seller, bool)
	ExistsByCid(cid string) bool
	Update(seller *models.Seller)
	GetAll() []models.Seller
	Delete(id int)
}
