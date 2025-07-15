package carrier

import (
	"ProyectoFinal/internal/repository/carrier"
	"ProyectoFinal/pkg/models"
)

type carrierService struct {
	carrierRepo carrier.CarrierRepository
}

func NewCarrierService(carrierRepo carrier.CarrierRepository) CarrierService {
	return &carrierService{
		carrierRepo: carrierRepo,
	}
}

func (s *carrierService) Create(carrier *models.Carrier) (*models.Carrier, error) {
	newCarrier, err := s.carrierRepo.Create(carrier)
	if err != nil {
		return nil, err
	}

	return newCarrier, nil
}
