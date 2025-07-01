package router

import (
	"ProyectoFinal/internal/handler"
	"github.com/go-chi/chi/v5"
)

func SellerRoutes(ctr *handler.SellerHandler) chi.Router {
	r := chi.NewRouter()
	r.Post("/", ctr.Create())
	r.Get("/", ctr.GetAll())
	return r
}
