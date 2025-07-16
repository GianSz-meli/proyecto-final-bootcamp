package handler

import (
	utilsHandler "ProyectoFinal/internal/handler/utils"
	"ProyectoFinal/internal/service/seller"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
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

// Create handles HTTP POST requests for creating a new seller.
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

// Update handles HTTP PUT/PATCH requests for updating an existing seller.
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

		sellerUpdated, err := h.service.Update(id, &reqBody)

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

// GetAll handles HTTP GET requests for retrieving all sellers.
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

// GetById handles HTTP GET requests for retrieving a seller by its ID.
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

// Delete handles HTTP DELETE requests for deleting a seller by its ID.
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
