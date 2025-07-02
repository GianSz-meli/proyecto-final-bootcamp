package repository

import (
	"ProyectoFinal/pkg/models"
)

type ProductRepository interface {
	CreateProduct(newProd models.Product) (models.Product, error)
	FindAllProducts()(p map[int]models.Product, err error)
	FindProductsById(id int) (models.Product, error) 
}
