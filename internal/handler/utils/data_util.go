package utils

import (
	customErrors "ProyectoFinal/pkg/errors"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// ValidateRequestData validates the request data
func ValidateRequestData(s any) error {
	if err := validate.Struct(s); err != nil {
		return customErrors.WrapErrUnprocessableEntity(err)
	}
	return nil
}

// GetParamInt gets a parameter from the request, converts it to an integer and returns an error if the conversion fails
func GetParamInt(r *http.Request, paramName string) (int, error) {
	value := chi.URLParam(r, paramName)
	result, err := strconv.Atoi(value)

	if result <= 0 {
		return 0, customErrors.WrapErrBadRequest(errors.New("id must be greater than 0"))
	}

	if err != nil {
		return 0, customErrors.WrapErrBadRequest(err)
	}
	return result, nil
}
