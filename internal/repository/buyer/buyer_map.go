package buyer

import (
	"ProyectoFinal/pkg/models"
	"sort"
)

type jsonRepository struct {
	db        map[int]models.Buyer
	idCounter int
}

func NewBuyerJsonRepository(data map[int]models.Buyer) Repository {
	return &jsonRepository{
		db:        data,
		idCounter: checkCounter(data),
	}
}

func checkCounter(data map[int]models.Buyer) int {
	idCounter := 0
	for _, buyer := range data {
		if buyer.Id > idCounter {
			idCounter = buyer.Id
		}
	}
	return idCounter
}

func (r *jsonRepository) Save(buyer models.Buyer) models.Buyer {
	r.idCounter++
	buyer.Id = r.idCounter
	r.db[buyer.Id] = buyer

	return buyer
}

func (r *jsonRepository) GetById(id int) (models.Buyer, bool) {
	buyer, ok := r.db[id]
	return buyer, ok
}

func (r *jsonRepository) GetAll() []models.Buyer {
	var data []models.Buyer

	for _, buyer := range r.db {
		data = append(data, buyer)
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].Id < data[j].Id
	})

	return data
}

func (r *jsonRepository) Update(buyer models.Buyer) models.Buyer {
	r.db[buyer.Id] = buyer
	return buyer
}

func (r *jsonRepository) Delete(id int) {
	delete(r.db, id)
}

func (r *jsonRepository) ExistsByCardNumberId(id string) bool {
	for _, buyer := range r.db {
		if buyer.CardNumberId == id {
			return true
		}
	}
	return false
}
