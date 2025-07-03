package repository

import (
	"ProyectoFinal/internal/repository/utils"
	"ProyectoFinal/pkg/models"
)

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

type SectionMap struct {
	db        map[int]models.Section
	idCounter int
}

func (r *SectionMap) GetAll() (s []models.Section, err error) {
	s = make([]models.Section, 0, len(r.db))
	for _, value := range r.db {
		s = append(s, value)
	}
	return
}

func (r *SectionMap) GetById(id int) (models.Section, error) {
	section, exists := r.db[id]
	if !exists {
		return models.Section{}, nil
	}
	return section, nil
}

func (r *SectionMap) Create(section models.Section) (s models.Section, err error) {
	r.idCounter++
	section.ID = r.idCounter
	r.db[section.ID] = section
	return section, nil
}

func (r *SectionMap) Update(id int, section models.Section) (models.Section, error) {
	section.ID = id
	r.db[id] = section
	return section, nil
}

func (r *SectionMap) Delete(id int) error {
	delete(r.db, id)
	return nil
}

func (r *SectionMap) ExistBySectionNumber(sectionNumber int) bool {
	for _, section := range r.db {
		if section.SectionNumber == sectionNumber {
			return true
		}
	}
	return false
}
