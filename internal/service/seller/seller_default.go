package seller

import (
	"ProyectoFinal/internal/repository/seller"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
)

type SellerDefault struct {
	repository seller.SellerRepository
}

func NewSellerService(repository seller.SellerRepository) SellerService {
	return &SellerDefault{repository: repository}
}

func (s *SellerDefault) Create(seller models.Seller) (models.Seller, error) {
	exists, err := s.repository.ExistsByCid(seller.Cid)
	if err != nil {
		return models.Seller{}, err
	}
	if exists {
		newError := errors.WrapErrAlreadyExist("seller", "cid", seller.Cid)
		return models.Seller{}, newError
	}

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

func (s *SellerDefault) Update(id int, seller models.Seller) (models.Seller, error) {
	current, err := s.repository.GetById(id)
	if err != nil {
		return models.Seller{}, err
	}

	if current.Cid != seller.Cid {
		exists, err := s.repository.ExistsByCid(seller.Cid)
		if err != nil {
			return models.Seller{}, err
		}
		if exists {
			newError := errors.WrapErrAlreadyExist("seller", "cid", seller.Cid)
			return models.Seller{}, newError
		}
	}

	seller, err = s.repository.Update(&seller)

	if err != nil {
		return models.Seller{}, err
	}

	return seller, nil
}
