package router

import (
	"ProyectoFinal/internal/handler"

	"github.com/go-chi/chi/v5"
)

func GetCarrierRouter(h *handler.CarrierHandler) chi.Router {
	rt := chi.NewRouter()

	rt.Post("/", h.Create())
	return rt
}
