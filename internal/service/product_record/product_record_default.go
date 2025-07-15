package service

import (
	repository "ProyectoFinal/internal/repository/product_record"
	pkgErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	
)

func NewProductRecordDefault(rp repository.ProductRecordRepository) *ProductRecordDefault {
	return &ProductRecordDefault{rp: rp}
}

type ProductRecordDefault struct {
	// rp is the repository that will be used by the service
	rp repository.ProductRecordRepository
}

func (s *ProductRecordDefault) CreateProduct(newProd models.ProductRecord) (models.ProductRecord, error) {
	   exists, _ := s.rp.ExistsProductRecordID(newProd.ProductID)
    if exists {
        newError := pkgErrors.WrapErrAlreadyExist("product", "productID", newProd.ProductID)
        return models.ProductRecord{}, newError
    }
	prodReturn, err := s.rp.CreateProductRecord(newProd)
	if err != nil {
		return models.ProductRecord{}, err
	}
	return prodReturn, nil
}