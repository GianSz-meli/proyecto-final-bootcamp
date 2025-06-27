package loader

import "ProyectoFinal/pkg/models"

type SectionLoader interface {
	Load() (s map[int]models.Section, err error)
}
