package router

import (
	"ProyectoFinal/internal/handler"
	"github.com/go-chi/chi/v5"
)

func GetBuyerRouter(h *handler.BuyerHandler) chi.Router {
	rt := chi.NewRouter()

	rt.Get("/", h.GetAll())
	rt.Get("/{id}", h.GetById())
	rt.Post("/", h.Create())
	rt.Patch("/{id}", h.Update())
	rt.Delete("/{id}", h.Delete())
	rt.Get("/reportPurchaseOrders", h.GetAllWithOrderCount())
	rt.Get("/reportPurchaseOrders/{id}", h.GetByIdWithOrderCount())
	return rt
}
