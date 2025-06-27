package models

type SuccessResponse struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data"`
}
