package locality

import "ProyectoFinal/pkg/models"

// LocalityRepository defines methods for locality data management.
type LocalityRepository interface {

	// Create inserts a new locality record into the repository.
	Create(locality models.Locality) (models.Locality, error)

	// GetById retrieves a locality by its ID.
	GetById(id int) (*models.Locality, error)

	// GetSellersByIdLocality retrieves the seller report for a specific locality by its ID.
	GetSellersByIdLocality(idLocality int) (models.SellersByLocalityReport, error)

	// GetSellersByLocalities retrieves seller reports for all localities.
	GetSellersByLocalities() ([]models.SellersByLocalityReport, error)

	// ReportCarriersByLocality retrieves the carrier report for a specific locality by its ID or all localities if ID is nil.
	ReportCarriersByLocality(id *int) ([]models.CarrierReport, error)
}
