package utils

import (
	"ProyectoFinal/pkg/models"
)

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
