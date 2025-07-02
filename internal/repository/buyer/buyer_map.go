package buyer

import (
	"ProyectoFinal/pkg/models"
	"encoding/json"
	"fmt"
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
		return models.Buyer{}, err
	}

	return buyer, nil
}

func (r *jsonRepository) GetById(id int) (models.Buyer, error) {
	buyer, ok := r.db[id]
	if !ok {
		return models.Buyer{}, fmt.Errorf("buyer with id %v not found", id)
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
	prev, _ := r.db[buyer.Id]
	r.db[buyer.Id] = buyer

	if err := r.flush(); err != nil {
		r.db[buyer.Id] = prev
		return models.Buyer{}, err
	}

	return buyer, nil
}

func (r *jsonRepository) Delete(id int) error {
	prev, ok := r.db[id]

	if !ok {
		return fmt.Errorf("buyer with id %v not found", id)
	}

	delete(r.db, id)

	if err := r.flush(); err != nil {
		r.db[id] = prev
		return err
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
		return fmt.Errorf("failed to marshal json : %w", err)
	}

	return os.WriteFile(r.filePath, newJson, 0644)
}
