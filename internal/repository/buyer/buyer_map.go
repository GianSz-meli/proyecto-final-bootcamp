package buyer

import (
	"ProyectoFinal/internal/repository/utils"
	"ProyectoFinal/pkg/models"
	"sort"
)

type buyerMap struct {
	db        map[int]models.Buyer
	idCounter int
}

func NewBuyerRepository(data map[int]models.Buyer) Repository {
	return &buyerMap{
		db:        data,
		idCounter: utils.GetLastId(data),
	}
}

func (r *buyerMap) Create(buyer models.Buyer) models.Buyer {
	r.idCounter++
	buyer.Id = r.idCounter
	r.db[buyer.Id] = buyer
	return buyer
}

func (r *buyerMap) GetById(id int) (models.Buyer, bool) {
	buyer, ok := r.db[id]
	return buyer, ok
}

func (r *buyerMap) GetAll() []models.Buyer {
	var data []models.Buyer
	for _, buyer := range r.db {
		data = append(data, buyer)
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].Id < data[j].Id
	})

	return data
}

func (r *buyerMap) Update(buyer models.Buyer) models.Buyer {
	r.db[buyer.Id] = buyer
	return buyer
}

func (r *buyerMap) Delete(id int) {
	delete(r.db, id)
}

func (r *buyerMap) ExistsByCardNumberId(id string) bool {
	for _, buyer := range r.db {
		if buyer.CardNumberId == id {
			return true
		}
	}
	return false
}
