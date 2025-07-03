package handler

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
	warehouses := h.warehouseService.GetAllWarehouses()
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
		errors.HandleError(w, errors.WrapErrBadRequest(err))
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
		Data: createdWarehouse.ModelToDoc(),
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
		errors.HandleError(w, errors.WrapErrBadRequest(err))
		return
	}

	if err := utils.ValidateRequestData(updateRequest); err != nil {
		errors.HandleError(w, err)
		return
	}

	currentWarehouse, err := h.warehouseService.GetWarehouseById(id)
	if err != nil {
		errors.HandleError(w, err)
		return
	}

	if updated := utils.UpdateFields(currentWarehouse, &updateRequest); !updated {
		newError := errors.WrapErrUnprocessableEntity(fmt.Errorf("no fields provided for update"))
		errors.HandleError(w, newError)
		return
	}

	updatedWarehouse, err := h.warehouseService.UpdateWarehouse(id, currentWarehouse)
	if err != nil {
		errors.HandleError(w, err)
		return
	}

	responseBody := models.SuccessResponse{
		Data: updatedWarehouse.ModelToDoc(),
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
