package errors

import (
	"errors"
	"fmt"
	"github.com/bootcamp-go/web/response"
	"net/http"
)

type ApiError struct {
	StatusCode int `json:"status_code"`
}

var (
	ErrGeneral             = errors.New("internal server error")
	ErrNotFound            = errors.New("not found")
	ErrAlreadyExists       = errors.New("resource already exists")
	ErrBadRequest          = errors.New("bad request")
	ErrUnprocessableEntity = errors.New("validation failed")
	mapErr                 = map[error]ApiError{
		ErrGeneral:             NewErrInternalServer(),
		ErrNotFound:            NewErrNotFound(),
		ErrAlreadyExists:       NewErrAlreadyExists(),
		ErrBadRequest:          NewErrBadRequest(),
		ErrUnprocessableEntity: NewErrUnprocessableEntity(),
	}
)

func NewErrInternalServer() ApiError {
	return ApiError{
		StatusCode: http.StatusInternalServerError,
	}
}

func NewErrNotFound() ApiError {
	return ApiError{
		StatusCode: http.StatusNotFound,
	}
}

func NewErrAlreadyExists() ApiError {
	return ApiError{
		StatusCode: http.StatusConflict,
	}
}

func NewErrBadRequest() ApiError {
	return ApiError{
		StatusCode: http.StatusBadRequest,
	}
}

func NewErrUnprocessableEntity() ApiError {
	return ApiError{
		StatusCode: http.StatusUnprocessableEntity,
	}
}

func getMappedError(err error) *ApiError {
	for baseError, mappedError := range mapErr {
		if errors.Is(err, baseError) {
			return &mappedError
		}
	}
	return nil
}
func HandleError(w http.ResponseWriter, err error) {
	if mappedError := getMappedError(err); mappedError != nil {
		response.Error(w, mappedError.StatusCode, err.Error())
		return
	}
	response.Error(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

func WrapErrAlreadyExist(domain, property string, value int) error {
	return fmt.Errorf("%w : %s with %s %d already exists", ErrAlreadyExists, domain, property, value)
}
func WrapErrBadRequest(err error) error {
	return fmt.Errorf("%w : %s", ErrBadRequest, err.Error())
}
