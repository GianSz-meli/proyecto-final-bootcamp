package errors

import (
	"errors"
	"net/http"

	"github.com/bootcamp-go/web/response"
)

type ApiError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func (a *ApiError) Error() string {
	return a.Message
}

var (
	ErrGeneral = errors.New("unexpected general error")

	ErrSectionNotFound      = errors.New("section not found")
	ErrSectionAlreadyExists = errors.New("section already exists")
	ErrInvalidSectionID     = errors.New("invalid section ID")
	ErrInvalidSectionData   = errors.New("invalid section data")
	ErrSectionNumberExists  = errors.New("section number already exists")

	mapErr = map[error]ApiError{
		ErrGeneral:              NewErrInternalServer(ErrGeneral.Error()),
		ErrSectionNotFound:      NewErrNotFound(ErrSectionNotFound.Error()),
		ErrSectionAlreadyExists: NewErrConflict(ErrSectionAlreadyExists.Error()),
		ErrInvalidSectionID:     NewErrBadRequest(ErrInvalidSectionID.Error()),
		ErrInvalidSectionData:   NewErrBadRequest(ErrInvalidSectionData.Error()),
		ErrSectionNumberExists:  NewErrConflict(ErrSectionNumberExists.Error()),
	}
)

func NewErrInternalServer(msg string) ApiError {
	return ApiError{
		StatusCode: http.StatusInternalServerError,
		Message:    msg,
	}
}

func NewErrNotFound(msg string) ApiError {
	return ApiError{
		StatusCode: http.StatusNotFound,
		Message:    msg,
	}
}

func NewErrBadRequest(msg string) ApiError {
	return ApiError{
		StatusCode: http.StatusBadRequest,
		Message:    msg,
	}
}

func NewErrConflict(msg string) ApiError {
	return ApiError{
		StatusCode: http.StatusConflict,
		Message:    msg,
	}
}

func HandleError(w http.ResponseWriter, err error) {
	mappedError, ok := mapErr[err]

	if !ok {
		response.Error(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	response.Error(w, mappedError.StatusCode, mappedError.Message)
}
