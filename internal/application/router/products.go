package router

import (
	"ProyectoFinal/internal/handler"

	"github.com/go-chi/chi/v5"
)

func ProductRoutes(hd *handler.ProductHandler) chi.Router {
	r := chi.NewRouter()
	r.Post("/", hd.CreateProduct)
	r.Get("/", hd.FindAllProducts)
	return r
}
