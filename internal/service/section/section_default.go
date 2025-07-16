package service

import (
	repository "ProyectoFinal/internal/repository/section"
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
	section, err = s.rp.GetById(id)
	if err != nil {
		return models.Section{}, err
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
