package repository

import (
	"ProyectoFinal/pkg/models"
	"database/sql"
	"errors"
	"log"
)

type SectionMySQL struct {
	db *sql.DB
}

func NewSectionMySQL(db *sql.DB) *SectionMySQL {
	return &SectionMySQL{db: db}
}

func (r *SectionMySQL) GetAll() ([]models.Section, error) {
	rows, err := r.db.Query(SectionSelectAll)
	if err != nil {
		log.Println("[SectionMySQL][GetAll] error en Query:", err)
		return nil, err
	}
	defer rows.Close()

	var sections []models.Section
	for rows.Next() {
		var s models.Section
		err := rows.Scan(&s.ID, &s.SectionNumber, &s.CurrentTemperature, &s.MinimumTemperature, &s.CurrentCapacity, &s.MinimumCapacity, &s.ProductTypeID, &s.WarehouseID, &s.MaximumCapacity)
		if err != nil {
			log.Println("[SectionMySQL][GetAll] error en Scan:", err)
			return nil, err
		}
		sections = append(sections, s)
	}
	return sections, nil
}

func (r *SectionMySQL) GetById(id int) (models.Section, bool) {
	var s models.Section
	err := r.db.QueryRow(SectionSelectById, id).Scan(
		&s.ID, &s.SectionNumber, &s.CurrentTemperature, &s.MinimumTemperature, &s.CurrentCapacity, &s.MinimumCapacity, &s.ProductTypeID, &s.WarehouseID, &s.MaximumCapacity,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return s, false
		}
		return s, false
	}
	return s, true
}

func (r *SectionMySQL) Create(section models.Section) (models.Section, error) {
	tx, err := r.db.Begin()
	if err != nil {
		log.Println("[SectionMySQL][Create] error en Begin:", err)
		return section, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	res, err := tx.Exec(SectionInsert,
		section.SectionNumber, section.CurrentTemperature, section.MinimumTemperature, section.CurrentCapacity, section.MinimumCapacity, section.ProductTypeID, section.WarehouseID, section.MaximumCapacity)
	if err != nil {
		log.Println("[SectionMySQL][Create] error en Exec:", err)
		tx.Rollback()
		return section, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Println("[SectionMySQL][Create] error en LastInsertId:", err)
		tx.Rollback()
		return section, err
	}
	section.ID = int(id)
	if err := tx.Commit(); err != nil {
		log.Println("[SectionMySQL][Create] error en Commit:", err)
		return section, err
	}
	return section, nil
}

func (r *SectionMySQL) Update(id int, section models.Section) (models.Section, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return section, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	_, err = tx.Exec(SectionUpdate,
		section.SectionNumber, section.CurrentTemperature, section.MinimumTemperature, section.CurrentCapacity, section.MinimumCapacity, section.ProductTypeID, section.WarehouseID, section.MaximumCapacity, id)
	if err != nil {
		tx.Rollback()
		return section, err
	}
	if err := tx.Commit(); err != nil {
		return section, err
	}
	section.ID = id
	return section, nil
}

func (r *SectionMySQL) Delete(id int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	_, err = tx.Exec(SectionDelete, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (r *SectionMySQL) ExistBySectionNumber(sectionNumber int) bool {
	var exists bool
	err := r.db.QueryRow(SectionExistsByNumber, sectionNumber).Scan(&exists)
	return err == nil && exists
}
