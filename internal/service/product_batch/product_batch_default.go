package product_batch

import (
	productbatchrepo "ProyectoFinal/internal/repository/product_batch"
	"ProyectoFinal/pkg/errors"
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
		// Manejar errores específicos del repositorio
		if err.Error() == "batch_number already exists" {
			return models.ProductBatch{}, errors.WrapErrAlreadyExist("product batch", "batch_number", productBatch.BatchNumber)
		}
		if err.Error() == "product_id does not exist" {
			return models.ProductBatch{}, errors.WrapErrNotFound("product", "id", productBatch.ProductID)
		}
		if err.Error() == "section_id does not exist" {
			return models.ProductBatch{}, errors.WrapErrNotFound("section", "id", productBatch.SectionID)
		}
		return models.ProductBatch{}, err
	}

	return createdProductBatch, nil
}

func (s *productBatchService) GetProductCountBySection(sectionID *int) ([]models.SectionProductReport, error) {
	reports, err := s.repository.GetProductCountBySection(sectionID)
	if err != nil {
		return nil, err
	}

	// Si se especificó una sección y no se encontró ningún reporte, devolver 404
	if sectionID != nil && len(reports) == 0 {
		return nil, errors.WrapErrNotFound("section", "id", *sectionID)
	}

	return reports, nil
}
