package locality

import (
	"ProyectoFinal/pkg/models"
)

type LocalityRepository interface {
	GetById(id int) (*models.Locality, error)
}