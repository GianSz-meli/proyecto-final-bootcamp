package handler

import (
	"ProyectoFinal/internal/handler/utils"
	"ProyectoFinal/internal/service/purchase_order"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"log"
	"net/http"
)

type PurchaseOrderHandler struct {
	service purchase_order.Service
}

func NewpPurchaseOrderHandler(newService purchase_order.Service) *PurchaseOrderHandler {
	return &PurchaseOrderHandler{
		service: newService,
	}
}

func (h *PurchaseOrderHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto models.PurchaseOrderCreateDTO

		if err := request.JSON(r, &dto); err != nil {
			wrappedErr := errors.WrapErrBadRequest(err)
			log.Println(wrappedErr)
			errors.HandleError(w, wrappedErr)
			return
		}

		if validateDtoErr := utils.ValidateRequestData(dto); validateDtoErr != nil {
			log.Println(validateDtoErr)
			errors.HandleError(w, validateDtoErr)
			return
		}

		resp, respErr := h.service.Create(dto.CreateDtoToModel())
		if respErr != nil {
			log.Println(respErr)
			errors.HandleError(w, respErr)
			return
		}

		response.JSON(w, http.StatusCreated, models.SuccessResponse{Data: resp.ModelToDoc()})
	}
}
