package locality

import (
	"ProyectoFinal/internal/repository/locality"
	"ProyectoFinal/pkg/models"
)

type LocalityDefault struct {
	repository locality.LocalityRepository
}

func NewLocalityService(repository locality.LocalityRepository) LocalityService {
	return &LocalityDefault{repository: repository}
}

func (s *LocalityDefault) Create(locality models.Locality) (models.Locality, error) {
	newlocality, err := s.repository.Create(locality)

	if err != nil {
		return models.Locality{}, err
	}

	return newlocality, nil
}

func (s *LocalityDefault) GetById(id int) (models.Locality, error) {
	locality, err := s.repository.GetById(id)
	if err != nil {
		return models.Locality{}, err
	}
	return *locality, nil
}
