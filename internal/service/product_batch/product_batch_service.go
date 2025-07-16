package product_batch

import "ProyectoFinal/pkg/models"

// ProductBatchService defines the contract for product batch business logic operations
type ProductBatchService interface {
	Create(productBatch models.ProductBatch) (models.ProductBatch, error)
	GetProductCountBySection(sectionID *int) ([]models.SectionProductReport, error)
}
