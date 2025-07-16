package utils
import "ProyectoFinal/pkg/models"

func GetLastId[T any](db map[int]T) int {
	lastId := 0
	for id := range db {
		if id > lastId {
			lastId = id
		}
	}
	return lastId
}
type RowScanner interface {
	Scan(dest ...any) error
}

func LocalityScan(rowScanner RowScanner, locality *models.Locality) error {
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

func SellerScan(rowScanner RowScanner, seller *models.Seller) error {
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