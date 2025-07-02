package repository

import (
	"ProyectoFinal/pkg/models"
)

type ProductRepository interface {
	CreateProduct(newProd models.Product) (models.Product, error)
}
