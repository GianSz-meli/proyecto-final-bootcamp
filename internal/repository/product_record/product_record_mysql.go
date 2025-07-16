package repository

import (
	pkgErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
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

func (r *ProductRecordSQL) GetRecordsProduct(prodID int) (models.ReportProductData, error) {
	var res models.ReportProductData

	err := r.db.QueryRow(QueryReport, prodID).Scan(&res.ProductID, &res.Description, &res.RecordsCount)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			newError := pkgErrors.WrapErrNotFound("product", "id", prodID)
			return models.ReportProductData{}, newError
		}
		return models.ReportProductData{}, err
	}
	return res, nil
}

func (r *ProductRecordSQL) GetRecordsProductAll() ([]models.ReportProductData, error) {

	rows, err := r.db.Query(QueryGetAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.ReportProductData
	for rows.Next() {
		var res models.ReportProductData
		if err := rows.Scan(&res.ProductID, &res.Description, &res.RecordsCount); err != nil {
			return nil, err
		}
		results = append(results, res)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
