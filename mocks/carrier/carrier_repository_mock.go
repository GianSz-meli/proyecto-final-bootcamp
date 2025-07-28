package carrier

import (
	"ProyectoFinal/pkg/models"

	"github.com/stretchr/testify/mock"
)

type MockCarrierRepository struct {
	mock.Mock
}

func (m *MockCarrierRepository) Create(carrier *models.Carrier) (*models.Carrier, error) {
	args := m.Called(carrier)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Carrier), args.Error(1)
}
