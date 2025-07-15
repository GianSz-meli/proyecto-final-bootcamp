package errors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/go-sql-driver/mysql"
)

type ApiError struct {
	StatusCode int `json:"status_code"`
}

var (
	ErrGeneral             = errors.New("internal server error")
	ErrNotFound            = errors.New("not found")
	ErrConflict            = errors.New("conflict")
	ErrBadRequest          = errors.New("bad request")
	ErrUnprocessableEntity = errors.New("unprocessable entity")
	mapErr                 = map[error]ApiError{
		ErrGeneral:             NewErrInternalServer(),
		ErrNotFound:            NewErrNotFound(),
		ErrConflict:            NewErrConflict(),
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

func NewErrConflict() ApiError {
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
	err = HandleMysqlError(err)
	if mappedError := getMappedError(err); mappedError != nil {
		response.Error(w, mappedError.StatusCode, err.Error())
		return
	}

	response.Error(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

func WrapErrConflict(domain, property string, value any) error {
	return fmt.Errorf("%w : %s with %s %v already exists", ErrConflict, domain, property, value)
}

func WrapErrBadRequest(err error) error {
	return fmt.Errorf("%w : %s", ErrBadRequest, err.Error())
}

func WrapErrUnprocessableEntity(err error) error {
	return fmt.Errorf("%w : %s", ErrUnprocessableEntity, err.Error())
}

func WrapErrNotFound(domain, property string, value any) error {
	return fmt.Errorf("%w : %s with %s %v not found", ErrNotFound, domain, property, value)
}
