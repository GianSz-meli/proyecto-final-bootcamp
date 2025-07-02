package handler

import (
	service "ProyectoFinal/internal/service/products"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"strconv"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"

	//"github.com/go-playground/validator/v10"
	"net/http"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var reqBody models.ProductDoc

	if err := request.JSON(r, &reqBody); err != nil {
		newErr := errors.WrapErrBadRequest(err)
		errors.HandleError(w, newErr)
		return
	}

	if err := validate.Struct(reqBody); err != nil {
		newError := errors.WrapErrBadRequest(err)
		errors.HandleError(w, newError)
		return
	}

	model := reqBody.DocToModel()

	product, err := h.service.CreateProduct(model)

	if err != nil {
		errors.HandleError(w, err)
		return
	}

	body := map[string]models.ProductDoc{
		"data": product.ModelToDoc(),
	}
	response.JSON(w, http.StatusCreated, body)

}

func (h *ProductHandler) FindAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.FindAllProducts()
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, nil)
		return
	}
	data := make(map[int]models.ProductDoc)
	for key, value := range products {
		data[key] = value.ModelToDoc()
	}

	response.JSON(w, http.StatusOK, map[string]any{
		"message": "success",
		"data":    data,
	})
}

func (h *ProductHandler) FindProductsById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "id inválido")
		return
	}
	product, err := h.service.FindProductsById(id)
	if err != nil {
		response.Error(w, http.StatusNotFound, "producto no encontrado")
		return
	}
// manejar los errores 
	data := product.ModelToDoc()

	response.JSON(w, http.StatusOK, map[string]any{
		"message": "success",
		"data":    data,
	})
}