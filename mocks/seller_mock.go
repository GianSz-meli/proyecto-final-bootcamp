package mocks

import "ProyectoFinal/pkg/models"

type MockSellerService struct {
	CreateFunc  func(seller models.Seller) (models.Seller, error)
	GetAllFunc  func() ([]models.Seller, error)
	GetByIdFunc func(id int) (models.Seller, error)
	DeleteFunc  func(id int) error
	UpdateFunc  func(seller models.Seller) (models.Seller, error)
	Spy         struct {
		CountCreateFunc  int
		CountGetAllFunc  int
		CountGetByIdFunc int
		CountDeleteFunc  int
		CountUpdateFunc  int
	}
}

func (m *MockSellerService) Create(seller models.Seller) (models.Seller, error) {
	m.Spy.CountCreateFunc++
	return m.CreateFunc(seller)
}

func (m *MockSellerService) GetAll() ([]models.Seller, error) {
	m.Spy.CountGetAllFunc++
	return m.GetAllFunc()
}

func (m *MockSellerService) GetById(id int) (models.Seller, error) {
	m.Spy.CountGetByIdFunc++
	return m.GetByIdFunc(id)
}

func (m *MockSellerService) Delete(id int) error {
	m.Spy.CountDeleteFunc++
	return m.DeleteFunc(id)
}

func (m *MockSellerService) Update(seller models.Seller) (models.Seller, error) {
	m.Spy.CountUpdateFunc++
	return m.UpdateFunc(seller)
}
