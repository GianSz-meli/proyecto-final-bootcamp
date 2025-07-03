package repository

import (
	"ProyectoFinal/internal/repository/utils"
	"ProyectoFinal/pkg/models"
)

func NewProductMap(db map[int]models.Product) *ProductMap {
	return &ProductMap{db: db, lastID: utils.GetLastId[models.Product](db)}
}

type ProductMap struct {
	db     map[int]models.Product
	lastID int
}

func (r *ProductMap) ExistsProdCode(prodCode string) bool {
	for _, product := range r.db {
		if product.ProductCode == prodCode {
			return true
		}
	}
	return false
}

func (r *ProductMap) CreateProduct(newProd models.Product) (models.Product, error) {
	r.lastID++
	newProd.ID = r.lastID
	r.db[newProd.ID] = newProd
	return newProd, nil
}

func (r *ProductMap) FindAllProducts() (p map[int]models.Product, err error) {
	v := make(map[int]models.Product)
	for key, value := range r.db {
		v[key] = value
	}
	return v, nil
}

func (r *ProductMap) FindProductsById(id int) (models.Product, bool) {
	product, ok := r.db[id]
	return product, ok
}

func (r *ProductMap) UpdateProduct(id int, prod models.Product) (models.Product, error) {
	prod.ID = id
	r.db[id] = prod
	return prod, nil
}

func (r *ProductMap) DeleteProduct(id int) {
	delete(r.db, id)
}
