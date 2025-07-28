package seller

import (
	"ProyectoFinal/internal/handler/utils"
	"ProyectoFinal/internal/repository/seller"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"fmt"
	"log"
)

type SellerDefault struct {
	repository seller.SellerRepository
}

func NewSellerService(repository seller.SellerRepository) SellerService {
	return &SellerDefault{repository: repository}
}

func (s *SellerDefault) Create(seller models.Seller) (models.Seller, error) {

	newSeller, err := s.repository.Create(seller)

	if err != nil {
		return models.Seller{}, err
	}

	return newSeller, nil
}

func (s *SellerDefault) GetAll() ([]models.Seller, error) {
	sellers, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}
	return sellers, nil

}

func (s *SellerDefault) GetById(id int) (models.Seller, error) {
	seller, err := s.repository.GetById(id)
	if err != nil {
		return models.Seller{}, err
	}
	return *seller, nil
}

func (s *SellerDefault) Delete(id int) error {

	if err := s.repository.Delete(id); err != nil {
		return err
	}

	return nil
}

func (s *SellerDefault) Update(id int, reqBody *models.UpdateSellerRequest) (models.Seller, error) {
	sellerToUpdate, err := s.repository.GetById(id)

	if err != nil {
		return models.Seller{}, err
	}
	if !utils.UpdateFields(sellerToUpdate, reqBody) {
		newError := fmt.Errorf("%w : no fields provided for update", errors.ErrUnprocessableEntity)
		log.Println(newError)
		return models.Seller{}, newError
	}
	_, err = s.repository.Update(sellerToUpdate)
	if err != nil {
		return models.Seller{}, err
	}

	return *sellerToUpdate, nil
}
