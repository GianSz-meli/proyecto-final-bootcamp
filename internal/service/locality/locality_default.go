package locality

import (
	"ProyectoFinal/internal/repository/locality"
	pkgErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"errors"
)

type LocalityDefault struct {
	repository locality.LocalityRepository
}

func NewLocalityService(repository locality.LocalityRepository) LocalityService {
	return &LocalityDefault{repository: repository}
}

func (s *LocalityDefault) Create(locality models.Locality) (models.Locality, error) {
	if locality.Id > 0 {
		exist, err := s.repository.GetById(locality.Id)
		if err != nil {
			if !errors.Is(err, pkgErrors.ErrNotFound) {
				return models.Locality{}, err
			}
		}
		if exist != nil {
			newError := pkgErrors.WrapErrAlreadyExist("locality", "id", locality.Id)
			return models.Locality{}, newError
		}
	}

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
