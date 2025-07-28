package locality

import (
	"ProyectoFinal/internal/repository/utils"
	"ProyectoFinal/pkg/models"
)

// SellersByLocalityScan scans a database row into a SellersByLocalityReport model.
// The rowScanner parameter must implement the RowScanner interface.
func SellersByLocalityScan(rowScanner utils.RowScanner, sellersByLocality *models.SellersByLocalityReport) error {
	if err := rowScanner.Scan(
		&sellersByLocality.SellersCount,
		&sellersByLocality.LocalityId,
		&sellersByLocality.LocalityName,
	); err != nil {
		return err
	}
	return nil
}
