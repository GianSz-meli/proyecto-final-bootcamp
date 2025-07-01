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

func (s *SectionDefault) GetAll() (sections map[int]models.Section, err error) {
	sections, err = s.rp.GetAll()
	if err != nil {
		return nil, err
	}
	return sections, nil
}

func (s *SectionDefault) GetByID(id int) (section models.Section, err error) {
	if id <= 0 {
		return models.Section{}, errors.ErrInvalidSectionID
	}

	section, err = s.rp.GetByID(id)
	if err != nil {
		return models.Section{}, err
	}
	return section, nil
}

func (s *SectionDefault) Create(section models.Section) (createdSection models.Section, err error) {
	if err := s.validateSection(section); err != nil {
		return models.Section{}, errors.ErrInvalidSectionData
	}

	sections, err := s.rp.GetAll()
	if err != nil {
		return models.Section{}, err
	}

	for _, existingSection := range sections {
		if existingSection.SectionNumber == section.SectionNumber {
			return models.Section{}, errors.ErrSectionNumberExists
		}
	}

	createdSection, err = s.rp.Create(section)
	if err != nil {
		return models.Section{}, err
	}
	return createdSection, nil
}

func (s *SectionDefault) Update(id int, section models.Section) (updatedSection models.Section, err error) {
	if id <= 0 {
		return models.Section{}, errors.ErrInvalidSectionID
	}

	_, err = s.rp.GetByID(id)
	if err != nil {
		return models.Section{}, err
	}

	if section.SectionNumber != 0 {
		sections, err := s.rp.GetAll()
		if err != nil {
			return models.Section{}, err
		}

		for _, sec := range sections {
			if sec.ID != id && sec.SectionNumber == section.SectionNumber {
				return models.Section{}, errors.ErrSectionNumberExists
			}
		}
	}

	if err := s.validateSection(section); err != nil {
		return models.Section{}, errors.ErrInvalidSectionData
	}

	updatedSection, err = s.rp.Update(id, section)
	if err != nil {
		return models.Section{}, err
	}
	return updatedSection, nil
}

func (s *SectionDefault) Delete(id int) (err error) {
	if id <= 0 {
		return errors.ErrInvalidSectionID
	}

	_, err = s.rp.GetByID(id)
	if err != nil {
		return err
	}

	err = s.rp.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SectionDefault) validateSection(section models.Section) error {
	if section.SectionNumber <= 0 {
		return errors.ErrInvalidSectionData
	}

	if section.CurrentTemperature < -50 || section.CurrentTemperature > 50 {
		return errors.ErrInvalidSectionData
	}

	if section.MinimumTemperature < -50 || section.MinimumTemperature > 50 {
		return errors.ErrInvalidSectionData
	}

	if section.CurrentTemperature < section.MinimumTemperature {
		return errors.ErrInvalidSectionData
	}

	if section.CurrentCapacity < 0 {
		return errors.ErrInvalidSectionData
	}

	if section.MinimumCapacity < 0 {
		return errors.ErrInvalidSectionData
	}

	if section.MaximumCapacity <= 0 {
		return errors.ErrInvalidSectionData
	}

	if section.MinimumCapacity > section.MaximumCapacity {
		return errors.ErrInvalidSectionData
	}

	if section.CurrentCapacity > section.MaximumCapacity {
		return errors.ErrInvalidSectionData
	}

	if section.CurrentCapacity < section.MinimumCapacity {
		return errors.ErrInvalidSectionData
	}

	if section.WarehouseID <= 0 {
		return errors.ErrInvalidSectionData
	}

	if section.ProductTypeID <= 0 {
		return errors.ErrInvalidSectionData
	}

	return nil
}
