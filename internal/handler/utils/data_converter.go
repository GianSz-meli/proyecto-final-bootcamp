package utils

import (
	"ProyectoFinal/pkg/errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// GetParamInt gets a parameter from the request, converts it to an integer and returns an error if the conversion fails
func GetParamInt(r *http.Request, paramName string) (int, error) {
	value := chi.URLParam(r, paramName)
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0, errors.WrapErrBadRequest(err)
	}
	return result, nil
}
