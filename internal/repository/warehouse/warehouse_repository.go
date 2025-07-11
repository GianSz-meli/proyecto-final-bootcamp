package warehouse

import "ProyectoFinal/pkg/models"

type WarehouseRepository interface {
	GetAll() []models.Warehouse
	GetById(id int) *models.Warehouse
	Create(warehouse models.Warehouse) models.Warehouse
	Update(id int, warehouse models.Warehouse) models.Warehouse
	ExistsByCode(code string) bool
	Delete(id int)
}
