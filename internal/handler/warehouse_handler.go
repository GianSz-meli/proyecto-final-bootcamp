package handler

import (
	"encoding/json"
	"net/http"

	"ProyectoFinal/internal/handler/utils"
	"ProyectoFinal/internal/service/warehouse"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"

	"github.com/bootcamp-go/web/response"
	"github.com/go-playground/validator/v10"
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
	for _, warehouse := range warehouses {
		warehousesDoc = append(warehousesDoc, warehouse.ModelToDoc())
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

	warehouse, err := h.warehouseService.GetWarehouseById(id)
	if err != nil {
		errors.HandleError(w, err)
		return
	}

	responseBody := models.SuccessResponse{
		Data: warehouse.ModelToDoc(),
	}
	response.JSON(w, http.StatusOK, responseBody)
}

func (h *WarehouseHandler) CreateWarehouse(w http.ResponseWriter, r *http.Request) {
	var createRequest models.CreateWarehouseRequest
	if err := json.NewDecoder(r.Body).Decode(&createRequest); err != nil {
		errors.HandleError(w, errors.WrapErrBadRequest(err))
		return
	}

	validate := validator.New()
	if err := validate.Struct(createRequest); err != nil {
		errors.HandleError(w, errors.WrapErrUnprocessableEntity(err))
		return
	}

	warehouse := createRequest.DocToModel()

	createdWarehouse, err := h.warehouseService.CreateWarehouse(warehouse)
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

	validate := validator.New()
	if err := validate.Struct(updateRequest); err != nil {
		errors.HandleError(w, errors.WrapErrUnprocessableEntity(err))
		return
	}

	warehouse := updateRequest.DocToModel()
	updatedWarehouse, err := h.warehouseService.UpdateWarehouse(id, warehouse)
	if err != nil {
		errors.HandleError(w, err)
		return
	}

	responseBody := models.SuccessResponse{
		Data: updatedWarehouse.ModelToDoc(),
	}
	response.JSON(w, http.StatusOK, responseBody)
}