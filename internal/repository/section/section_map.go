package repository

import (
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
)

func NewSectionMap(db map[int]models.Section) *SectionMap {
	defaultDb := make(map[int]models.Section)
	if db != nil {
		defaultDb = db
	}

	return &SectionMap{db: defaultDb, lastID: len(defaultDb)}
}

type SectionMap struct {
	db     map[int]models.Section
	lastID int
}

func (r *SectionMap) GetAll() (s map[int]models.Section, err error) {
	s = make(map[int]models.Section)
	for key, value := range r.db {
		s[key] = value
	}
	return
}

func (r *SectionMap) GetByID(id int) (models.Section, error) {
	var s, exist = r.db[id]
	if !exist {
		return models.Section{}, errors.ErrNotFound
	}
	return s, nil
}

func (r *SectionMap) Create(section models.Section) (s models.Section, err error) {
	r.lastID++
	section.ID = r.lastID
	r.db[section.ID] = section
	return
}

func (r *SectionMap) Update(id int, section models.Section) (models.Section, error) {
	_, exist := r.db[id]
	if !exist {
		return models.Section{}, errors.ErrNotFound
	}

	r.db[id] = section
	return section, nil
}

func (r *SectionMap) Delete(id int) error {
	_, exist := r.db[id]
	if !exist {
		return errors.ErrNotFound
	}

	delete(r.db, id)
	return nil
}

func (r *SectionMap) ExistBySectionNumber(sectionNumber int) (bool, error) {
	for _, section := range r.db {
		if section.SectionNumber == sectionNumber {
			return true, nil
		}
	}
	return false, nil
}
