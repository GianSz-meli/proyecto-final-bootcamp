package models

type ApiError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func (a *ApiError) Error() string {
	return a.Message
}
