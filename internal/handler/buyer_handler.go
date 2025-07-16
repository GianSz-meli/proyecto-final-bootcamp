package handler

import (
	"ProyectoFinal/internal/handler/utils"
	"ProyectoFinal/internal/service/buyer"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"log"
	"net/http"
)

type BuyerHandler struct {
	service buyer.Service
}

// NewBuyerHandler creates and returns a new buyer HTTP handler instance.
func NewBuyerHandler(newService buyer.Service) *BuyerHandler {
	return &BuyerHandler{
		service: newService,
	}
}

// GetById returns an HTTP handler function that retrieves a specific buyer by ID.
// This endpoint handles GET requests to fetch individual buyer information.
func (h *BuyerHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, idErr := utils.GetParamInt(r, "id")
		if idErr != nil {
			log.Println(idErr)
			errors.HandleError(w, idErr)
			return
		}

		resp, respErr := h.service.GetById(id)
		if respErr != nil {
			log.Println(respErr)
			errors.HandleError(w, respErr)
			return
		}

		response.JSON(w, http.StatusOK, models.SuccessResponse{Data: resp.ModelToDoc()})
	}
}

// GetAll returns an HTTP handler function that retrieves all buyers.
// This endpoint handles GET requests to fetch the complete list of buyers.
func (h *BuyerHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := h.service.GetAll()
		if err != nil {
			log.Println(err)
			errors.HandleError(w, err)
			return
		}

		docs := make([]models.BuyerDoc, 0, len(resp))

		for _, b := range resp {
			docs = append(docs, b.ModelToDoc())
		}

		response.JSON(w, http.StatusOK, models.SuccessResponse{Data: docs})
	}
}

// Create returns an HTTP handler function that creates a new buyer.
// This endpoint handles POST requests to add new buyers to the system.
func (h *BuyerHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto models.BuyerCreateDTO

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

// Update returns an HTTP handler function that updates an existing buyer.
// This endpoint handles PATCH requests to modify buyer information.
func (h *BuyerHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, idErr := utils.GetParamInt(r, "id")
		if idErr != nil {
			log.Println(idErr)
			errors.HandleError(w, idErr)
			return
		}

		var dto models.BuyerUpdateDTO

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

		resp, updateErr := h.service.PatchUpdate(id, &dto)
		if updateErr != nil {
			log.Println(updateErr)
			errors.HandleError(w, updateErr)
			return
		}

		response.JSON(w, http.StatusOK, models.SuccessResponse{Data: resp.ModelToDoc()})
	}
}

// Delete returns an HTTP handler function that deletes a buyer by ID.
// This endpoint handles DELETE requests to remove buyers from the system.
func (h *BuyerHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, idErr := utils.GetParamInt(r, "id")
		if idErr != nil {
			log.Println(idErr)
			errors.HandleError(w, idErr)
			return
		}

		if err := h.service.Delete(id); err != nil {
			log.Println(err)
			errors.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// GetAllOrByIdWithOrderCount handles requests to get buyers with their purchase order count.
// If no 'id' query parameter is provided, returns all buyers with their order counts.
// If 'id' query parameter is provided, returns the specific buyer with their order count.
func (h *BuyerHandler) GetAllOrByIdWithOrderCount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, hasId, err := utils.GetOptionalQueryParamInt(r, "id")
		if err != nil {
			log.Println(err)
			errors.HandleError(w, err)
			return
		}

		if !hasId {
			h.handleGetAllWithOrderCount(w)
			return
		}

		h.handleGetByIdWithOrderCount(w, id)
	}
}

// handleGetAllWithOrderCount returns an HTTP handler function that retrieves all buyers with their order counts.
// This endpoint provides comprehensive buyer reporting with purchase order statistics for business analysis.
func (h *BuyerHandler) handleGetAllWithOrderCount(w http.ResponseWriter) {
	resp, err := h.service.GetAllWithOrderCount()
	if err != nil {
		log.Println(err)
		errors.HandleError(w, err)
		return
	}

	docs := make([]models.BuyerWithOrderCountDoc, 0, len(resp))

	for _, b := range resp {
		docs = append(docs, b.ModelToDoc())
	}

	response.JSON(w, http.StatusOK, models.SuccessResponse{Data: docs})
}

// handleGetByIdWithOrderCount returns an HTTP handler function that retrieves a buyer with their order count.
// This endpoint provides business intelligence by combining buyer data with purchase order statistics.
func (h *BuyerHandler) handleGetByIdWithOrderCount(w http.ResponseWriter, id int) {
	resp, respErr := h.service.GetByIdWithOrderCount(id)
	if respErr != nil {
		log.Println(respErr)
		errors.HandleError(w, respErr)
		return
	}

	response.JSON(w, http.StatusOK, models.SuccessResponse{Data: resp.ModelToDoc()})
}
