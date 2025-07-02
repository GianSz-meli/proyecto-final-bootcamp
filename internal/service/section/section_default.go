package service

import (
	repository "ProyectoFinal/internal/repository/section"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
)

func NewSectionDefault(rp repository.SectionRepository) *SectionDefault {
	return &SectionDefault{rp: rp}
}

type SectionDefault struct {
	rp repository.SectionRepository
}

func (s *SectionDefault) GetAll() (sections []models.Section, err error) {
	sections, err = s.rp.GetAll()
	if err != nil {
		return nil, err
	}
	return sections, nil
}

func (s *SectionDefault) GetByID(id int) (section models.Section, err error) {
	section, err = s.rp.GetByID(id)
	if err != nil {
		return models.Section{}, errors.WrapErrNotFound("Section", "id", id)
	}
	return section, nil
}

func (s *SectionDefault) Create(section models.Section) (createdSection models.Section, err error) {
	exists := s.rp.ExistBySectionNumber(section.SectionNumber)
	if exists {
		return models.Section{}, errors.WrapErrAlreadyExist("Section", "section_number", section.SectionNumber)
	}

	createdSection, err = s.rp.Create(section)
	if err != nil {
		return models.Section{}, err
	}
	return createdSection, nil
}

func (s *SectionDefault) Update(id int, section models.Section) (updatedSection models.Section, err error) {
	existingSection, err := s.rp.GetByID(id)
	if err != nil {
		return models.Section{}, errors.WrapErrNotFound("Section", "id", id)
	}

	if section.SectionNumber != 0 && section.SectionNumber != existingSection.SectionNumber {
		exists := s.rp.ExistBySectionNumber(section.SectionNumber)
		if exists {
			return models.Section{}, errors.WrapErrAlreadyExist("Section", "section_number", section.SectionNumber)
		}
	}

	mergedSection := existingSection
	if section.SectionNumber != 0 {
		mergedSection.SectionNumber = section.SectionNumber
	}
	if section.CurrentTemperature != 0 {
		mergedSection.CurrentTemperature = section.CurrentTemperature
	}
	if section.MinimumTemperature != 0 {
		mergedSection.MinimumTemperature = section.MinimumTemperature
	}
	if section.CurrentCapacity != 0 {
		mergedSection.CurrentCapacity = section.CurrentCapacity
	}
	if section.MinimumCapacity != 0 {
		mergedSection.MinimumCapacity = section.MinimumCapacity
	}
	if section.MaximumCapacity != 0 {
		mergedSection.MaximumCapacity = section.MaximumCapacity
	}
	if section.WarehouseID != 0 {
		mergedSection.WarehouseID = section.WarehouseID
	}
	if section.ProductTypeID != 0 {
		mergedSection.ProductTypeID = section.ProductTypeID
	}
	if section.ProductBatches != nil {
		mergedSection.ProductBatches = section.ProductBatches
	}

	updatedSection, err = s.rp.Update(id, mergedSection)
	if err != nil {
		return models.Section{}, err
	}
	return updatedSection, nil
}

func (s *SectionDefault) Delete(id int) (err error) {
	_, err = s.rp.GetByID(id)
	if err != nil {
		return errors.WrapErrNotFound("Section", "id", id)
	}

	err = s.rp.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
