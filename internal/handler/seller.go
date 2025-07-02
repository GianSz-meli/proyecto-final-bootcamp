package handler

import (
	"ProyectoFinal/internal/service/seller"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"fmt"
	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

type SellerHandler struct {
	service seller.SellerService
}

func NewSellerHandler(service seller.SellerService) *SellerHandler {
	return &SellerHandler{service: service}
}

var validate = validator.New()

func (h *SellerHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqBody models.SellerDoc

		if err := request.JSON(r, &reqBody); err != nil {
			newErr := errors.WrapErrBadRequest(err)
			errors.HandleError(w, newErr)
			return
		}

		if err := validate.Struct(reqBody); err != nil {
			newError := fmt.Errorf("%w : %s", errors.ErrUnprocessableEntity, err.Error())
			errors.HandleError(w, newError)
			return
		}

		model := reqBody.DocToModel()

		seller, err := h.service.Create(model)

		if err != nil {
			errors.HandleError(w, err)
			return
		}

		body := map[string][]models.SellerDoc{
			"data": {seller.ModelToDoc()},
		}
		response.JSON(w, http.StatusCreated, body)

	}

}

func (h *SellerHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		sellers := h.service.GetAll()

		var sellersDoc []models.SellerDoc

		for _, seller := range sellers {
			sellersDoc = append(sellersDoc, seller.ModelToDoc())
		}

		body := map[string][]models.SellerDoc{
			"data": sellersDoc,
		}
		response.JSON(w, http.StatusOK, body)
	}
}

func (h *SellerHandler) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqId := chi.URLParam(r, "id")

		id, err := strconv.Atoi(reqId)

		if err != nil {
			errors.HandleError(w, errors.ErrBadRequest)
			return
		}

		seller, err := h.service.GetById(id)

		if err != nil {
			errors.HandleError(w, err)
			return
		}

		body := map[string][]models.SellerDoc{
			"data": {seller.ModelToDoc()},
		}

		response.JSON(w, http.StatusOK, body)
	}
}
