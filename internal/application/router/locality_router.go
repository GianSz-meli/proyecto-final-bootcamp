package router

import (
	"ProyectoFinal/internal/handler/locality"
	"github.com/go-chi/chi/v5"
)

// GetLocalityRouter initializes and returns a Chi router with locality-related routes.
func GetLocalityRouter(handler *locality.LocalityHandler) chi.Router {
	r := chi.NewRouter()
	r.Post("/", handler.Create())
	r.Get("/reportSellers", handler.GetSellersByLocality())
	r.Get("/reportCarriers", handler.ReportCarriersByLocality())
	return r
}
