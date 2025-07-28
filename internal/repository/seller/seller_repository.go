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
