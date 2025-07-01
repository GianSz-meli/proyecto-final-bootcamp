package repository

import (
	"ProyectoFinal/pkg/models"
)

type Repository interface {
	Create(seller *models.Seller) error
	GetById(id int) (models.Seller, bool)
	ExistsByCid(cid int) bool
	GetAll() map[int]models.Seller
}
type SellerRepository struct {
	db map[int]models.Seller
}

func NewSellerRepository(db map[int]models.Seller) Repository {
	return &SellerRepository{db: db}
}

func (r *SellerRepository) Create(seller *models.Seller) error {
	id := len(r.db) + 1
	seller.Id = id
	r.db[id] = *seller
	return nil
}

func (r *SellerRepository) GetById(id int) (models.Seller, bool) {
	seller, ok := r.db[id]
	return seller, ok
}

func (r *SellerRepository) ExistsByCid(cid int) bool {
	sellers := r.db
	for _, seller := range sellers {
		if seller.Cid == cid {
			return true
		}
	}
	return false
}

func (r *SellerRepository) GetAll() map[int]models.Seller {
	return r.db
}
