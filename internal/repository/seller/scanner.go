package seller

import (
	"ProyectoFinal/internal/repository/utils"
	"ProyectoFinal/pkg/models"
)

func SellerScan(rowScanner utils.RowScanner, seller *models.Seller) error {
	if err := rowScanner.Scan(
		&seller.Id,
		&seller.Cid,
		&seller.CompanyName,
		&seller.Address,
		&seller.Telephone,
		&seller.LocalityId,
	); err != nil {
		return err
	}
	return nil
}
