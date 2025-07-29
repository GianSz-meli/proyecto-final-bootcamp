package mocks

import (
	"ProyectoFinal/pkg/models"

	"github.com/stretchr/testify/mock"
)

type MockProductRecordService struct {
	mock.Mock
}

func (m *MockProductRecordService) CreateProductRecord(newProd models.ProductRecord) (models.ProductRecord, error) {
	args := m.Called(newProd)
	return args.Get(0).(models.ProductRecord), args.Error(1)
}

func (m *MockProductRecordService) GetRecordsProduct(prodID *int) (models.ReportProductData, error) {
	args := m.Called(prodID)
	return args.Get(0).(models.ReportProductData), args.Error(1)
}

func (m *MockProductRecordService) GetRecordsProductAll() ([]models.ReportProductData, error) {
	args := m.Called()
	return args.Get(0).([]models.ReportProductData), args.Error(1)
}
