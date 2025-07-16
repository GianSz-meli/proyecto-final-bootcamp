package service

import "ProyectoFinal/pkg/models"

type ProductRecordService interface {
	CreateProductRecord(newProd models.ProductRecord) (models.ProductRecord, error)
	GetRecordsProduct(prodID *int) (models.ReportProductData, error)
	GetRecordsProductAll() ([]models.ReportProductData, error) 
}
