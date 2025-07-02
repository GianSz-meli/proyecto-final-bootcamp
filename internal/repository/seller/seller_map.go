package seller

import (
	"ProyectoFinal/pkg/models"
)

type SellerMap struct {
	db map[int]models.Seller
}

func NewSellerRepository(db map[int]models.Seller) SellerRepository {
	return &SellerMap{db: db}
}

func (r *SellerMap) Create(seller *models.Seller) error {
	id := len(r.db) + 1
	seller.Id = id
	r.db[id] = *seller
	return nil
}

func (r *SellerMap) GetById(id int) (models.Seller, bool) {
	s, ok := r.db[id]
	return s, ok
}

func (r *SellerMap) ExistsByCid(cid int) bool {
	sellers := r.db
	for _, s := range sellers {
		if s.Cid == cid {
			return true
		}
	}
	return false
}

func (r *SellerMap) GetAll() []models.Seller {
	var sellers []models.Seller
	for _, s := range r.db {
		sellers = append(sellers, s)
	}
	return sellers
}
