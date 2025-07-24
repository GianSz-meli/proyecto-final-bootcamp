package router

import (
	"ProyectoFinal/internal/handler/purchase_order"
	"github.com/go-chi/chi/v5"
)

// GetPurchaseOrderRouter creates and configures a Chi router with purchase order-related HTTP routes.
func GetPurchaseOrderRouter(h *purchase_order.PurchaseOrderHandler) chi.Router {
	rt := chi.NewRouter()

	rt.Post("/", h.Create())

	return rt
}
