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
	"strconv"
)

type LocalityHandler struct {
	service locality.LocalityService
}

func NewLocalityHandler(service locality.LocalityService) *LocalityHandler {
	return &LocalityHandler{service: service}
}

// Create handles the HTTP request for creating a new locality.
func (h *LocalityHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody models.LocalityCreateRequest

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
		body := models.SuccessResponse{Data: []models.LocalityCreateRequest{locality.ModelToDoc()}}
		response.JSON(w, http.StatusCreated, body)

	}

}

// GetById handles the HTTP request for retrieving a locality by its ID.
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
			Data: []models.Locality{locality},
		}
		response.JSON(w, http.StatusOK, body)
	}
}

// GetSellersByLocality handles the HTTP request for retrieving seller info by locality ID.
// If the 'id' query parameter is not present, it returns sellers for all localities.
func (h *LocalityHandler) GetSellersByLocality() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idReq := r.URL.Query().Get("id")
		if idReq == "" {
			h.GetSellersByLocalities(w, r)
			return
		}
		id, err := strconv.Atoi(idReq)

		if err != nil {
			log.Println(err)
			errors.HandleError(w, err)
			return
		}

		sellersByLocality, err := h.service.GetSellersByIdLocality(id)

		if err != nil {
			log.Println(err)
			errors.HandleError(w, err)
			return
		}

		body := models.SuccessResponse{
			Data: sellersByLocality,
		}
		response.JSON(w, http.StatusOK, body)
	}
}

// GetSellersByLocalities handles the HTTP request for retrieving seller info for all localities.
func (h *LocalityHandler) GetSellersByLocalities(w http.ResponseWriter, r *http.Request) {
	sellersByLocalities, err := h.service.GetSellersByLocalities()

	if err != nil {
		log.Println(err)
		errors.HandleError(w, err)
		return
	}

	body := models.SuccessResponse{
		Data: sellersByLocalities,
	}
	response.JSON(w, http.StatusOK, body)
}

// ReportCarriersByLocality handles the HTTP request for retrieving carrier info 
// by locality ID or by all localities if ID sent on the query params is nil.
func (h *LocalityHandler) ReportCarriersByLocality() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := utilsHandler.GetQueryInt(r, "id")
		if err != nil {
			errors.HandleError(w, errors.WrapErrBadRequest(err))
			return
		}

		report, serviceErr := h.service.ReportCarriersByLocality(id)
		if serviceErr != nil {
			errors.HandleError(w, serviceErr)
			return
		}

		response.JSON(w, http.StatusOK, models.SuccessResponse{Data: report})
	}
}