package handler

import (
	"ProyectoFinal/internal/handler/utils"
	service "ProyectoFinal/internal/service/product_record"
	pkgErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"log"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
)

type ProductRecordHandler struct {
	service service.ProductRecordService
}

func NewProductRecordHandler(service service.ProductRecordService) *ProductRecordHandler {
	return &ProductRecordHandler{service: service}
}

func (h *ProductRecordHandler) CreateProductRecord(w http.ResponseWriter, r *http.Request) {
	var reqBody models.ProductRecordDoc

	if err := request.JSON(r, &reqBody); err != nil {
		newErr := pkgErrors.WrapErrBadRequest(err)
		pkgErrors.HandleError(w, newErr)
		return
	}

	if err := utils.ValidateRequestData(reqBody); err != nil {
		pkgErrors.HandleError(w, err)
		return
	}
	model := reqBody.DocToModel()
	product, err := h.service.CreateProductRecord(model)
	if err != nil {
		log.Println(err)
		pkgErrors.HandleError(w, err)
		return
	}

	body := models.SuccessResponse{Data: product.ModelToDoc()}
	response.JSON(w, http.StatusCreated, body)
}

func (h *ProductRecordHandler) GetProductRecordsCount(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	productID, err := strconv.Atoi(idStr)
	if err != nil {
		pkgErrors.HandleError(w, err)
		return
	}
	report, err := h.service.GetRecordsProduct(productID)
	if err != nil {
		log.Printf("[ProductRecordHandler][GetProductRecords] error: %v", err)
		pkgErrors.HandleError(w, err)
		return
	}

	response.JSON(w, http.StatusOK, report)
}
