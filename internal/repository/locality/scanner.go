package locality

import (
	"ProyectoFinal/internal/repository/utils"
	"ProyectoFinal/pkg/models"
)

func LocalityScan(rowScanner utils.RowScanner, locality *models.Locality) error {
	if err := rowScanner.Scan(
		&locality.Id,
		&locality.LocalityName,
		&locality.Province.Id,
		&locality.Province.ProvinceName,
		&locality.Province.Country.Id,
		&locality.Province.Country.CountryName,
	); err != nil {
		return err
	}
	return nil
}

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
