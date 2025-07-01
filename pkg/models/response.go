package models

type SuccessResponse[T any] struct {
	Data []T `json:"data"`
}
