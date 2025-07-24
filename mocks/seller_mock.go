package mocks

import "ProyectoFinal/pkg/models"

type MockSellerService struct {
	CreateFunc  func(seller models.Seller) (models.Seller, error)
	GetAllFunc  func() ([]models.Seller, error)
	GetByIdFunc func(id int) (models.Seller, error)
	DeleteFunc  func(id int) error
	UpdateFunc  func(seller models.Seller) (models.Seller, error)
}

func (m *MockSellerService) Create(seller models.Seller) (models.Seller, error) {
	return m.CreateFunc(seller)
}

func (m *MockSellerService) GetAll() ([]models.Seller, error) {
	return m.GetAllFunc()
}

func (m *MockSellerService) GetById(id int) (models.Seller, error) {
	return m.GetByIdFunc(id)
}

func (m *MockSellerService) Delete(id int) error {
	return m.DeleteFunc(id)
}

func (m *MockSellerService) Update(seller models.Seller) (models.Seller, error) {
	return m.UpdateFunc(seller)
}
