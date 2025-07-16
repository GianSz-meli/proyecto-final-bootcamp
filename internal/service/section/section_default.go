package service

import (
	repository "ProyectoFinal/internal/repository/section"
	"ProyectoFinal/pkg/models"
)

// NewSectionDefault creates a new instance of SectionDefault with the provided repository
func NewSectionDefault(rp repository.SectionRepository) *SectionDefault {
	return &SectionDefault{rp: rp}
}

// SectionDefault implements SectionService interface and provides business logic for section operations
type SectionDefault struct {
	rp repository.SectionRepository
}

// GetAll retrieves all sections by delegating to the repository layer
func (s *SectionDefault) GetAll() (sections []models.Section, err error) {
	sections, err = s.rp.GetAll()
	if err != nil {
		return nil, err
	}
	return sections, nil
}

// GetById retrieves a section by its ID by delegating to the repository layer
// Propagates any errors from the repository (including not found errors)
func (s *SectionDefault) GetById(id int) (section models.Section, err error) {
	section, err = s.rp.GetById(id)
	if err != nil {
		return models.Section{}, err
	}
	return section, nil
}

// Create stores a new section by delegating to the repository layer
// Returns the created section with the generated ID
func (s *SectionDefault) Create(section models.Section) (createdSection models.Section, err error) {
	createdSection, err = s.rp.Create(section)
	if err != nil {
		return models.Section{}, err
	}
	return createdSection, nil
}

// Update modifies an existing section by ID by delegating to the repository layer
// First validates that the section exists, then performs the update
// Returns the updated section with the provided ID
func (s *SectionDefault) Update(id int, section models.Section) (updatedSection models.Section, err error) {
	_, err = s.rp.GetById(id)
	if err != nil {
		return models.Section{}, err
	}

	section.ID = id
	updatedSection, err = s.rp.Update(id, section)
	if err != nil {
		return models.Section{}, err
	}
	return updatedSection, nil
}

// Delete removes a section by ID by delegating to the repository layer
// First validates that the section exists, then performs the deletion
func (s *SectionDefault) Delete(id int) (err error) {
	_, err = s.rp.GetById(id)
	if err != nil {
		return err
	}

	err = s.rp.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
