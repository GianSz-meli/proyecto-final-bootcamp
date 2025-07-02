package repository

import (
	"ProyectoFinal/pkg/models"
	"errors"
)

func NewProductMap(db map[int]models.Product) *ProductMap {
	return &ProductMap{db: db, lastID: len(db)}
}


type ProductMap struct {
	db     map[int]models.Product
	lastID int
}


func (r *ProductMap) CreateProduct(newProd models.Product) (models.Product, error) {
	for _, product := range r.db {
		if product.ProductCode == newProd.ProductCode {
			return models.Product{}, errors.New("the product code already exists")
		}
	}
	newProd.ID = r.lastID + 1
	r.db[newProd.ID] = newProd
	return newProd, nil
}


func (r *ProductMap) FindAllProducts()(p map[int]models.Product, err error) {
	p = make(map[int]models.Product)
	for key, value := range r.db {
		p[key] = value
	}
	return
}
func (r *ProductMap) FindProductsById(id int) (models.Product, error) {

	for _, product := range r.db {
		if product.ID == id {
			return product, nil
		}
	}
	//maanejar error 
	return models.Product{}, errors.New("no vehicle was found with this id")
}