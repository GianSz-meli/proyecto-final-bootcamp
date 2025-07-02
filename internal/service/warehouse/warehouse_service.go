package warehouse

import "ProyectoFinal/pkg/models"

type WarehouseService interface {
	GetAllWarehouses() []models.Warehouse
	GetWarehouseById(id int) (*models.Warehouse, error)
	CreateWarehouse(warehouse models.Warehouse) (*models.Warehouse, error)
	UpdateWarehouse(id int, warehouse models.Warehouse) (*models.Warehouse, error)
	DeleteWarehouse(id int) error
}
