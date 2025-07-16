package repository

import (
	"ProyectoFinal/internal/repository/utils"
	pkgErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
)

// NewSectionMap creates a new instance of SectionMap with the provided data
// If no data is provided, it initializes with an empty map
func NewSectionMap(db map[int]models.Section) *SectionMap {
	defaultDb := make(map[int]models.Section)
	if db != nil {
		defaultDb = db
	}

	return &SectionMap{
		db:        defaultDb,
		idCounter: utils.GetLastId[models.Section](defaultDb),
	}
}

// SectionMap implements SectionRepository interface using in-memory map storage
type SectionMap struct {
	db        map[int]models.Section
	idCounter int
}

// GetAll retrieves all sections from the in-memory map
func (r *SectionMap) GetAll() (s []models.Section, err error) {
	s = make([]models.Section, 0, len(r.db))
	for _, value := range r.db {
		s = append(s, value)
	}
	return
}

// GetById retrieves a section by its ID from the in-memory map
// Returns a not found error if the section doesn't exist
func (r *SectionMap) GetById(id int) (models.Section, error) {
	section, exists := r.db[id]
	if !exists {
		return models.Section{}, pkgErrors.WrapErrNotFound("Section", "id", id)
	}
	return section, nil
}

// Create stores a new section in the in-memory map with an auto-generated ID
// Returns the created section with the generated ID
func (r *SectionMap) Create(section models.Section) (s models.Section, err error) {
	r.idCounter++
	section.ID = r.idCounter
	r.db[section.ID] = section
	return section, nil
}

// Update modifies an existing section in the in-memory map by ID
// Returns the updated section with the provided ID
func (r *SectionMap) Update(id int, section models.Section) (models.Section, error) {
	section.ID = id
	r.db[id] = section
	return section, nil
}

// Delete removes a section from the in-memory map by its ID
func (r *SectionMap) Delete(id int) error {
	delete(r.db, id)
	return nil
}

// ExistBySectionNumber checks if a section exists with the given section number
// Returns true if the section exists, false otherwise
func (r *SectionMap) ExistBySectionNumber(sectionNumber string) bool {
	for _, section := range r.db {
		if section.SectionNumber == sectionNumber {
			return true
		}
	}
	return false
}
