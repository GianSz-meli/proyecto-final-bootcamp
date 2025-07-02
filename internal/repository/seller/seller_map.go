package seller

import "ProyectoFinal/pkg/models"

type SellerMap struct {
	db map[int]models.Seller
}

func NewSellerRepository(db map[int]models.Seller) SellerRepository {
	return &SellerMap{db: db}
}

func (r *SellerMap) Create(seller *models.Seller) {
	id := len(r.db) + 1
	seller.Id = id
	r.db[id] = *seller
}

func (r *SellerMap) GetById(id int) (models.Seller, bool) {
	seller, ok := r.db[id]
	return seller, ok
}

func (r *SellerMap) ExistsByCid(cid int) bool {
	sellers := r.db
	for _, seller := range sellers {
		if seller.Cid == cid {
			return true
		}
	}
	return false
}

func (r *SellerMap) Update(seller *models.Seller) {
	r.db[seller.Id] = *seller
}
