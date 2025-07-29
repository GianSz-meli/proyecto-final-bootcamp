package products

import (
	"ProyectoFinal/pkg/models"
)

type ProductRepository interface {
	CreateProduct(newProd models.Product) (models.Product, error)
	FindAllProducts() (map[int]models.Product, error)
	FindProductsById(id int) (models.Product, error)
	UpdateProduct(id int, prod models.Product) (models.Product, error)
	DeleteProduct(id int) error     
}
