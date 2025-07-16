package router

import (
	"ProyectoFinal/internal/handler"

	"github.com/go-chi/chi/v5"
)

func GetProductRecordRouter(h *handler.ProductRecordHandler) chi.Router {
	rt := chi.NewRouter()

	rt.Post("/", h.CreateProductRecord)
	return rt
}