package repository

import (
	"ProyectoFinal/pkg/models"
	"database/sql"
	
)

func NewProductRecordSQL(db *sql.DB) *ProductRecordSQL {
	return &ProductRecordSQL{db: db}
}

type ProductRecordSQL struct {
	db *sql.DB
}

func (r *ProductRecordSQL) ExistsProductRecordID(productID int) (bool, error) {
    var exists bool
    query := "SELECT EXISTS (SELECT 1 FROM products WHERE id = ?)"
    err := r.db.QueryRow(query, productID).Scan(&exists)
    if err != nil {
        return false, err
    }
    return exists, nil
}

func (r *ProductRecordSQL) CreateProductRecord(newRecord models.ProductRecord) (models.ProductRecord, error) {
	query := `
        INSERT INTO product_records 
            (last_update_date, purchase_price, sale_price, product_id)
        VALUES (?, ?, ?, ?)
    `
    res, err := r.db.Exec(
        query,
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