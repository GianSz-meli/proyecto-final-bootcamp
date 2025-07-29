package product_batch

import (
	"ProyectoFinal/pkg/models"

	"github.com/stretchr/testify/mock"
)

// MockProductBatchService - implementation of product batch service interface
type MockProductBatchService struct {
	mock.Mock
}

func (m *MockProductBatchService) Create(productBatch models.ProductBatch) (models.ProductBatch, error) {
	args := m.Called(productBatch)
	return args.Get(0).(models.ProductBatch), args.Error(1)
}

func (m *MockProductBatchService) GetProductCountBySection(sectionID *int) ([]models.SectionProductReport, error) {
	args := m.Called(sectionID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.SectionProductReport), args.Error(1)
}
