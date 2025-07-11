package handler

import (
	utilsHandler "ProyectoFinal/internal/handler/utils"
	"ProyectoFinal/internal/service/seller"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"fmt"
	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
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
			errors.HandleError(w, newErr)
			return
		}

		if err := utilsHandler.ValidateRequestData(reqBody); err != nil {
			errors.HandleError(w, err)
			return
		}

		model := reqBody.DocToModel()

		seller, err := h.service.Create(model)

		if err != nil {
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
			errors.HandleError(w, err)
			return
		}

		var reqBody models.UpdateSellerRequest

		if err = request.JSON(r, &reqBody); err != nil {
			newErr := errors.WrapErrBadRequest(err)
			errors.HandleError(w, newErr)
			return
		}

		if err := utilsHandler.ValidateRequestData(reqBody); err != nil {
			errors.HandleError(w, err)
			return
		}

		sellerToUpdate, err := h.service.GetById(id)

		if err != nil {
			newError := errors.WrapErrNotFound("seller", "id", id)
			errors.HandleError(w, newError)
			return
		}

		if updated := utilsHandler.UpdateFields(&sellerToUpdate, &reqBody); !updated {
			newError := fmt.Errorf("%w : no fields provided for update", errors.ErrUnprocessableEntity)
			errors.HandleError(w, newError)
			return
		}

		sellerUpdated, err := h.service.Update(id, sellerToUpdate)

		if err != nil {
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

		sellers := h.service.GetAll()

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
			errors.HandleError(w, err)
			return
		}

		seller, err := h.service.GetById(id)

		if err != nil {
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
			errors.HandleError(w, err)
			return
		}

		err = h.service.Delete(id)

		if err != nil {
			errors.HandleError(w, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
