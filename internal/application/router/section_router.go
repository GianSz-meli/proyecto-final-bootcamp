package router

import (
	"ProyectoFinal/internal/handler"

	"github.com/go-chi/chi/v5"
)

func GetSectionRouter(handler *handler.SectionDefault) chi.Router {
	r := chi.NewRouter()
	r.Post("/", handler.Create())
	r.Get("/", handler.GetAll())
	r.Get("/{id}", handler.GetById())
	r.Delete("/{id}", handler.Delete())
	r.Patch("/{id}", handler.Update())
	return r
}
