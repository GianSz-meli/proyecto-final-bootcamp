package service

import (
	repository "ProyectoFinal/internal/repository/product_record"
	"ProyectoFinal/pkg/models"
)

func NewProductRecordDefault(rp repository.ProductRecordRepository) *ProductRecordDefault {
	return &ProductRecordDefault{rp: rp}
}

type ProductRecordDefault struct {
	// rp is the repository that will be used by the service
	rp repository.ProductRecordRepository
}

func (s *ProductRecordDefault) CreateProductRecord(prod models.ProductRecord) (models.ProductRecord, error) {
	newProd, err := s.rp.CreateProductRecord(prod)
	if err != nil {
		return models.ProductRecord{}, err
	}
	return newProd, nil
}

func (s *ProductRecordDefault) GetRecordsProduct(prodID *int) (models.ReportProductData, error) {
	data, err := s.rp.GetRecordsProduct(*prodID)
	if err != nil {
		return models.ReportProductData{}, err
	}
	return data, nil
}

func (s *ProductRecordDefault) GetRecordsProductAll() ([]models.ReportProductData, error) {
	products, err := s.rp.GetRecordsProductAll()
	if err != nil {
		return []models.ReportProductData{}, err
	}
	return products, nil 
}
