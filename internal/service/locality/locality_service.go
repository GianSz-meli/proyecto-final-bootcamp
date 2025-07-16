package locality

import "ProyectoFinal/pkg/models"

type LocalityService interface {
	Create(locality models.Locality) (models.Locality, error)
	GetById(id int) (models.Locality, error)
	GetSellersByLocalities() ([]models.SellersByLocalityReport, error)
	GetSellersByIdLocality(idLocality int) (models.SellersByLocalityReport, error)
}
