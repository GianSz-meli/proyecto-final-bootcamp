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

func (s *SectionDefault) GetById(id int) (section models.Section, err error) {
	section, exists := s.rp.GetById(id)
	if !exists {
		return models.Section{}, errors.WrapErrNotFound("section", "id", id)
	}
	return section, nil
}

func (s *SectionDefault) Create(section models.Section) (createdSection models.Section, err error) {
	createdSection, err = s.rp.Create(section)
	if err != nil {
		return models.Section{}, err
	}
	return createdSection, nil
}

func (s *SectionDefault) Update(id int, section models.Section) (updatedSection models.Section, err error) {
	_, exists := s.rp.GetById(id)
	if !exists {
		return models.Section{}, errors.WrapErrNotFound("section", "id", id)
	}

	section.ID = id
	updatedSection, err = s.rp.Update(id, section)
	if err != nil {
		return models.Section{}, err
	}
	return updatedSection, nil
}

func (s *SectionDefault) Delete(id int) (err error) {
	_, exists := s.rp.GetById(id)
	if !exists {
		return errors.WrapErrNotFound("section", "id", id)
	}

	err = s.rp.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
