package product_batch

import "ProyectoFinal/pkg/models"

type ProductBatchService interface {
	Create(productBatch models.ProductBatch) (models.ProductBatch, error)
	GetProductCountBySection(sectionID *int) ([]models.SectionProductReport, error)
}
