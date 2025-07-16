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

// RowScanner is an interface for scanning database rows into destination variables.
type RowScanner interface {
	Scan(dest ...any) error
}
