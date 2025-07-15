package handler

import (
	"ProyectoFinal/internal/handler/utils"
	"ProyectoFinal/internal/service/product_batch"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"log"
	"net/http"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
)

type ProductBatchHandler struct {
	service product_batch.ProductBatchService
}

func NewProductBatchHandler(service product_batch.ProductBatchService) *ProductBatchHandler {
	return &ProductBatchHandler{
		service: service,
	}
}

func (h *ProductBatchHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody models.ProductBatchCreateRequest

		if err := request.JSON(r, &reqBody); err != nil {
			newErr := errors.WrapErrBadRequest(err)
			log.Println(newErr)
			errors.HandleError(w, newErr)
			return
		}

		if err := utils.ValidateRequestData(reqBody); err != nil {
			log.Println(err)
			errors.HandleError(w, err)
			return
		}

		model := reqBody.CreateRequestToModel()

		createdProductBatch, err := h.service.Create(model)
		if err != nil {
			log.Printf("[ProductBatchHandler][Create] error: %v", err)
			errors.HandleError(w, err)
			return
		}

		response.JSON(w, http.StatusCreated, models.SuccessResponse{
			Data: createdProductBatch.ModelToDoc(),
		})
	}
}

func (h *ProductBatchHandler) GetProductCountBySection() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sectionID, err := utils.GetParamInt(r, "id")
		if err != nil {
			log.Println(err)
			errors.HandleError(w, err)
			return
		}

		reports, err := h.service.GetProductCountBySection(&sectionID)
		if err != nil {
			log.Printf("[ProductBatchHandler][GetProductCountBySection] error: %v", err)
			errors.HandleError(w, err)
			return
		}

		response.JSON(w, http.StatusOK, models.SuccessResponse{
			Data: reports,
		})
	}
}
