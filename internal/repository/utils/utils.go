package utils

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
