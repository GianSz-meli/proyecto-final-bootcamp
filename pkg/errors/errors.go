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
	ErrGeneral = errors.New("unexepected general error")

	mapErr = map[error]ApiError{
		ErrGeneral: NewErrInternalServer(ErrGeneral.Error()),
	}
)

func NewErrInternalServer(msg string) ApiError {
	return ApiError{
		StatusCode: http.StatusInternalServerError,
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
