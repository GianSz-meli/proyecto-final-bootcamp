package warehouse

import "ProyectoFinal/pkg/models"

type WarehouseService interface {
	GetAllWarehouses() ([]models.Warehouse, error)
	GetWarehouseById(id int) (models.Warehouse, error)
	CreateWarehouse(warehouse models.Warehouse) (models.Warehouse, error)
	UpdateWarehouse(id int, warehouse models.UpdateWarehouseRequest) (models.Warehouse, error)
	DeleteWarehouse(id int) error
}
