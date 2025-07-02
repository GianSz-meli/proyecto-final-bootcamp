package service 

import (
	"ProyectoFinal/pkg/models"
	"ProyectoFinal/internal/repository/products"
)

func NewProductDefault (rp repository.ProductRepository) *ProductDefault {
	return &ProductDefault{rp: rp}
}

type ProductDefault struct {
	// rp is the repository that will be used by the service
	rp repository.ProductRepository
}


func (s *ProductDefault) CreateProduct(newProd models.Product) (models.Product, error){
	return s.rp.CreateProduct(newProd)
}