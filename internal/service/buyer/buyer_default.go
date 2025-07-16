package buyer

import (
	"ProyectoFinal/internal/handler/utils"
	"ProyectoFinal/internal/repository/buyer"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"fmt"
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

// PatchUpdate applies partial updates to a buyer, only modifying fields that are
// provided (non-nil) in the updateDTO. Returns error if no fields are provided.
func (s *buyerService) PatchUpdate(id int, updateDTO *models.BuyerUpdateDTO) (*models.Buyer, error) {
	buyerToUpdate, err := s.GetById(id)
	if err != nil {
		return nil, err
	}

	if !utils.UpdateFields(buyerToUpdate, updateDTO) {
		return nil, fmt.Errorf("%w : no fields provided for update", errors.ErrUnprocessableEntity)
	}

	return s.Update(id, buyerToUpdate)
}
