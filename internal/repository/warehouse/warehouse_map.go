package warehouse

import (
	"ProyectoFinal/pkg/models"
)

type MemoryWarehouseRepository struct {
	db map[int]models.Warehouse
	lastId int
}

func NewMemoryWarehouseRepository(db map[int]models.Warehouse) *MemoryWarehouseRepository {
	return &MemoryWarehouseRepository{
		db: db,
		lastId: GetLastId(db),
	}
}

func (r *MemoryWarehouseRepository) GetAll() []models.Warehouse {
	warehouses := make([]models.Warehouse, 0, len(r.db))
	for _, warehouse := range r.db {
		warehouses = append(warehouses, warehouse)
	}
	return warehouses
}

func (r *MemoryWarehouseRepository) GetById(id int) *models.Warehouse {
	warehouse, exists := r.db[id]
	if !exists {
		return nil
	}
	return &warehouse
}

func (r *MemoryWarehouseRepository) ExistsByCode(code string) bool {
	for _, warehouse := range r.db {
		if *warehouse.WarehouseCode == code {
			return true
		}
	}
	return false
}

func (r *MemoryWarehouseRepository) Create(warehouse models.Warehouse) *models.Warehouse {
	r.lastId++
	warehouse.ID = r.lastId
	r.db[r.lastId] = warehouse
	return &warehouse
}

func (r *MemoryWarehouseRepository) Update(id int, warehouse models.Warehouse) *models.Warehouse {
	r.db[id] = warehouse
	return &warehouse
}

func GetLastId(db map[int]models.Warehouse) int {
	lastId := 0
	for id := range db {
		if id > lastId {
			lastId = id
		}
	}
	return lastId
}

func (r *MemoryWarehouseRepository) Delete(id int) {
	panic("unimplemented")
}