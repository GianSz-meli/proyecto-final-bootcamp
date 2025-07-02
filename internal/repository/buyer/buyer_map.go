package buyer

import (
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"encoding/json"
	"os"
	"sort"
)

type jsonRepository struct {
	db        map[int]models.Buyer
	idCounter int
	filePath  string
}

func NewBuyerJsonRepository(data map[int]models.Buyer, path string) Repository {
	return &jsonRepository{
		db:        data,
		idCounter: checkCounter(data),
		filePath:  path,
	}
}

func checkCounter(data map[int]models.Buyer) int {
	idCounter := 0
	for _, car := range data {
		if car.Id > idCounter {
			idCounter = car.Id
		}
	}
	return idCounter
}

func (r *jsonRepository) Save(buyer models.Buyer) (models.Buyer, error) {
	r.idCounter++
	buyer.Id = r.idCounter
	r.db[buyer.Id] = buyer

	if err := r.flush(); err != nil {
		delete(r.db, buyer.Id)
		r.idCounter--
		return models.Buyer{}, errors.ErrGeneral
	}

	return buyer, nil
}

func (r *jsonRepository) GetById(id int) (models.Buyer, error) {
	buyer, ok := r.db[id]
	if !ok {
		return models.Buyer{}, errors.ErrNotFound
	}
	return buyer, nil
}

func (r *jsonRepository) GetAll() ([]models.Buyer, error) {
	var data []models.Buyer

	for _, buyer := range r.db {
		data = append(data, buyer)
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].Id < data[j].Id
	})

	return data, nil
}

func (r *jsonRepository) Update(buyer models.Buyer) (models.Buyer, error) {
	prev, ok := r.db[buyer.Id]
	if !ok {
		return models.Buyer{}, errors.ErrNotFound
	}
	r.db[buyer.Id] = buyer

	if err := r.flush(); err != nil {
		r.db[buyer.Id] = prev
		return models.Buyer{}, errors.ErrGeneral
	}

	return buyer, nil
}

func (r *jsonRepository) Delete(id int) error {
	prev, ok := r.db[id]

	if !ok {
		return errors.ErrNotFound
	}

	delete(r.db, id)

	if err := r.flush(); err != nil {
		r.db[id] = prev
		return errors.ErrGeneral
	}

	return nil
}

func (r *jsonRepository) ExistsByCardNumberId(id string) bool {
	for _, buyer := range r.db {
		if buyer.CardNumberId == id {
			return true
		}
	}
	return false
}

func (r *jsonRepository) flush() error {
	var data []models.Buyer

	for _, buyer := range r.db {
		data = append(data, buyer)
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].Id < data[j].Id
	})

	newJson, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return errors.ErrGeneral
	}

	if err := os.WriteFile(r.filePath, newJson, 0644); err != nil {
		return errors.ErrGeneral
	}

	return nil
}
