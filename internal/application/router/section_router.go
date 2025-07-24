package router

import (
	"ProyectoFinal/internal/handler"
	sectionHandler "ProyectoFinal/internal/handler/section"

	"github.com/go-chi/chi/v5"
)

func GetSectionRouter(handler *sectionHandler.SectionDefault, productBatchHandler *handler.ProductBatchHandler) chi.Router {
	r := chi.NewRouter()
	r.Post("/", handler.Create())
	r.Get("/", handler.GetAll())

	// Report endpoints - must go BEFORE routes with parameters
	r.Get("/reportProducts", productBatchHandler.GetProductCountByAllSections())
	r.Get("/reportProducts/{id}", productBatchHandler.GetProductCountBySection())

	r.Get("/{id}", handler.GetById())
	r.Delete("/{id}", handler.Delete())
	r.Patch("/{id}", handler.Update())

	return r
}
