package repository

import (
	"ProyectoFinal/pkg/models"
    pkgErrors "ProyectoFinal/pkg/errors"
	"database/sql"
	"errors"
)

func NewProductRecordSQL(db *sql.DB) *ProductRecordSQL {
	return &ProductRecordSQL{db: db}
}

type ProductRecordSQL struct {
	db *sql.DB
}

func (r *ProductRecordSQL) ExistsProductRecordID(productID int) (bool, error) {
    var exists bool
    err := r.db.QueryRow(QueryExists, productID).Scan(&exists)
    if err != nil {
        return false, err
    }
    return exists, nil
}

func (r *ProductRecordSQL) CreateProductRecord(newRecord models.ProductRecord) (models.ProductRecord, error) {
	
    res, err := r.db.Exec(
        QueryCreate,
        newRecord.LastUpdateDate, 
        newRecord.PurchasePrice,
        newRecord.SalePrice,
        newRecord.ProductID,
    )
    if err != nil {
        return newRecord, err
    }
    id, err := res.LastInsertId()
    if err != nil {
        return newRecord, err
    }
    newRecord.ID = int(id)
    return newRecord, nil
}

func (r *ProductRecordSQL) GetRecordsProduct(prodID int) (models.ReportProductData, error){
    var res models.ReportProductData

    err := r.db.QueryRow(QueryReport, prodID).Scan(&res.ProductID, &res.Description, &res.RecordsCount)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
		newError := pkgErrors.WrapErrNotFound("product record", "productID", prodID)
			return models.ReportProductData{}, newError
		}
        return models.ReportProductData{}, err
    }
    return res, nil
}


