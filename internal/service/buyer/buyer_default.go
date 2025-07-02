package buyer

import (
	"ProyectoFinal/internal/repository/buyer"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"github.com/go-playground/validator/v10"
	"strconv"
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
	intId, err := strconv.Atoi(buyerDTO.CardNumberId)
	if err != nil {
		return models.Buyer{}, errors.ErrBadRequest
	}

	if s.repository.ExistsByCardNumberId(buyerDTO.CardNumberId) {
		return models.Buyer{}, errors.WrapErrAlreadyExist("buyer", "card_number_id", intId)
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
		return models.Buyer{}, errors.WrapErrBadRequest(err)
	}

	existingBuyer, err := s.repository.GetById(id)
	if err != nil {
		return models.Buyer{}, err
	}

	if buyerDTO.CardNumberId != nil {
		intId, err := strconv.Atoi(*buyerDTO.CardNumberId)
		if err != nil {
			return models.Buyer{}, errors.ErrBadRequest
		}
		if *buyerDTO.CardNumberId != existingBuyer.CardNumberId &&
			s.repository.ExistsByCardNumberId(*buyerDTO.CardNumberId) {
			return models.Buyer{}, errors.WrapErrAlreadyExist("buyer", "card_number_id", intId)
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
