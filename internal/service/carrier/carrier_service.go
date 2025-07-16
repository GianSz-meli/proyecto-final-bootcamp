package carrier

import "ProyectoFinal/pkg/models"

type CarrierService interface {
	Create(carrier *models.Carrier) (*models.Carrier, error)
}
