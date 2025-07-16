package buyer

import (
	"ProyectoFinal/internal/repository/buyer"
	"ProyectoFinal/pkg/models"
)

type buyerService struct {
	repository buyer.Repository
}

// NewBuyerService creates and returns a new instance of the buyer service.
func NewBuyerService(newRepository buyer.Repository) Service {
	return &buyerService{
		repository: newRepository,
	}
}

// GetById retrieves a single buyer by their unique identifier.
func (s *buyerService) GetById(id int) (*models.Buyer, error) {
	return s.repository.GetById(id)
}

// GetAll retrieves all buyers from the system.
func (s *buyerService) GetAll() ([]*models.Buyer, error) {
	return s.repository.GetAll()
}

// Create adds a new buyer to the system.
func (s *buyerService) Create(buyer *models.Buyer) (*models.Buyer, error) {
	return s.repository.Create(buyer)
}

// Update modifies an existing buyer's information.
func (s *buyerService) Update(id int, buyer *models.Buyer) (*models.Buyer, error) {
	return s.repository.Update(buyer)
}

// Delete removes a buyer from the system.
func (s *buyerService) Delete(id int) error {
	return s.repository.Delete(id)
}

// GetByIdWithOrderCount retrieves a buyer along with their total purchase orders count.
func (s *buyerService) GetByIdWithOrderCount(id int) (*models.BuyerWithOrderCount, error) {
	return s.repository.GetByIdWithOrderCount(id)
}

// GetAllWithOrderCount retrieves all buyers along with their respective purchase orders count.
func (s *buyerService) GetAllWithOrderCount() ([]*models.BuyerWithOrderCount, error) {
	return s.repository.GetAllWithOrderCount()
}
