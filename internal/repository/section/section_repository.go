package repository

import "ProyectoFinal/pkg/models"

type SectionRepository interface {
	GetAll() (s []models.Section, err error)
	GetById(id int) (s models.Section, exists bool)
	Create(section models.Section) (s models.Section, err error)
	Update(id int, section models.Section) (s models.Section, err error)
	Delete(id int) (err error)
	ExistBySectionNumber(sectionNumber string) bool
}
