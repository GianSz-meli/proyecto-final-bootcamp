package repository

import (
	pkgErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"fmt"
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
			newerror := fmt.Errorf("%w: the product code already exists", pkgErrors.ErrUnprocessableEntity)
			return models.Product{}, newerror
		}
	}
	newProd.ID = r.lastID + 1
	r.db[newProd.ID] = newProd
	return newProd, nil
}

func (r *ProductMap) FindAllProducts() (p map[int]models.Product, err error) {
	v := make(map[int]models.Product)

	for key, value := range r.db {
		v[key] = value
	}

	if len(v) == 0 {
		newerror := fmt.Errorf("%w: not exits products", pkgErrors.ErrNotFound)
		return nil, newerror
	}
	return v, nil
}

func (r *ProductMap) FindProductsById(id int) (models.Product, error) {

	for _, product := range r.db {
		if product.ID == id {
			return product, nil
		}
	}
	newerror := fmt.Errorf("%w: no product was found with this id", pkgErrors.ErrNotFound)
	return models.Product{}, newerror
}

func (r *ProductMap) UpdateProduct(id int, prod models.Product) (models.Product, error) {
	prod, ok := r.db[id]
	if !ok {
		newerror := fmt.Errorf("%w: no product was found with this id", pkgErrors.ErrNotFound)
		return models.Product{}, newerror
	}
	prod.ID = id
	r.db[id] = prod
	return prod, nil
}

func (r *ProductMap) DeleteProduct(id int) error {
	if _, exists := r.db[id]; !exists {
		return pkgErrors.ErrNotFound
	}
	delete(r.db, id)
	return nil
}
