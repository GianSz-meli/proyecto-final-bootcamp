package handler

import (
	service "ProyectoFinal/internal/service/products"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"

	//"github.com/go-playground/validator/v10"
	"net/http"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var reqBody models.ProductDoc

	if err := request.JSON(r, &reqBody); err != nil {
		newErr := errors.WrapErrBadRequest(err)
		errors.HandleError(w, newErr)
		return
	}

	if err := validate.Struct(reqBody); err != nil {
		newError := errors.WrapErrBadRequest(err)
		errors.HandleError(w, newError)
		return
	}

	model := reqBody.DocToModel()

	product, err := h.service.CreateProduct(model)

	if err != nil {
		errors.HandleError(w, err)
		return
	}

	body := map[string]models.ProductDoc{
		"data": product.ModelToDoc(),
	}
	response.JSON(w, http.StatusCreated, body)

}
