package service

import "ProyectoFinal/pkg/models"

type ProductService interface {
	CreateProduct(newProd models.Product) (models.Product, error)	
}





