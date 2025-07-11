package models

type Model[T any] interface {
	ModelToDoc() T
}

type Dto[T any] interface {
	DocToModel() T
}
