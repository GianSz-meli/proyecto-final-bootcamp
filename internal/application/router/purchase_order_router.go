package router

import (
	"ProyectoFinal/internal/handler"
	"github.com/go-chi/chi/v5"
)

func GetPurchaseOrderRouter(h *handler.PurchaseOrderHandler) chi.Router {
	rt := chi.NewRouter()

	rt.Get("/getByBuyerId/{id}", h.GetByBuyerId())
	rt.Post("/", h.Create())

	return rt
}
