package handler

import (
	"net/http"

	"ProyectoFinal/internal/handler/utils"
	"ProyectoFinal/internal/service/warehouse"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"

	"github.com/bootcamp-go/web/response"
)

type WarehouseHandler struct {
	warehouseService warehouse.WarehouseService
}

func NewWarehouseHandler(warehouseService warehouse.WarehouseService) *WarehouseHandler {
	return &WarehouseHandler{
		warehouseService: warehouseService,
	}
}

func (h *WarehouseHandler) GetAllWarehouses(w http.ResponseWriter, r *http.Request) {
	warehouses := h.warehouseService.GetAllWarehouses()
	responseBody := models.SuccessResponse{
		Data: warehouses,
	}
	response.JSON(w, http.StatusOK, responseBody)
}

func (h *WarehouseHandler) GetWarehouseById(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetParamInt(r, "id")
	if err != nil {
		errors.HandleError(w, err)
		return
	}

	warehouse, err := h.warehouseService.GetWarehouseById(id)
	if err != nil {
		errors.HandleError(w, err)
		return
	}

	responseBody := models.SuccessResponse{
		Data: warehouse,
	}
	response.JSON(w, http.StatusOK, responseBody)
}
