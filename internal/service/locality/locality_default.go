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

func (l *LocalityDefault) Create(locality models.Locality) (models.Locality, error) {
	newlocality, err := l.repository.Create(locality)
	if err != nil {
		return models.Locality{}, err
	}

	return newlocality, nil
}

func (l *LocalityDefault) GetById(id int) (models.Locality, error) {
	locality, err := l.repository.GetById(id)
	if err != nil {
		return models.Locality{}, err
	}
	return *locality, nil
}

func (l *LocalityDefault) GetSellersByLocalities() ([]models.SellersByLocalityReport, error) {

	sellersByLocality, err := l.repository.GetSellersByLocalities()

	if err != nil {
		return nil, err
	}
	return sellersByLocality, nil
}

func (l *LocalityDefault) GetSellersByIdLocality(idLocality int) (models.SellersByLocalityReport, error) {
	sellerByLocality, err := l.repository.GetSellersByIdLocality(idLocality)

	if err != nil {
		return models.SellersByLocalityReport{}, err
	}

	return sellerByLocality, nil

}
