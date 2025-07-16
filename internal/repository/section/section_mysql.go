package repository

import (
	pkgErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"database/sql"
	"errors"
	"log"
)

// SectionMySQL implements SectionRepository interface using MySQL database
type SectionMySQL struct {
	db *sql.DB
}

// NewSectionMySQL creates a new instance of SectionMySQL with the provided database connection
func NewSectionMySQL(db *sql.DB) *SectionMySQL {
	return &SectionMySQL{db: db}
}

// GetAll retrieves all sections from the MySQL database
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

// GetById retrieves a section by its ID from the MySQL database
// Returns a not found error if the section doesn't exist
func (r *SectionMySQL) GetById(id int) (models.Section, error) {
	var s models.Section
	err := r.db.QueryRow(SectionSelectById, id).Scan(
		&s.ID, &s.SectionNumber, &s.CurrentTemperature, &s.MinimumTemperature, &s.CurrentCapacity, &s.MinimumCapacity, &s.ProductTypeID, &s.WarehouseID, &s.MaximumCapacity,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			newError := pkgErrors.WrapErrNotFound("Section", "id", id)
			return s, newError
		}
		return s, err
	}
	return s, nil
}

// Create stores a new section in the MySQL database within a transaction
// Returns the created section with the generated ID
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

// Update modifies an existing section in the MySQL database within a transaction
// Returns the updated section with the provided ID
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

// Delete removes a section from the MySQL database within a transaction
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

// ExistBySectionNumber checks if a section exists with the given section number
// Returns true if the section exists, false otherwise
func (r *SectionMySQL) ExistBySectionNumber(sectionNumber string) bool {
	var exists bool
	err := r.db.QueryRow(SectionExistsByNumber, sectionNumber).Scan(&exists)
	return err == nil && exists
}
