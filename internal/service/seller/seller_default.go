package seller

import (
	"ProyectoFinal/internal/repository/seller"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"fmt"
)

type SellerDefault struct {
	repository seller.SellerRepository
}

func NewSellerService(repository seller.SellerRepository) SellerService {
	return &SellerDefault{repository: repository}
}

func (s *SellerDefault) Create(seller models.Seller) (models.Seller, error) {
	if s.repository.ExistsByCid(seller.Cid) {
		newError := errors.WrapErrConflict("seller", "cid", seller.Cid)
		return models.Seller{}, newError
	}

	s.repository.Create(&seller)

	return seller, nil
}

func (s *SellerDefault) GetAll() []models.Seller {
	sellers := s.repository.GetAll()
	return sellers

}

func (s *SellerDefault) GetById(id int) (models.Seller, error) {
	seller, ok := s.repository.GetById(id)
	if !ok {
		newError := fmt.Errorf("%w : seller with id %d not found", errors.ErrNotFound, id)
		return models.Seller{}, newError
	}

	return seller, nil
}

func (s *SellerDefault) Delete(id int) error {
	if _, ok := s.repository.GetById(id); !ok {
		newError := fmt.Errorf("%w : seller with id %d not found", errors.ErrNotFound, id)
		return newError
	}

	s.repository.Delete(id)

	return nil
}

func (s *SellerDefault) Update(id int, seller models.Seller) (models.Seller, error) {
	current, ok := s.repository.GetById(id)
	if !ok {
		newError := fmt.Errorf("%w : seller with id %d not found", errors.ErrNotFound, id)
		return models.Seller{}, newError
	}

	if current.Cid != seller.Cid {
		if s.repository.ExistsByCid(seller.Cid) {
			newError := errors.WrapErrConflict("seller", "cid", seller.Cid)
			return models.Seller{}, newError
		}
	}

	s.repository.Update(&seller)

	return seller, nil
}
