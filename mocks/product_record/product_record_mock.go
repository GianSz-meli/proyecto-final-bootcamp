package mocks

import (
	"ProyectoFinal/pkg/models"

	"github.com/stretchr/testify/mock"
)

type MockProductRecordRepository struct {
	mock.Mock
}

func (m *MockProductRecordRepository) ExistsProductRecordID(productID int) (bool, error) {
	args := m.Called(productID)
	return args.Bool(0), args.Error(1)
}

func (m *MockProductRecordRepository) CreateProductRecord(newRecord models.ProductRecord) (models.ProductRecord, error) {
	args := m.Called(newRecord)
	return args.Get(0).(models.ProductRecord), args.Error(1)
}

func (m *MockProductRecordRepository) GetRecordsProduct(prodID int) (models.ReportProductData, error) {
	args := m.Called(prodID)
	return args.Get(0).(models.ReportProductData), args.Error(1)
}

func (m *MockProductRecordRepository) GetRecordsProductAll() ([]models.ReportProductData, error) {
	args := m.Called()
	return args.Get(0).([]models.ReportProductData), args.Error(1)
}

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
