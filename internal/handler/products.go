package handler

import (
	service "ProyectoFinal/internal/service/products"
	pkgErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"errors"
	"fmt"
	"strconv"
	"strings"

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
		newErr := pkgErrors.WrapErrBadRequest(err)
		pkgErrors.HandleError(w, newErr)
		return
	}

	if err := validate.Struct(reqBody); err != nil {
		newError := pkgErrors.WrapErrUnprocessableEntity(err)
		pkgErrors.HandleError(w, newError)
		return
	}

	model := reqBody.DocToModel()

	product, err := h.service.CreateProduct(model)

	if err != nil {
		pkgErrors.HandleError(w, err)
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
		pkgErrors.HandleError(w, err)
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
		pkgErrors.HandleError(w, err)
		return
	}
	product, err := h.service.FindProductsById(id)
	if err != nil {
		pkgErrors.HandleError(w, err)
		return
	}

	data := product.ModelToDoc()

	response.JSON(w, http.StatusOK, map[string]any{
		"message": "success",
		"data":    data,
	})
}
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkgErrors.HandleError(w, err)
		return
	}
	currentProd, err := h.service.FindProductsById(id)
	if err != nil {
		if errors.Is(err, pkgErrors.ErrNotFound) {
			newErr := fmt.Errorf("%w : product with id %d not found", pkgErrors.ErrNotFound, id)
			pkgErrors.HandleError(w, newErr)
		} else {
			pkgErrors.HandleError(w, err)
		}
		return
	}

	var reqBody models.ProductDoc
	if err := request.JSON(r, &reqBody); err != nil {
		newErr := pkgErrors.WrapErrBadRequest(err)
		pkgErrors.HandleError(w, newErr)
		return
	}

	if reqBody.ProductCode != nil && strings.TrimSpace(*reqBody.ProductCode) == "" {
		response.Error(w, http.StatusBadRequest, "'product_code' no puede ser vacío si se especifica")
		return
	}
	if reqBody.Description != nil && strings.TrimSpace(*reqBody.Description) == "" {
		response.Error(w, http.StatusBadRequest, "'description' no puede ser vacío si se especifica")
		return
	}
	if reqBody.Width != nil && *reqBody.Width <= 0 {
		response.Error(w, http.StatusBadRequest, "'width' debe ser mayor a 0 si se especifica")
		return
	}
	if reqBody.Height != nil && *reqBody.Height <= 0 {
		response.Error(w, http.StatusBadRequest, "'height' debe ser mayor a 0 si se especifica")
		return
	}
	if reqBody.Length != nil && *reqBody.Length <= 0 {
		response.Error(w, http.StatusBadRequest, "'length' debe ser mayor a 0 si se especifica")
		return
	}
	if reqBody.NetWeight != nil && *reqBody.NetWeight <= 0 {
		response.Error(w, http.StatusBadRequest, "'net_weight' debe ser mayor a 0 si se especifica")
		return
	}
	if reqBody.ExpirationRate != nil && *reqBody.ExpirationRate < 0 {
		response.Error(w, http.StatusBadRequest, "'expiration_rate' no puede ser negativo si se especifica")
		return
	}
	if reqBody.Temperature != nil && *reqBody.Temperature == 0 {
		response.Error(w, http.StatusBadRequest, "'recommended_freezing_temperature' no puede ser 0 si se especifica")
		return
	}
	if reqBody.FreezingRate != nil && *reqBody.FreezingRate < 0 {
		response.Error(w, http.StatusBadRequest, "'freezing_rate' no puede ser negativo si se especifica")
		return
	}
	if reqBody.ProductType != nil {
		if reqBody.ProductType.ID <= 0 {
			response.Error(w, http.StatusBadRequest, "'product_type.id' debe ser mayor a 0 si se especifica")
			return
		}
		if strings.TrimSpace(reqBody.ProductType.Description) == "" {
			response.Error(w, http.StatusBadRequest, "'product_type.description' no puede ser vacío si se especifica")
			return
		}
	}
	updateProd := currentProd
	update, err := h.service.UpdateProduct(id, updateProd)
	if err != nil {
		pkgErrors.HandleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, map[string]any{
		"message": "success",
		"data":    update,
	})

}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		pkgErrors.HandleError(w, err)
		return
	}
	err = h.service.DeleteProduct(id)
	if err != nil {
		pkgErrors.HandleError(w, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
