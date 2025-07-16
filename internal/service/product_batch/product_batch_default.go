package product_batch

import (
	productbatchrepo "ProyectoFinal/internal/repository/product_batch"
	"ProyectoFinal/pkg/models"
)

// productBatchService implements ProductBatchService interface and provides business logic for product batch operations
type productBatchService struct {
	repository productbatchrepo.ProductBatchRepository
}

// NewProductBatchService creates a new instance of productBatchService with the provided repository
func NewProductBatchService(repository productbatchrepo.ProductBatchRepository) ProductBatchService {
	return &productBatchService{
		repository: repository,
	}
}

// Create stores a new product batch by delegating to the repository layer
// Returns the created product batch with the generated ID
func (s *productBatchService) Create(productBatch models.ProductBatch) (models.ProductBatch, error) {
	createdProductBatch, err := s.repository.Create(productBatch)
	if err != nil {
		return models.ProductBatch{}, err
	}

	return createdProductBatch, nil
}

// GetProductCountBySection retrieves a report of product counts by section by delegating to the repository layer
// If sectionID is provided, returns data for that specific section only
// If sectionID is nil, returns data for all sections
func (s *productBatchService) GetProductCountBySection(sectionID *int) ([]models.SectionProductReport, error) {
	reports, err := s.repository.GetProductCountBySection(sectionID)
	if err != nil {
		return nil, err
	}

	return reports, nil
}
