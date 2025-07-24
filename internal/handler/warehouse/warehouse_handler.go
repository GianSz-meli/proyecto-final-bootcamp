package warehouse

import (
	"encoding/json"
	"fmt"
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
	warehouses, err := h.warehouseService.GetAllWarehouses()
	if err != nil {
		errors.HandleError(w, err)
		return
	}
	warehousesDoc := make([]models.WarehouseDocument, 0, len(warehouses))
	for _, wh := range warehouses {
		warehousesDoc = append(warehousesDoc, wh.ModelToDoc())
	}
	responseBody := models.SuccessResponse{
		Data: warehousesDoc,
	}
	response.JSON(w, http.StatusOK, responseBody)
}

func (h *WarehouseHandler) GetWarehouseById(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetParamInt(r, "id")
	if err != nil {
		errors.HandleError(w, err)
		return
	}

	wh, err := h.warehouseService.GetWarehouseById(id)
	if err != nil {
		errors.HandleError(w, err)
		return
	}

	responseBody := models.SuccessResponse{
		Data: wh.ModelToDoc(),
	}
	response.JSON(w, http.StatusOK, responseBody)
}

func (h *WarehouseHandler) CreateWarehouse(w http.ResponseWriter, r *http.Request) {
	var createRequest models.CreateWarehouseRequest
	if err := json.NewDecoder(r.Body).Decode(&createRequest); err != nil {
		errors.HandleError(w, errors.WrapErrBadRequest(fmt.Errorf("it was not possible to decode json")))
		return
	}

	if err := utils.ValidateRequestData(createRequest); err != nil {
		errors.HandleError(w, err)
		return
	}

	wh := createRequest.DocToModel()

	createdWarehouse, err := h.warehouseService.CreateWarehouse(wh)
	if err != nil {
		errors.HandleError(w, err)
		return
	}

	responseBody := models.SuccessResponse{
		Data: createdWarehouse.ModelToCreateDoc(),
	}
	response.JSON(w, http.StatusCreated, responseBody)
}

func (h *WarehouseHandler) UpdateWarehouse(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetParamInt(r, "id")
	if err != nil {
		errors.HandleError(w, err)
		return
	}

	var updateRequest models.UpdateWarehouseRequest
	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		errors.HandleError(w, errors.WrapErrBadRequest(fmt.Errorf("it was not possible to decode json")))
		return
	}

	if err := utils.ValidateRequestData(updateRequest); err != nil {
		errors.HandleError(w, err)
		return
	}

	updatedWarehouse, err := h.warehouseService.UpdateWarehouse(id, updateRequest)
	if err != nil {
		errors.HandleError(w, err)
		return
	}

	responseBody := models.SuccessResponse{
		Data: updatedWarehouse.ModelToUpdateDoc(),
	}
	response.JSON(w, http.StatusOK, responseBody)
}

func (h *WarehouseHandler) DeleteWarehouse(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetParamInt(r, "id")
	if err != nil {
		errors.HandleError(w, err)
		return
	}

	err = h.warehouseService.DeleteWarehouse(id)
	if err != nil {
		errors.HandleError(w, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
