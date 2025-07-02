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
	if s.repository.ExistsByCid(seller.Cid) {
		newError := errors.WrapErrAlreadyExist("seller", "cid", seller.Cid)
		return models.Seller{}, newError
	}

	if err := s.repository.Create(&seller); err != nil {
		return models.Seller{}, errors.ErrGeneral
	}

	return seller, nil
}

func (s *SellerDefault) GetAll() []models.Seller {
	sellers := s.repository.GetAll()
	return sellers

}
