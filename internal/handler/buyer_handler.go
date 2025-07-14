package handler

import (
	"ProyectoFinal/internal/handler/utils"
	"ProyectoFinal/internal/service/buyer"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"fmt"
	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"log"
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

		buyerToUpdate, existsErr := h.service.GetById(id)
		if existsErr != nil {
			log.Println(existsErr)
			errors.HandleError(w, existsErr)
			return
		}

		if updated := utils.UpdateFields(&buyerToUpdate, &dto); !updated {
			newError := fmt.Errorf("%w : no fields provided for update", errors.ErrUnprocessableEntity)
			log.Println(newError)
			errors.HandleError(w, newError)
			return
		}

		resp, updateErr := h.service.Update(id, buyerToUpdate)
		if updateErr != nil {
			log.Println(updateErr)
			errors.HandleError(w, updateErr)
			return
		}

		response.JSON(w, http.StatusOK, models.SuccessResponse{Data: resp.ModelToDoc()})
	}
}

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
