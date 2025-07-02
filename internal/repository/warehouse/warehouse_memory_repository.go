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
