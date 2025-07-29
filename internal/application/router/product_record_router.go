package router

import (
	"ProyectoFinal/internal/handler/product_record"

	"github.com/go-chi/chi/v5"
)

func GetProductRecordRouter(h *product_record.ProductRecordHandler) chi.Router {
	rt := chi.NewRouter()

	rt.Post("/", h.CreateProductRecord)
	return rt
}
