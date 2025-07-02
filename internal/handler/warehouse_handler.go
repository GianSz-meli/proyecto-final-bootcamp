package handler

import (
	"net/http"

	"ProyectoFinal/internal/service/warehouse"
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
