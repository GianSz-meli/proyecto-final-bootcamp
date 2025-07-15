package repository

import "ProyectoFinal/pkg/models"

type ProductBatchRepository interface {
	Create(productBatch models.ProductBatch) (models.ProductBatch, error)
	GetProductCountBySection(sectionID *int) ([]models.SectionProductReport, error)
}
