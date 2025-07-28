package buyer

import (
	"ProyectoFinal/pkg/models"
	"github.com/stretchr/testify/mock"
)

type MockBuyerRepository struct {
	mock.Mock
}

func (m *MockBuyerRepository) GetById(id int) (*models.Buyer, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Buyer), args.Error(1)
}

func (m *MockBuyerRepository) GetAll() ([]*models.Buyer, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.Buyer), args.Error(1)
}

func (m *MockBuyerRepository) Create(buyer *models.Buyer) (*models.Buyer, error) {
	args := m.Called(buyer)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Buyer), args.Error(1)
}

func (m *MockBuyerRepository) Update(buyer *models.Buyer) (*models.Buyer, error) {
	args := m.Called(buyer)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Buyer), args.Error(1)
}

func (m *MockBuyerRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockBuyerRepository) GetByIdWithOrderCount(id int) (*models.BuyerWithOrderCount, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.BuyerWithOrderCount), args.Error(1)
}

func (m *MockBuyerRepository) GetAllWithOrderCount() ([]*models.BuyerWithOrderCount, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.BuyerWithOrderCount), args.Error(1)
}
