package locality

import (
	"ProyectoFinal/pkg/models"
	"github.com/stretchr/testify/mock"
)

type MockLocalityRepository struct {
	mock.Mock
}

func (m *MockLocalityRepository) Create(locality models.Locality) (models.Locality, error) {
	args := m.Called(locality)
	return args.Get(0).(models.Locality), args.Error(1)
}

func (m *MockLocalityRepository) GetSellersByIdLocality(idLocality int) (models.SellersByLocalityReport, error) {
	args := m.Called(idLocality)
	return args.Get(0).(models.SellersByLocalityReport), args.Error(1)
}

func (m *MockLocalityRepository) GetSellersByLocalities() ([]models.SellersByLocalityReport, error) {
	args := m.Called()
	return args.Get(0).([]models.SellersByLocalityReport), args.Error(1)
}

func (m *MockLocalityRepository) ReportCarriersByLocality(id *int) ([]models.CarrierReport, error) {
	args := m.Called(id)
	return args.Get(0).([]models.CarrierReport), args.Error(1)
}
