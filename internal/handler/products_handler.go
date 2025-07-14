package handler

import (
	"ProyectoFinal/internal/handler/utils"
	service "ProyectoFinal/internal/service/products"
	pkgErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"fmt"
	"log"
	"net/http"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
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

	if err := utils.ValidateRequestData(reqBody); err != nil {
		pkgErrors.HandleError(w, err)
		return
	}
	println("antes doctomodel")
	model := reqBody.DocToModel()
	println("despues doctomodel")
	product, err := h.service.CreateProduct(model)
	if err != nil {
		log.Println(err)
		pkgErrors.HandleError(w, err)
		return
	}

	body := models.SuccessResponse{Data: product.ModelToDoc()}
	response.JSON(w, http.StatusCreated, body)
}

func (h *ProductHandler) FindAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.FindAllProducts()
	if err != nil {
		log.Println(err)
		pkgErrors.HandleError(w, err)
		return
	}
	data := make(map[int]models.ProductDoc)
	for key, value := range products {
		data[key] = value.ModelToDoc()
	}
	body := models.SuccessResponse{Data: data}
	response.JSON(w, http.StatusOK, body)
}

func (h *ProductHandler) FindProductsById(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetParamInt(r, "id")
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
	body := models.SuccessResponse{Data: data}
	response.JSON(w, http.StatusOK, body)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetParamInt(r, "id")
	if err != nil {
		log.Println(err)
		pkgErrors.HandleError(w, err)
		return
	}

	currentProd, err := h.service.FindProductsById(id)
	if err != nil {
		log.Println(err)
		pkgErrors.HandleError(w, err)
		return
	}

	var reqBody models.ProductDocUpdate
	if err := request.JSON(r, &reqBody); err != nil {
		newErr := pkgErrors.WrapErrBadRequest(err)
		pkgErrors.HandleError(w, newErr)
		return
	}

	if err := utils.ValidateRequestData(reqBody); err != nil {
		pkgErrors.HandleError(w, err)
		return
	}

	if updated := utils.UpdateFields(&currentProd, &reqBody); !updated {
		newError := fmt.Errorf("%w : no fields provided for update", pkgErrors.ErrUnprocessableEntity)
		pkgErrors.HandleError(w, newError)
		return
	}
	
	update, err := h.service.UpdateProduct(id, currentProd)
	if err != nil {
		log.Println(err)
		pkgErrors.HandleError(w, err)
		return
	}
	body := models.SuccessResponse{Data: update}
	response.JSON(w, http.StatusCreated, body)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := utils.GetParamInt(r, "id")
	if err != nil {
		log.Println(err)
		pkgErrors.HandleError(w, err)
		return
	}
	err = h.service.DeleteProduct(id)
	if err != nil {
		log.Println(err)
		pkgErrors.HandleError(w, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
