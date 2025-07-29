package router

import (
	"ProyectoFinal/internal/handler/inbound_order"

	"github.com/go-chi/chi/v5"
)

func InboundOrderRoutes(ctr *inbound_order.InboundOrderHandler) chi.Router {
	r := chi.NewRouter()
	r.Post("/", ctr.Create())
	return r
}
