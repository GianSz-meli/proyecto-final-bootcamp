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
	ErrAlreadyExists       = errors.New("resource already exists")
	ErrBadRequest          = errors.New("bad request")
	ErrUnprocessableEntity = errors.New("unprocessable entity")

	mapErr = map[error]ApiError{
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

	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		handleMySQLError(w, mysqlErr)
		return
	}

	response.Error(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

func handleMySQLError(w http.ResponseWriter, mysqlErr *mysql.MySQLError) {
	switch mysqlErr.Number {
	case 1062: // Duplicate entry for key
		response.Error(w, http.StatusConflict, "Resource already exists. The given data violates a unique constraint")
	case 1452: // Foreign key constraint fails
		response.Error(w, http.StatusBadRequest, "Invalid reference. The given data references an invalid or missing record")
	case 1048: // Column cannot be null
		response.Error(w, http.StatusBadRequest, "Required field cannot be empty")
	case 1366: // Incorrect data type
		response.Error(w, http.StatusBadRequest, "Invalid data type")
	default:
		response.Error(w, http.StatusInternalServerError, "Technical error when connecting to the database")
	}
}

func WrapErrAlreadyExist(domain, property string, value any) error {
	return fmt.Errorf("%w : %s with %s %v already exists", ErrAlreadyExists, domain, property, value)
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
