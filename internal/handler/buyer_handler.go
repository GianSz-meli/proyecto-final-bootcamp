package handler

import (
	"ProyectoFinal/internal/handler/utils"
	"ProyectoFinal/internal/service/buyer"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/bootcamp-go/web/response"
	"net/http"
)

type BuyerHandler struct {
	service buyer.Service
}

func NewBuyerHandler(newService buyer.Service) *BuyerHandler {
	return &BuyerHandler{
		service: newService,
	}
}

func (h *BuyerHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := h.service.GetAll()

		docs := make([]models.BuyerDoc, 0, len(resp))
		for _, b := range resp {
			docs = append(docs, b.ModelToDoc())
		}
		response.JSON(w, http.StatusOK, models.SuccessResponse{Data: docs})
	}
}

func (h *BuyerHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var dto models.BuyerCreateDTO
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			errors.HandleError(w, errors.WrapErrBadRequest(err))
			return
		}

		if err := utils.ValidateRequestData(dto); err != nil {
			errors.HandleError(w, err)
			return
		}

		resp, err := h.service.Create(dto.CreateDtoToModel())
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		doc := resp.ModelToDoc()
		response.JSON(w, http.StatusCreated, models.SuccessResponse{Data: doc})
	}
}

func (h *BuyerHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := utils.GetParamInt(r, "id")
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		resp, err := h.service.GetById(id)
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		doc := resp.ModelToDoc()
		response.JSON(w, http.StatusOK, models.SuccessResponse{Data: doc})
	}
}

func (h *BuyerHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := utils.GetParamInt(r, "id")
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		if err := h.service.Delete(id); err != nil {
			errors.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func (h *BuyerHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := utils.GetParamInt(r, "id")
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		var dto models.BuyerUpdateDTO
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			errors.HandleError(w, errors.WrapErrBadRequest(err))
			return
		}

		if err := utils.ValidateRequestData(dto); err != nil {
			errors.HandleError(w, errors.WrapErrUnprocessableEntity(err))
			return
		}

		existingBuyer, err := h.service.GetById(id)
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		if updated := utils.UpdateFields(&existingBuyer, &dto); !updated {
			newError := fmt.Errorf("%w : no fields provided for update", errors.ErrUnprocessableEntity)
			errors.HandleError(w, newError)
			return
		}

		resp, err := h.service.Update(id, existingBuyer)
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		doc := resp.ModelToDoc()
		response.JSON(w, http.StatusOK, models.SuccessResponse{Data: doc})
	}
}
