package handler

import (
	"ProyectoFinal/internal/service/buyer"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
	"strconv"
)

type BuyerHandler struct {
	service   buyer.Service
	validator *validator.Validate
}

func NewBuyerHandler(newService buyer.Service) *BuyerHandler {
	return &BuyerHandler{
		service:   newService,
		validator: validator.New(),
	}
}

func (h *BuyerHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := h.service.GetAll()

		if err != nil {
			errors.HandleError(w, err)
			return
		}

		newResp := models.SuccessResponse{
			Data: resp,
		}

		writeJSON(w, http.StatusOK, newResp)
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
			errors.HandleError(w, err)
			return
		}

		if err := h.validator.Struct(dto); err != nil {
			errors.HandleError(w, errors.ErrUnprocessableEntity)
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

		writeJSON(w, http.StatusCreated, newResp)
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

		resp, err := h.service.GetById(intId)
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		newResp := models.SuccessResponse{
			Data: resp,
		}

		writeJSON(w, http.StatusOK, newResp)
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

		resp, err := h.service.GetById(intId)
		if err != nil {
			errors.HandleError(w, err)
			return
		}

		newResp := models.SuccessResponse{
			Data: resp,
		}

		writeJSON(w, http.StatusOK, newResp)

	}
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	jsonResp, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		errors.HandleError(w, err)
		return
	}
	w.Write(jsonResp)
}
