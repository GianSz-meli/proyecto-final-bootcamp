package service

import (
	repository "ProyectoFinal/internal/repository/products"
	pkgErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"fmt"
)

func NewProductDefault(rp repository.ProductRepository) *ProductDefault {
	return &ProductDefault{rp: rp}
}

type ProductDefault struct {
	// rp is the repository that will be used by the service
	rp repository.ProductRepository
}

func (s *ProductDefault) CreateProduct(newProd models.Product) (models.Product, error) {
	if s.rp.ExistsProdCode(newProd.ProductCode) {
		newError := pkgErrors.WrapErrConflict("product", "product code", newProd.ProductCode)
		return models.Product{}, newError
	}
	s.rp.CreateProduct(newProd)
	return newProd, nil
}

func (s *ProductDefault) FindAllProducts() (p map[int]models.Product, err error) {
	return s.rp.FindAllProducts()
}

func (s *ProductDefault) FindProductsById(id int) (models.Product, error) {
	product, err := s.rp.FindProductsById(id)
	if err != nil {
		newError := fmt.Errorf("%w : product with id %d not found", pkgErrors.ErrNotFound, id)
		return models.Product{}, newError
	}
	return product, nil
}

func (s *ProductDefault) UpdateProduct(id int, prod models.Product) (models.Product, error) {
	currentProd, err := s.rp.FindProductsById(id)
	if err != nil {
		newError := fmt.Errorf("%w : product with id %d not found", pkgErrors.ErrNotFound, id)
		return models.Product{}, newError
	}
	if currentProd.ProductCode != prod.ProductCode {
		if s.rp.ExistsProdCode(prod.ProductCode) {
			newError := pkgErrors.WrapErrConflict("product", "product code", prod.ProductCode)
			return models.Product{}, newError
		}
	}
	return s.rp.UpdateProduct(id, prod)
}

func (s *ProductDefault) DeleteProduct(id int) error {
	_, err := s.rp.FindProductsById(id)
	if err != nil {
		newError := fmt.Errorf("%w : product with id %d not found", pkgErrors.ErrNotFound, id)
		return newError
	}
	s.rp.DeleteProduct(id)
	return nil
}
