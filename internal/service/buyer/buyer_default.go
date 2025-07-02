package buyer

import (
	"ProyectoFinal/internal/repository/buyer"
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
		return models.Buyer{}, fmt.Errorf("buyer with card number id %v already exists", buyerDTO.CardNumberId)
	}

	return s.repository.Save(models.DTOToBuyer(buyerDTO))
}

func (s *buyerService) GetById(id int) (models.Buyer, error) {
	return s.repository.GetById(id)
}

func (s *buyerService) GetAll() ([]models.Buyer, error) {
	return s.repository.GetAll()
}

func (s *buyerService) Update(id int, buyerDTO models.BuyerUpdateDTO) (models.Buyer, error) {
	if err := s.validator.Struct(buyerDTO); err != nil {
		return models.Buyer{}, fmt.Errorf("invalid input: %w", err)
	}

	existingBuyer, err := s.repository.GetById(id)
	if err != nil {
		return models.Buyer{}, err
	}

	if buyerDTO.CardNumberId != nil {
		if *buyerDTO.CardNumberId != existingBuyer.CardNumberId &&
			s.repository.ExistsByCardNumberId(*buyerDTO.CardNumberId) {
			return models.Buyer{}, fmt.Errorf("buyer with card number id %v already exists", *buyerDTO.CardNumberId)
		}
		existingBuyer.CardNumberId = *buyerDTO.CardNumberId
	}

	if buyerDTO.FirstName != nil {
		existingBuyer.FirstName = *buyerDTO.FirstName
	}

	if buyerDTO.LastName != nil {
		existingBuyer.LastName = *buyerDTO.LastName
	}

	updatedBuyer, err := s.repository.Update(existingBuyer)
	if err != nil {
		return models.Buyer{}, err
	}

	return updatedBuyer, nil
}

func (s *buyerService) Delete(id int) error {
	return s.repository.Delete(id)
}
