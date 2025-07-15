package handler

import (
	utilsHandler "ProyectoFinal/internal/handler/utils"
	"ProyectoFinal/internal/service/locality"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"log"
	"net/http"
)

type LocalityHandler struct {
	service locality.LocalityService
}

func NewLocalityHandler(service locality.LocalityService) *LocalityHandler {
	return &LocalityHandler{service: service}
}

func (h *LocalityHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody models.LocalityDoc

		if err := request.JSON(r, &reqBody); err != nil {
			newErr := errors.WrapErrBadRequest(err)
			log.Println(newErr)
			errors.HandleError(w, newErr)
			return
		}

		if err := utilsHandler.ValidateRequestData(reqBody); err != nil {
			log.Println(err)
			errors.HandleError(w, err)
			return
		}

		model := reqBody.DocToModel()

		locality, err := h.service.Create(model)

		if err != nil {
			log.Println(err)
			errors.HandleError(w, err)
			return
		}
		body := models.SuccessResponse{Data: []models.LocalityDoc{locality.ModelToDoc()}}
		response.JSON(w, http.StatusCreated, body)

	}

}

func (h *LocalityHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := utilsHandler.GetParamInt(r, "id")
		if err != nil {
			log.Println(err)
			errors.HandleError(w, err)
			return
		}

		locality, err := h.service.GetById(id)

		if err != nil {
			log.Println(err)
			errors.HandleError(w, err)
			return
		}

		body := models.SuccessResponse{
			Data: []models.LocalityDoc{locality.ModelToDoc()},
		}
		response.JSON(w, http.StatusOK, body)
	}
}
