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
		return models.Section{}, errors.ErrSectionNotFound
	}
	return s, nil
}

func (r *SectionMap) Create(section models.Section) (s models.Section, err error) {
	for _, value := range r.db {
		if value.SectionNumber == section.SectionNumber {
			return models.Section{}, errors.ErrSectionNumberExists
		}
	}
	r.lastID++
	section.ID = r.lastID
	r.db[section.ID] = section
	return
}

func (r *SectionMap) Update(id int, section models.Section) (models.Section, error) {
	var s, exist = r.db[id]
	if !exist {
		return models.Section{}, errors.ErrSectionNotFound
	}

	if section.SectionNumber != 0 {
		s.SectionNumber = section.SectionNumber
	}

	if section.CurrentTemperature != 0 {
		s.CurrentTemperature = section.CurrentTemperature
	}

	if section.MinimumTemperature != 0 {
		s.MinimumTemperature = section.MinimumTemperature
	}

	if section.CurrentCapacity != 0 {
		s.CurrentCapacity = section.CurrentCapacity
	}

	if section.MinimumCapacity != 0 {
		s.MinimumCapacity = section.MinimumCapacity
	}

	if section.MaximumCapacity != 0 {
		s.MaximumCapacity = section.MaximumCapacity
	}

	if section.WarehouseID != 0 {
		s.WarehouseID = section.WarehouseID
	}

	if section.ProductTypeID != 0 {
		s.ProductTypeID = section.ProductTypeID
	}

	if section.ProductBatches != nil {
		s.ProductBatches = section.ProductBatches
	}

	r.db[id] = s
	return s, nil
}

func (r *SectionMap) Delete(id int) error {
	_, exist := r.db[id]
	if !exist {
		return errors.ErrSectionNotFound
	}

	delete(r.db, id)
	return nil
}
