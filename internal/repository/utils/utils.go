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
		&locality.ProvinceName,
		&locality.CountryName,
	); err != nil {
		return err
	}
	return nil
}
