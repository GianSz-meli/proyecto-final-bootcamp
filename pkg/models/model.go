package models

// Model is a generic interface that defines a method for converting a model to its corresponding DTO.
type Model[T any] interface {
	// ModelToDoc converts the model to its corresponding DTO type.
	ModelToDoc() T
}

// Dto is a generic interface that defines a method for converting a DTO to its corresponding model.
type Dto[T any] interface {
	// DocToModel converts the DTO to its corresponding model type.
	DocToModel() T
}
