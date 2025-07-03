package seller

import (
	"ProyectoFinal/internal/repository/utils"
	"ProyectoFinal/pkg/models"
)

type SellerMap struct {
	db     map[int]models.Seller
	lastId int
}

func NewSellerRepository(db map[int]models.Seller) SellerRepository {
	return &SellerMap{
		db:     db,
		lastId: utils.GetLastId[models.Seller](db),
	}
}

func (r *SellerMap) Create(seller *models.Seller) {
	r.lastId++
	id := r.lastId
	seller.Id = id
	r.db[id] = *seller
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

func (r *SellerMap) Delete(id int) {
	delete(r.db, id)
}

func (r *SellerMap) Update(seller *models.Seller) {
	r.db[seller.Id] = *seller
}
