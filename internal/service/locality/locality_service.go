package locality

import "ProyectoFinal/pkg/models"

// LocalityService defines business logic methods for managing localities and their sellers.
type LocalityService interface {
	// Create adds a new locality.
	Create(locality models.Locality) (models.Locality, error)

	// GetById returns a locality by its ID.
	GetById(id int) (models.Locality, error)

	// GetSellersByLocalities returns seller reports for all localities.
	GetSellersByLocalities() ([]models.SellersByLocalityReport, error)

	// GetSellersByIdLocality returns the seller report for a specific locality by its ID.
	GetSellersByIdLocality(idLocality int) (models.SellersByLocalityReport, error)
}
