package locality

import "ProyectoFinal/pkg/models"

type LocalityRepository interface {
	Create(locality models.Locality) (models.Locality, error)
	GetById(id int) (*models.Locality, error)
}
