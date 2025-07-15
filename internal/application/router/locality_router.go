package router

import (
	"ProyectoFinal/internal/handler"
	"github.com/go-chi/chi/v5"
)

func GetLocalityRouter(handler *handler.LocalityHandler) chi.Router {
	r := chi.NewRouter()
	r.Post("/", handler.Create())
	r.Get("/{id}", handler.GetById())
	return r
}
