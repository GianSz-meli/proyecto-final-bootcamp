package warehouse

import (
	"ProyectoFinal/internal/repository/utils"
	"ProyectoFinal/pkg/models"
)

type MemoryWarehouseRepository struct {
	db     map[int]models.Warehouse
	lastId int
}

func NewMemoryWarehouseRepository(db map[int]models.Warehouse) *MemoryWarehouseRepository {
	return &MemoryWarehouseRepository{
		db:     db,
		lastId: utils.GetLastId(db),
	}
}

func (r *MemoryWarehouseRepository) GetAll() ([]models.Warehouse, error) {
	warehouses := make([]models.Warehouse, 0, len(r.db))
	for _, warehouse := range r.db {
		warehouses = append(warehouses, warehouse)
	}
	return warehouses, nil
}

func (r *MemoryWarehouseRepository) GetById(id int) (*models.Warehouse, error) {
	warehouse, exists := r.db[id]
	if !exists {
		return nil, nil
	}
	return &warehouse, nil
}

func (r *MemoryWarehouseRepository) ExistsByCode(code string) (bool, error) {
	for _, warehouse := range r.db {
		if warehouse.WarehouseCode == code {
			return true, nil
		}
	}
	return false, nil
}

func (r *MemoryWarehouseRepository) Create(warehouse models.Warehouse) (models.Warehouse, error) {
	r.lastId++
	warehouse.ID = r.lastId
	r.db[r.lastId] = warehouse
	return warehouse, nil
}

func (r *MemoryWarehouseRepository) Update(id int, warehouse models.Warehouse) (models.Warehouse, error) {
	r.db[id] = warehouse
	return warehouse, nil
}

func (r *MemoryWarehouseRepository) Delete(id int) error {
	delete(r.db, id)
	return nil
}
