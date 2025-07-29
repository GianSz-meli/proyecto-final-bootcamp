package router

import (
	"ProyectoFinal/internal/handler/product_batch"

	"github.com/go-chi/chi/v5"
)

func GetProductBatchRouter(handler *product_batch.ProductBatchHandler) chi.Router {
	r := chi.NewRouter()
	r.Post("/", handler.Create())
	return r
}
