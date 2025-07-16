package handler

import (
	utilsHandler "ProyectoFinal/internal/handler/utils"
	"ProyectoFinal/internal/service/seller"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"fmt"
	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"log"
	"net/http"
)

type SellerHandler struct {
	service seller.SellerService
}

func NewSellerHandler(service seller.SellerService) *SellerHandler {
	return &SellerHandler{service: service}
}

func (h *SellerHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody models.CreateSellerRequest

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

		seller, err := h.service.Create(model)

		if err != nil {
			log.Println(err)
			errors.HandleError(w, err)
			return
		}
		body := models.SuccessResponse{Data: []models.SellerDoc{seller.ModelToDoc()}}
		response.JSON(w, http.StatusCreated, body)

	}

}

func (h *SellerHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := utilsHandler.GetParamInt(r, "id")

		if err != nil {
			log.Println(err)
			errors.HandleError(w, err)
			return
		}

		var reqBody models.UpdateSellerRequest

		if err = request.JSON(r, &reqBody); err != nil {
			newErr := errors.WrapErrBadRequest(err)
			log.Println(newErr)
			errors.HandleError(w, newErr)
			return
		}

		if err = utilsHandler.ValidateRequestData(reqBody); err != nil {
			log.Println(err)
			errors.HandleError(w, err)
			return
		}

		sellerToUpdate, err := h.service.GetById(id)

		if err != nil {
			log.Println(err)
			errors.HandleError(w, err)
			return
		}

		if updated := utilsHandler.UpdateFields(&sellerToUpdate, &reqBody); !updated {
			newError := fmt.Errorf("%w : no fields provided for update", errors.ErrUnprocessableEntity)
			log.Println(newError)
			errors.HandleError(w, newError)
			return
		}

		sellerUpdated, err := h.service.Update(sellerToUpdate)

		if err != nil {
			log.Println(err)
			errors.HandleError(w, err)
			return
		}

		body := models.SuccessResponse{
			Data: []models.SellerDoc{sellerUpdated.ModelToDoc()},
		}

		response.JSON(w, http.StatusOK, body)

	}
}

func (h *SellerHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		sellers, err := h.service.GetAll()

		if err != nil {
			log.Println(err)
			errors.HandleError(w, err)
			return
		}

		var sellersDoc []models.SellerDoc

		for _, seller := range sellers {
			sellersDoc = append(sellersDoc, seller.ModelToDoc())
		}

		body := models.SuccessResponse{
			Data: sellersDoc,
		}
		response.JSON(w, http.StatusOK, body)
	}
}

func (h *SellerHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := utilsHandler.GetParamInt(r, "id")
		if err != nil {
			log.Println(err)
			errors.HandleError(w, err)
			return
		}

		seller, err := h.service.GetById(id)

		if err != nil {
			log.Println(err)
			errors.HandleError(w, err)
			return
		}

		body := models.SuccessResponse{
			Data: []models.SellerDoc{seller.ModelToDoc()},
		}
		response.JSON(w, http.StatusOK, body)
	}
}

func (h *SellerHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := utilsHandler.GetParamInt(r, "id")
		if err != nil {
			log.Println(err)
			errors.HandleError(w, err)
			return
		}

		err = h.service.Delete(id)

		if err != nil {
			log.Println(err)
			errors.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
