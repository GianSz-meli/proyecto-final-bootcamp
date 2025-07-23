package handler_test

import (
	"ProyectoFinal/pkg/models"
)


type MockProductService struct {
	CreateProductFunc      func(models.Product) (models.Product, error)
	FindAllProductsFunc    func() (map[int]models.Product, error)
	FindProductsByIdFunc   func(int) (models.Product, error)
	UpdateProductFunc      func(int, models.Product) (models.Product, error)
	DeleteProductFunc      func(int) error
}


func (m *MockProductService) CreateProduct(p models.Product) (models.Product, error) {
	if m.CreateProductFunc != nil {
		return m.CreateProductFunc(p)
	}
	return models.Product{}, nil
}
func (m *MockProductService) FindAllProducts() (map[int]models.Product, error) {
	if m.FindAllProductsFunc != nil {
		return m.FindAllProductsFunc()
	}
	return nil, nil
}
func (m *MockProductService) FindProductsById(id int) (models.Product, error) {
	if m.FindProductsByIdFunc != nil {
		return m.FindProductsByIdFunc(id)
	}
	return models.Product{}, nil
}
func (m *MockProductService) UpdateProduct(id int, p models.Product) (models.Product, error) {
	if m.UpdateProductFunc != nil {
		return m.UpdateProductFunc(id, p)
	}
	return models.Product{}, nil
}
func (m *MockProductService) DeleteProduct(id int) error {
	if m.DeleteProductFunc != nil {
		return m.DeleteProductFunc(id)
	}
	return nil
}




