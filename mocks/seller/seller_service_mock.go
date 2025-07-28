package seller

import (
	"ProyectoFinal/pkg/models"
	"github.com/stretchr/testify/mock"
)

// MockSellerService - implementation of seller interface
type MockSellerService struct {
	mock.Mock
}

func (m *MockSellerService) Create(seller models.Seller) (models.Seller, error) {
	args := m.Called(seller)
	return args.Get(0).(models.Seller), args.Error(1)
}

func (m *MockSellerService) GetAll() ([]models.Seller, error) {
	args := m.Called()
	return args.Get(0).([]models.Seller), args.Error(1)
}

func (m *MockSellerService) GetById(id int) (models.Seller, error) {
	args := m.Called(id)
	return args.Get(0).(models.Seller), args.Error(1)
}

func (m *MockSellerService) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockSellerService) Update(id int, reqBody *models.UpdateSellerRequest) (models.Seller, error) {
	args := m.Called(id, reqBody)
	return args.Get(0).(models.Seller), args.Error(1)
}
