package service

import "ProyectoFinal/pkg/models"

// SectionService defines the contract for section business logic operations
type SectionService interface {
	GetAll() (sections []models.Section, err error)
	GetById(id int) (s models.Section, err error)
	Create(section models.Section) (createdSection models.Section, err error)
	Update(id int, section models.Section) (updatedSection models.Section, err error)
	Delete(id int) (err error)
}
