package warehouse

import "ProyectoFinal/pkg/models"

type WarehouseService interface {
	GetAllWarehouses() []models.Warehouse
}
