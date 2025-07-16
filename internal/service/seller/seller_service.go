package seller

import "ProyectoFinal/pkg/models"

// SellerService defines business logic methods for managing sellers.
type SellerService interface {
	// Create adds a new seller.
	Create(seller models.Seller) (models.Seller, error)

	// GetAll retrieves all sellers.
	GetAll() ([]models.Seller, error)

	// GetById retrieves a seller by its ID.
	GetById(id int) (models.Seller, error)

	// Delete removes a seller by its ID.
	Delete(id int) error

	// Update modifies an existing seller.
	Update(seller models.Seller) (models.Seller, error)
}
