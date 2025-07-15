package product_batch

import (
	productbatchrepo "ProyectoFinal/internal/repository/product_batch"
	"ProyectoFinal/pkg/models"
)

type productBatchService struct {
	repository productbatchrepo.ProductBatchRepository
}

func NewProductBatchService(repository productbatchrepo.ProductBatchRepository) ProductBatchService {
	return &productBatchService{
		repository: repository,
	}
}

func (s *productBatchService) Create(productBatch models.ProductBatch) (models.ProductBatch, error) {
	createdProductBatch, err := s.repository.Create(productBatch)
	if err != nil {
		return models.ProductBatch{}, err
	}

	return createdProductBatch, nil
}

func (s *productBatchService) GetProductCountBySection(sectionID *int) ([]models.SectionProductReport, error) {
	reports, err := s.repository.GetProductCountBySection(sectionID)
	if err != nil {
		return nil, err
	}

	return reports, nil
}
