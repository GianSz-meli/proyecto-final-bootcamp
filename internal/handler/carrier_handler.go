package handler

import (
	"ProyectoFinal/internal/handler/utils"
	"ProyectoFinal/internal/service/carrier"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"net/http"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
)

type CarrierHandler struct {
	service carrier.CarrierService
}

func NewCarrierHandler(newService carrier.CarrierService) *CarrierHandler {
	return &CarrierHandler{
		service: newService,
	}
}

func (h *CarrierHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto models.CarrierCreateDTO

		if err := request.JSON(r, &dto); err != nil {
			errors.HandleError(w, errors.WrapErrBadRequest(err))
			return
		}

		if validateDtoErr := utils.ValidateRequestData(dto); validateDtoErr != nil {
			errors.HandleError(w, validateDtoErr)
			return
		}

		resp, respErr := h.service.Create(dto.CreateDtoToModel())
		if respErr != nil {
			errors.HandleError(w, respErr)
			return
		}

		response.JSON(w, http.StatusCreated, models.SuccessResponse{Data: resp.ModelToDoc()})
	}
}
