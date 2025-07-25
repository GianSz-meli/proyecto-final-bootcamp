package router

import (
	"ProyectoFinal/internal/handler/seller"

	"github.com/go-chi/chi/v5"
)

// GetSellerRouter initializes and returns a Chi router with routes for seller resource operations.
func GetSellerRouter(handler *seller.SellerHandler) chi.Router {
	r := chi.NewRouter()
	r.Post("/", handler.Create())
	r.Get("/", handler.GetAll())
	r.Get("/{id}", handler.GetById())
	r.Delete("/{id}", handler.Delete())
	r.Patch("/{id}", handler.Update())
	return r
}
