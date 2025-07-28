package locality

import (
	"ProyectoFinal/pkg/models"
	"github.com/stretchr/testify/mock"
)

// MockLocalityService - implementation of locality interface
type MockLocalityService struct {
	mock.Mock
}

func (m *MockLocalityService) Create(locality models.Locality) (models.Locality, error) {
	args := m.Called(locality)
	return args.Get(0).(models.Locality), args.Error(1)
}

func (m *MockLocalityService) GetSellersByLocalities() ([]models.SellersByLocalityReport, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.SellersByLocalityReport), args.Error(1)
}

func (m *MockLocalityService) GetSellersByIdLocality(idLocality int) (models.SellersByLocalityReport, error) {
	args := m.Called(idLocality)
	return args.Get(0).(models.SellersByLocalityReport), args.Error(1)
}

func (m *MockLocalityService) ReportCarriersByLocality(id *int) ([]models.CarrierReport, error) {
	args := m.Called(id)
	return args.Get(0).([]models.CarrierReport), args.Error(1)
}
