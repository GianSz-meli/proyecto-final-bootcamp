package carrier

import (
	"ProyectoFinal/pkg/models"
)

type CarrierRepository interface {
	Create(carrier *models.Carrier) (*models.Carrier, error)
}
