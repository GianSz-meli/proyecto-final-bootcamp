package buyer

import (
	"ProyectoFinal/internal/repository/buyer"
	"ProyectoFinal/pkg/models"
)

type buyerService struct {
	repository buyer.Repository
}

func NewBuyerService(newRepository buyer.Repository) Service {
	return &buyerService{
		repository: newRepository,
	}
}

func (s *buyerService) GetById(id int) (*models.Buyer, error) {
	return s.repository.GetById(id)
}

func (s *buyerService) GetAll() ([]*models.Buyer, error) {
	return s.repository.GetAll()
}

func (s *buyerService) Create(buyer *models.Buyer) (*models.Buyer, error) {
	return s.repository.Create(buyer)
}

func (s *buyerService) Update(id int, buyer *models.Buyer) (*models.Buyer, error) {
	return s.repository.Update(buyer)
}

func (s *buyerService) Delete(id int) error {
	return s.repository.Delete(id)
}
