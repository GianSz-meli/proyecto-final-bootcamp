package warehouse

import "ProyectoFinal/pkg/models"

type WarehouseRepository interface {
	GetAll() []models.Warehouse
}
