package warehouse

import (
	"ProyectoFinal/pkg/models"
)

type MemoryWarehouseRepository struct {
	db map[int]models.Warehouse
}

func NewMemoryWarehouseRepository(db map[int]models.Warehouse) *MemoryWarehouseRepository {
	return &MemoryWarehouseRepository{
		db: db,
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
	// Generate a new ID (simple implementation - in production you'd want a more robust solution)
	newID := 1
	for id := range r.db {
		if id >= newID {
			newID = id + 1
		}
	}

	warehouse.ID = newID
	r.db[newID] = warehouse
	return &warehouse
}

func (r *MemoryWarehouseRepository) Update(id int, warehouse models.Warehouse) *models.Warehouse {
	r.db[id] = warehouse
	return &warehouse
}
