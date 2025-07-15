package repository

import (
	"ProyectoFinal/pkg/models"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type ProductBatchMySQL struct {
	db *sql.DB
}

func NewProductBatchMySQL(db *sql.DB) *ProductBatchMySQL {
	return &ProductBatchMySQL{db: db}
}

func (r *ProductBatchMySQL) Create(productBatch models.ProductBatch) (models.ProductBatch, error) {
	// Verificar que batch_number no exista
	if r.ExistsByBatchNumber(productBatch.BatchNumber) {
		return productBatch, errors.New("batch_number already exists")
	}

	// Verificar que product_id existe
	if !r.ProductExists(productBatch.ProductID) {
		return productBatch, errors.New("product_id does not exist")
	}

	// Verificar que section_id existe
	if !r.SectionExists(productBatch.SectionID) {
		return productBatch, errors.New("section_id does not exist")
	}

	tx, err := r.db.Begin()
	if err != nil {
		log.Println("[ProductBatchMySQL][Create] error en Begin:", err)
		return productBatch, err
	}
	defer func() {
		if err != nil {
			log.Println("[ProductBatchMySQL][Create] rolling back transaction due to error:", err)
			tx.Rollback()
		}
	}()

	// Convertir manufacturing_hour de int a formato TIME (HH:MM:SS)
	manufacturingHourStr := fmt.Sprintf("%02d:00:00", productBatch.ManufacturingHour)

	res, err := tx.Exec(ProductBatchInsert,
		productBatch.BatchNumber, productBatch.CurrentQuantity, productBatch.CurrentTemperature,
		productBatch.DueDate, productBatch.InitialQuantity, productBatch.ManufacturingDate,
		manufacturingHourStr, productBatch.MinimumTemperature, productBatch.ProductID, productBatch.SectionID)
	if err != nil {
		log.Println("[ProductBatchMySQL][Create] error en Exec:", err)
		tx.Rollback()
		return productBatch, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println("[ProductBatchMySQL][Create] error en LastInsertId:", err)
		tx.Rollback()
		return productBatch, err
	}

	productBatch.ID = int(id)
	if err := tx.Commit(); err != nil {
		log.Println("[ProductBatchMySQL][Create] error en Commit:", err)
		return productBatch, err
	}

	return productBatch, nil
}

// Validation methods
func (r *ProductBatchMySQL) ExistsByBatchNumber(batchNumber string) bool {
	var exists bool
	err := r.db.QueryRow(ProductBatchExistsByNumber, batchNumber).Scan(&exists)
	if err != nil {
		log.Println("[ProductBatchMySQL][ExistsByBatchNumber] error:", err)
		return false
	}
	return exists
}

func (r *ProductBatchMySQL) ProductExists(productID int) bool {
	var exists bool
	err := r.db.QueryRow(ProductExistsById, productID).Scan(&exists)
	if err != nil {
		log.Println("[ProductBatchMySQL][ProductExists] error:", err)
		return false
	}
	return exists
}

func (r *ProductBatchMySQL) SectionExists(sectionID int) bool {
	var exists bool
	err := r.db.QueryRow(SectionExistsById, sectionID).Scan(&exists)
	if err != nil {
		log.Println("[ProductBatchMySQL][SectionExists] error:", err)
		return false
	}
	return exists
}

// Report methods
func (r *ProductBatchMySQL) GetProductCountBySection(sectionID *int) ([]models.SectionProductReport, error) {
	var rows *sql.Rows
	var err error

	if sectionID != nil {
		rows, err = r.db.Query(ProductBatchCountBySectionId, *sectionID)
	} else {
		rows, err = r.db.Query(ProductBatchCountBySection)
	}

	if err != nil {
		log.Println("[ProductBatchMySQL][GetProductCountBySection] error en Query:", err)
		return nil, err
	}
	defer rows.Close()

	var reports []models.SectionProductReport
	for rows.Next() {
		var report models.SectionProductReport
		err := rows.Scan(&report.SectionID, &report.SectionNumber, &report.ProductsCount)
		if err != nil {
			log.Println("[ProductBatchMySQL][GetProductCountBySection] error en Scan:", err)
			return nil, err
		}
		reports = append(reports, report)
	}

	return reports, nil
}
