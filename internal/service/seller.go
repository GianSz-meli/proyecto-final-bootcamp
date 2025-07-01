package service

import (
	"ProyectoFinal/internal/repository"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
)

type Service interface {
	Create(seller models.Seller) (models.Seller, error)
}

type SellerService struct {
	repository repository.Repository
}

func NewSellerService(repository repository.Repository) Service {
	return &SellerService{repository: repository}
}

func (s *SellerService) Create(seller models.Seller) (models.Seller, error) {
	if s.repository.ExistsByCid(seller.Cid) {
		newError := errors.WrapErrAlreadyExist("seller", "cid", seller.Cid)
		return models.Seller{}, newError
	}

	if err := s.repository.Create(&seller); err != nil {
		return models.Seller{}, errors.ErrGeneral
	}

	return seller, nil
}
