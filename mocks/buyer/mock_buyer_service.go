package buyer

import (
	"ProyectoFinal/pkg/models"
	"github.com/stretchr/testify/mock"
)

type MockBuyerService struct {
	mock.Mock
}

func (m *MockBuyerService) Create(buyer *models.Buyer) (*models.Buyer, error) {
	args := m.Called(buyer)
	return args.Get(0).(*models.Buyer), args.Error(1)
}

func (m *MockBuyerService) GetById(id int) (*models.Buyer, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Buyer), args.Error(1)
}

func (m *MockBuyerService) GetAll() ([]*models.Buyer, error) {
	args := m.Called()
	return args.Get(0).([]*models.Buyer), args.Error(1)
}

func (m *MockBuyerService) Update(id int, buyer *models.Buyer) (*models.Buyer, error) {
	args := m.Called(id, buyer)
	return args.Get(0).(*models.Buyer), args.Error(1)
}

func (m *MockBuyerService) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockBuyerService) GetByIdWithOrderCount(id int) (*models.BuyerWithOrderCount, error) {
	args := m.Called(id)
	return args.Get(0).(*models.BuyerWithOrderCount), args.Error(1)
}

func (m *MockBuyerService) GetAllWithOrderCount() ([]*models.BuyerWithOrderCount, error) {
	args := m.Called()
	return args.Get(0).([]*models.BuyerWithOrderCount), args.Error(1)
}

func (m *MockBuyerService) PatchUpdate(id int, updateDTO *models.BuyerUpdateDTO) (*models.Buyer, error) {
	args := m.Called(id, updateDTO)
	return args.Get(0).(*models.Buyer), args.Error(1)
}
