package repository

import (
	"ProyectoFinal/pkg/models"
)

type ProductRepository interface {
	ExistsProdCode(prodCode string) bool
    CreateProduct(newProd models.Product) (models.Product, error)
    FindAllProducts() (map[int]models.Product, error)
    FindProductsById(id int) (models.Product, bool)
    UpdateProduct(id int, prod models.Product) (models.Product, error)
    DeleteProduct(id int)
}
