package router

import (
	"ProyectoFinal/internal/handler"

	"github.com/go-chi/chi/v5"
)

func GetProductBatchRouter(handler *handler.ProductBatchHandler) chi.Router {
	r := chi.NewRouter()
	r.Post("/", handler.Create())
	return r
}
