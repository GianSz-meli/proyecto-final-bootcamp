package router

import (
	"ProyectoFinal/internal/handler"
	"github.com/go-chi/chi/v5"
)

// GetLocalityRouter initializes and returns a Chi router with locality-related routes.
func GetLocalityRouter(handler *handler.LocalityHandler) chi.Router {
	r := chi.NewRouter()
	r.Post("/", handler.Create())
	r.Get("/{id}", handler.GetById())
	r.Get("/reportSellers", handler.GetSellersByLocality())
	return r
}
