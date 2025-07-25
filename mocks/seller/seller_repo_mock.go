package seller

import (
	"ProyectoFinal/pkg/models"
	"github.com/stretchr/testify/mock"
)

type MockSellerRepository struct {
	mock.Mock
}

func (m *MockSellerRepository) Create(seller models.Seller) (models.Seller, error) {
	args := m.Called(seller)
	return args.Get(0).(models.Seller), args.Error(1)
}
func (m *MockSellerRepository) GetById(id int) (*models.Seller, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Seller), args.Error(1)
}
func (m *MockSellerRepository) Update(seller *models.Seller) (models.Seller, error) {
	args := m.Called(seller)
	return args.Get(0).(models.Seller), args.Error(1)
}
func (m *MockSellerRepository) GetAll() ([]models.Seller, error) {
	args := m.Called()
	return args.Get(0).([]models.Seller), args.Error(1)
}
func (m *MockSellerRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(1)
}
