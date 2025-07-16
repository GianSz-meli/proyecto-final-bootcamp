package repository

import "ProyectoFinal/pkg/models"

// ProductBatchRepository defines the contract for product batch data operations
type ProductBatchRepository interface {
	Create(productBatch models.ProductBatch) (models.ProductBatch, error)
	GetProductCountBySection(sectionID *int) ([]models.SectionProductReport, error)
}
