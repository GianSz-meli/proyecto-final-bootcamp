package buyer

import (
	"ProyectoFinal/internal/repository/buyer"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type buyerService struct {
	repository buyer.Repository
	validator  *validator.Validate
}

func NewBuyerService(newRepository buyer.Repository) Service {
	return &buyerService{
		repository: newRepository,
		validator:  validator.New(),
	}
}

func (s *buyerService) Save(buyerDTO models.BuyerCreateDTO) (models.Buyer, error) {
	if s.repository.ExistsByCardNumberId(buyerDTO.CardNumberId) {
		return models.Buyer{}, errors.WrapErrAlreadyExistString("buyer", "card_number_id", buyerDTO.CardNumberId)
	}

	return s.repository.Save(models.DTOToBuyer(buyerDTO)), nil
}

func (s *buyerService) GetById(id int) (models.Buyer, error) {
	existingBuyer, ok := s.repository.GetById(id)
	if !ok {
		return models.Buyer{}, fmt.Errorf("buyer with id %d not found", id)
	}
	return existingBuyer, nil
}

func (s *buyerService) GetAll() []models.Buyer {
	return s.repository.GetAll()
}

func (s *buyerService) Update(id int, buyerDTO models.BuyerUpdateDTO) (models.Buyer, error) {
	if err := s.validator.Struct(buyerDTO); err != nil {
		return models.Buyer{}, errors.WrapErrBadRequest(err)
	}

	existingBuyer, ok := s.repository.GetById(id)
	if !ok {
		return models.Buyer{}, fmt.Errorf("buyer with id %d not found", id)
	}

	if buyerDTO.CardNumberId != nil {
		if *buyerDTO.CardNumberId != existingBuyer.CardNumberId &&
			s.repository.ExistsByCardNumberId(*buyerDTO.CardNumberId) {
			return models.Buyer{}, errors.WrapErrAlreadyExistString("buyer", "card_number_id", *buyerDTO.CardNumberId)
		}
		existingBuyer.CardNumberId = *buyerDTO.CardNumberId
	}

	if buyerDTO.FirstName != nil {
		existingBuyer.FirstName = *buyerDTO.FirstName
	}

	if buyerDTO.LastName != nil {
		existingBuyer.LastName = *buyerDTO.LastName
	}

	return s.repository.Update(existingBuyer), nil
}

func (s *buyerService) Delete(id int) error {
	_, ok := s.repository.GetById(id)
	if !ok {
		return fmt.Errorf("buyer with id %d not found", id)
	}

	s.repository.Delete(id)
	return nil
}
