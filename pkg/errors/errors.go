package errors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/bootcamp-go/web/response"
)

// ApiError represents a custom API error with an associated HTTP status code.
type ApiError struct {
	StatusCode int `json:"status_code"`
}

var (
	// ErrGeneral indicates an internal server error.
	ErrGeneral = errors.New("internal server error")

	// ErrNotFound indicates a resource was not found.
	ErrNotFound = errors.New("not found")

	// ErrConflict indicates a resource conflict.
	ErrConflict = errors.New("conflict")

	// ErrBadRequest indicates the request is invalid.
	ErrBadRequest = errors.New("bad request")

	// ErrUnprocessableEntity indicates the request cannot be processed due to semantic errors.
	ErrUnprocessableEntity = errors.New("unprocessable entity")

	// mapErr maps common errors to their corresponding ApiError definitions.
	mapErr = map[error]ApiError{
		ErrGeneral:             NewErrInternalServer(),
		ErrNotFound:            NewErrNotFound(),
		ErrConflict:            NewErrConflict(),
		ErrBadRequest:          NewErrBadRequest(),
		ErrUnprocessableEntity: NewErrUnprocessableEntity(),
	}
)

// NewErrInternalServer returns an ApiError with status 500 Internal Server Error.
func NewErrInternalServer() ApiError {
	return ApiError{
		StatusCode: http.StatusInternalServerError,
	}
}

// NewErrNotFound returns an ApiError with status 404 Not Found.
func NewErrNotFound() ApiError {
	return ApiError{
		StatusCode: http.StatusNotFound,
	}
}

// NewErrConflict returns an ApiError with status 409 Conflict.
func NewErrConflict() ApiError {
	return ApiError{
		StatusCode: http.StatusConflict,
	}
}

// NewErrBadRequest returns an ApiError with status 400 Bad Request.
func NewErrBadRequest() ApiError {
	return ApiError{
		StatusCode: http.StatusBadRequest,
	}
}

// NewErrUnprocessableEntity returns an ApiError with status 422 Unprocessable Entity.
func NewErrUnprocessableEntity() ApiError {
	return ApiError{
		StatusCode: http.StatusUnprocessableEntity,
	}
}

// getMappedError returns the mapped ApiError for the given error, or nil if not found.
func getMappedError(err error) *ApiError {
	for baseError, mappedError := range mapErr {
		if errors.Is(err, baseError) {
			return &mappedError
		}
	}
	return nil
}

// HandleError processes the provided error, maps it to an API error, and writes the appropriate HTTP response.
func HandleError(w http.ResponseWriter, err error) {
	err = HandleMysqlError(err)
	if mappedError := getMappedError(err); mappedError != nil {
		response.Error(w, mappedError.StatusCode, err.Error())
		return
	}

	response.Error(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

// WrapErrConflict wraps a conflict error with domain, property, and value information.
func WrapErrConflict(domain, property string, value any) error {
	return fmt.Errorf("%w : %s with %s %v already exists", ErrConflict, domain, property, value)
}

// WrapErrBadRequest wraps a bad request error with the underlying error message.
func WrapErrBadRequest(err error) error {
	return fmt.Errorf("%w : %s", ErrBadRequest, err.Error())
}

// WrapErrUnprocessableEntity wraps an unprocessable entity error with the underlying error message.
func WrapErrUnprocessableEntity(err error) error {
	return fmt.Errorf("%w : %s", ErrUnprocessableEntity, err.Error())
}

// WrapErrNotFound wraps a not found error with domain, property, and value information.
func WrapErrNotFound(domain, property string, value any) error {
	return fmt.Errorf("%w : %s with %s %v not found", ErrNotFound, domain, property, value)
}
