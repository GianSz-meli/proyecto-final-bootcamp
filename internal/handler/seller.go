package handler

import (
	"ProyectoFinal/internal/service"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type SellerHandler struct {
	service service.Service
}

func NewSellerHandler(service service.Service) *SellerHandler {
	return &SellerHandler{service: service}
}

var validate = validator.New()

func (h *SellerHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody models.SellerDoc

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

		seller, err := h.service.Create(model)

		body := seller.ModelToDoc()
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		response.JSON(w, http.StatusCreated, body)

	}

}
