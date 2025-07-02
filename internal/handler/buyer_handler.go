package handler

import (
	"ProyectoFinal/internal/handler/utils"
	"ProyectoFinal/internal/service/buyer"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"strconv"
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

		newResp := models.SuccessResponse{
			Data: resp,
		}

		response.JSON(w, http.StatusOK, newResp)
	}
}

func (h *BuyerHandler) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		var dto models.BuyerCreateDTO
		if err := json.Unmarshal(body, &dto); err != nil {
			errors.HandleError(w, errors.ErrBadRequest)
			return
		}

		if err := utils.ValidateRequestData(dto); err != nil {
			errors.HandleError(w, errors.WrapErrUnprocessableEntity(err))
			return
		}

		resp, err := h.service.Save(dto)
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		newResp := models.SuccessResponse{
			Data: resp,
		}

		response.JSON(w, http.StatusCreated, newResp)
	}
}

func (h *BuyerHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		if intId <= 0 {
			errors.HandleError(w, errors.WrapErrBadRequest(fmt.Errorf("id must be greather than 0")))
			return
		}

		resp, err := h.service.GetById(intId)
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		newResp := models.SuccessResponse{
			Data: resp,
		}

		response.JSON(w, http.StatusOK, newResp)
	}
}

func (h *BuyerHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		if intId <= 0 {
			errors.HandleError(w, errors.WrapErrBadRequest(fmt.Errorf("id must be greather than 0")))
			return
		}

		if err := h.service.Delete(intId); err != nil {
			errors.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func (h *BuyerHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		intId, err := strconv.Atoi(id)
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		if intId <= 0 {
			errors.HandleError(w, errors.WrapErrBadRequest(fmt.Errorf("id must be greather than 0")))
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		var dto models.BuyerUpdateDTO
		if err := json.Unmarshal(body, &dto); err != nil {
			errors.HandleError(w, errors.ErrBadRequest)
			return
		}

		resp, err := h.service.Update(intId, dto)
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		newResp := models.SuccessResponse{
			Data: resp,
		}
		response.JSON(w, http.StatusOK, newResp)
	}
}
