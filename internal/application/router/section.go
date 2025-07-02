package router

import (
	"ProyectoFinal/internal/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SectionRoutes(sh *handler.SectionDefault) http.Handler {
	r := chi.NewRouter()

	r.Get("/", sh.GetAll())
	r.Get("/{id}", sh.GetByID())
	r.Post("/", sh.Create())
	r.Patch("/{id}", sh.Update())
	r.Delete("/{id}", sh.Delete())

	return r
}
