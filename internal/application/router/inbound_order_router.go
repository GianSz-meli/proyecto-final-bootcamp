package router

import (
	"ProyectoFinal/internal/handler"

	"github.com/go-chi/chi/v5"
)

func InboundOrderRoutes(ctr *handler.InboundOrderHandler) chi.Router {
	r := chi.NewRouter()
	r.Post("/", ctr.Create())
	return r
}
