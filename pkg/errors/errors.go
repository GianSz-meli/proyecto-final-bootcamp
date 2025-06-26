package errors

import (
	"ProyectoFinal/internal/handler/utils"
	"ProyectoFinal/pkg/models"
	"errors"
	"net/http"
)

var (
	ErrNotFound            = errors.New("not found")
	ErrBadRequest          = errors.New("bad request")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrInternalServerError = errors.New("internal server error")
)

var mapErr = map[error]models.ApiError{
	ErrNotFound:            NewErrNotFound(),
	ErrBadRequest:          NewErrBadRequest(),
	ErrUnauthorized:        NewErrUnauthorized(),
	ErrInternalServerError: NewErrInternalServer(),
}

func NewErrNotFound() models.ApiError {
	return models.ApiError{
		StatusCode: http.StatusNotFound,
		Message:    ErrNotFound.Error(),
	}
}
func NewErrBadRequest() models.ApiError {
	return models.ApiError{
		StatusCode: http.StatusBadRequest,
		Message:    ErrBadRequest.Error(),
	}
}
func NewErrUnauthorized() models.ApiError {
	return models.ApiError{
		StatusCode: http.StatusBadRequest,
		Message:    ErrUnauthorized.Error(),
	}
}
func NewErrInternalServer() models.ApiError {
	return models.ApiError{
		StatusCode: http.StatusBadRequest,
		Message:    ErrInternalServerError.Error(),
	}
}

func HandleError(w http.ResponseWriter, err error, customMsg ...string) {
	mappedError, ok := mapErr[err]
	if !ok {
		apiError := NewErrInternalServer()
		utils.SendError(w, apiError)
		return
	}
	if len(customMsg) > 0 && customMsg[0] != "" {
		//msgError := fmt.Errorf("%s %w", customMsg[0], err)
		mappedError.Message = customMsg[0]
	}
	utils.SendError(w, mappedError)
	return
}
