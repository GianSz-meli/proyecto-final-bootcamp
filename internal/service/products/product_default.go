package service

import (
	repository "ProyectoFinal/internal/repository/products"
	"ProyectoFinal/pkg/models"
)

func NewProductDefault(rp repository.ProductRepository) *ProductDefault {
	return &ProductDefault{rp: rp}
}

type ProductDefault struct {
	// rp is the repository that will be used by the service
	rp repository.ProductRepository
}

func (s *ProductDefault) CreateProduct(newProd models.Product) (models.Product, error) {
	return s.rp.CreateProduct(newProd)
}

func (s *ProductDefault) FindAllProducts() (p map[int]models.Product, err error) {
	return s.rp.FindAllProducts()
}

func (s *ProductDefault) FindProductsById(id int) (models.Product, error) {
	return s.rp.FindProductsById(id)
}

func (s *ProductDefault) UpdateProduct(id int, prod models.Product) (models.Product, error) {
	return s.rp.UpdateProduct(id, prod)
}

func (s *ProductDefault) DeleteProduct(id int) error {
	return s.rp.DeleteProduct(id)
}
