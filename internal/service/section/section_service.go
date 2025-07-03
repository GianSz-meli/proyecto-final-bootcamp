package service

import "ProyectoFinal/pkg/models"

type SectionService interface {
	GetAll() (s []models.Section, err error)
	GetByID(id int) (s models.Section, err error)
	Create(section models.Section) (s models.Section, err error)
	Update(id int, section models.Section) (s models.Section, err error)
	Delete(id int) (err error)
}
