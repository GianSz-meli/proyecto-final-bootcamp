package warehouse

import "ProyectoFinal/pkg/models"

type WarehouseRepository interface {
	GetAll() ([]models.Warehouse, error)
	GetById(id int) (*models.Warehouse, error)
	Create(warehouse models.Warehouse) (models.Warehouse, error)
	Update(id int, warehouse models.Warehouse) (models.Warehouse, error)
	ExistsByCode(code string) (bool, error)
	Delete(id int) error
}
