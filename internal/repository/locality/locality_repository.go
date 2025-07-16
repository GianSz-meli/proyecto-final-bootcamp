package locality

import "ProyectoFinal/pkg/models"

type LocalityRepository interface {
	Create(locality models.Locality) (models.Locality, error)
	GetById(id int) (*models.Locality, error)
	GetSellersByIdLocality(idLocality int) (models.SellersByLocalityReport, error)
	GetSellersByLocalities() ([]models.SellersByLocalityReport, error)
}
