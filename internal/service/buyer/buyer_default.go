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

func (s *buyerService) GetById(id int) (*models.Buyer, error) {
	foundBuyer, err := s.repository.GetById(id)
	if err != nil {
		return nil, err
	}

	return foundBuyer, nil
}

func (s *buyerService) GetAll() ([]*models.Buyer, error) {
	buyers, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return buyers, nil
}

func (s *buyerService) Create(buyer *models.Buyer) (*models.Buyer, error) {
	exists, err := s.repository.ExistsByCardNumberId(buyer.CardNumberId)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.WrapErrConflict("buyer", "card number id", buyer.CardNumberId)
	}

	newBuyer, createErr := s.repository.Create(buyer)
	if createErr != nil {
		return nil, createErr
	}

	return newBuyer, nil
}

func (s *buyerService) Update(id int, buyer *models.Buyer) (*models.Buyer, error) {
	existingBuyer, alreadyExistsErr := s.repository.GetById(id)
	if alreadyExistsErr != nil {
		return nil, alreadyExistsErr
	}

	if buyer.CardNumberId != existingBuyer.CardNumberId {
		exists, err := s.repository.ExistsByCardNumberId(buyer.CardNumberId)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, errors.WrapErrConflict("buyer", "card number id", buyer.CardNumberId)
		}
	}

	updatedBuyer, updateErr := s.repository.Update(buyer)
	if updateErr != nil {
		return nil, updateErr
	}

	return updatedBuyer, nil
}

func (s *buyerService) Delete(id int) error {
	if _, err := s.repository.GetById(id); err != nil {
		return err
	}

	if deleteErr := s.repository.Delete(id); deleteErr != nil {
		return deleteErr
	}

	return nil
}
