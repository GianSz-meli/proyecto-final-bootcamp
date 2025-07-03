package buyer

import (
	"ProyectoFinal/internal/repository/buyer"
	"ProyectoFinal/pkg/errors"
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

func (s *buyerService) Save(buyer models.Buyer) (models.Buyer, error) {
	if s.repository.ExistsByCardNumberId(buyer.CardNumberId) {
		return models.Buyer{}, errors.WrapErrAlreadyExist("buyer", "card number id", buyer.CardNumberId)
	}

	return s.repository.Save(buyer), nil
}

func (s *buyerService) GetById(id int) (models.Buyer, error) {
	existingBuyer, ok := s.repository.GetById(id)
	if !ok {
		return models.Buyer{}, errors.WrapErrNotFound("buyer", "id", id)
	}

	return existingBuyer, nil
}

func (s *buyerService) GetAll() []models.Buyer {
	return s.repository.GetAll()
}

func (s *buyerService) Update(id int, buyer models.Buyer) (models.Buyer, error) {
	existingBuyer, ok := s.repository.GetById(id)
	if !ok {
		return models.Buyer{}, errors.WrapErrNotFound("buyer", "id", id)
	}

	if buyer.CardNumberId != existingBuyer.CardNumberId && s.repository.ExistsByCardNumberId(buyer.CardNumberId) {
		return models.Buyer{}, errors.WrapErrAlreadyExist("buyer", "card number id", buyer.CardNumberId)
	}

	return s.repository.Update(buyer), nil
}

func (s *buyerService) Delete(id int) error {
	_, ok := s.repository.GetById(id)
	if !ok {
		return errors.WrapErrNotFound("buyer", "id", id)
	}

	s.repository.Delete(id)
	return nil
}
