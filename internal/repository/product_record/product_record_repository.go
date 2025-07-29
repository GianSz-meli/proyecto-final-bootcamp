package repository

import (
	"ProyectoFinal/pkg/models"
)

type ProductRecordRepository interface {
	CreateProductRecord(newRecord models.ProductRecord) (models.ProductRecord, error)
	GetRecordsProduct(prodID int) (models.ReportProductData, error)
	GetRecordsProductAll() ([]models.ReportProductData, error)
}
