package handler

import (
	"net/http"
	"strconv"

	"ProyectoFinal/internal/handler/utils"
	inboundOrderService "ProyectoFinal/internal/service/inbound_order"
	pkgErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
)

type InboundOrderHandler struct {
	service inboundOrderService.Service
}

func NewInboundOrderHandler(service inboundOrderService.Service) *InboundOrderHandler {
	return &InboundOrderHandler{
		service: service,
	}
}

func (h *InboundOrderHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody models.InboundOrderRequest

		if err := request.JSON(r, &reqBody); err != nil {
			newErr := pkgErrors.WrapErrBadRequest(err)
			pkgErrors.HandleError(w, newErr)
			return
		}

		if err := utils.ValidateRequestData(reqBody); err != nil {
			pkgErrors.HandleError(w, err)
			return
		}

		inboundOrder := reqBody.DocToModel()

		createdInboundOrder, err := h.service.Create(inboundOrder)

		if err != nil {
			pkgErrors.HandleError(w, err)
			return
		}

		responseBody := models.SuccessResponse{
			Data: createdInboundOrder.ModelToDoc(),
		}
		response.JSON(w, http.StatusCreated, responseBody)
	}
}

func (h *InboundOrderHandler) GetEmployeeInboundOrdersReport() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener el parámetro id de la query string
		idParam := r.URL.Query().Get("id")

		if idParam == "" {
			// Si no hay id, devolver el reporte de todos los empleados
			reports, err := h.service.GetEmployeeInboundOrdersReportAll()
			if err != nil {
				pkgErrors.HandleError(w, err)
				return
			}

			responseBody := models.SuccessResponse{
				Data: reports,
			}
			response.JSON(w, http.StatusOK, responseBody)
			return
		}

		// Si hay id, convertir a int y obtener el reporte específico
		id, err := strconv.Atoi(idParam)
		if err != nil {
			newErr := pkgErrors.WrapErrBadRequest(err)
			pkgErrors.HandleError(w, newErr)
			return
		}

		report, err := h.service.GetEmployeeInboundOrdersReportByEmployeeId(id)
		if err != nil {
			pkgErrors.HandleError(w, err)
			return
		}

		responseBody := models.SuccessResponse{
			Data: []models.EmployeeInboundOrdersReport{report},
		}
		response.JSON(w, http.StatusOK, responseBody)
	}
}
