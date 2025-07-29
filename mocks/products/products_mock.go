package mocks

import (
	"ProyectoFinal/pkg/models"

	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) ExistsProdCode(prodCode string) bool {
	args := m.Called(prodCode)
	return args.Bool(0)
}

func (m *MockProductRepository) CreateProduct(newProd models.Product) (models.Product, error) {
	args := m.Called(newProd)
	return args.Get(0).(models.Product), args.Error(1)
}

func (m *MockProductRepository) FindAllProducts() (map[int]models.Product, error) {
	args := m.Called()
	return args.Get(0).(map[int]models.Product), args.Error(1)
}

func (m *MockProductRepository) FindProductsById(id int) (models.Product, error) {
	args := m.Called(id)
	return args.Get(0).(models.Product), args.Error(1)
}

func (m *MockProductRepository) UpdateProduct(id int, prod models.Product) (models.Product, error) {
	args := m.Called(id, prod)
	return args.Get(0).(models.Product), args.Error(1)
}

func (m *MockProductRepository) DeleteProduct(id int) {
	m.Called(id)
}

type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) CreateProduct(p models.Product) (models.Product, error) {
	args := m.Called(p)
	return args.Get(0).(models.Product), args.Error(1)
}

func (m *MockProductService) FindAllProducts() (map[int]models.Product, error) {
	args := m.Called()
	return args.Get(0).(map[int]models.Product), args.Error(1)
}

func (m *MockProductService) FindProductsById(id int) (models.Product, error) {
	args := m.Called(id)
	return args.Get(0).(models.Product), args.Error(1)
}

func (m *MockProductService) UpdateProduct(id int, p models.ProductDocUpdate) (models.Product, error) {
	args := m.Called(id, p)
	return args.Get(0).(models.Product), args.Error(1)
}

func (m *MockProductService) DeleteProduct(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
