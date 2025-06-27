package loader

type Loader[T any] interface {
	Load() (map[int]T, error)
}
