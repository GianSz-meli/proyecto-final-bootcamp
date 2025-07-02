package router

import (
	"ProyectoFinal/internal/handler"
	"github.com/go-chi/chi/v5"
)

func GetBuyerRouter(h *handler.BuyerHandler) chi.Router {
	rt := chi.NewRouter()

	rt.Get("/", h.GetAll())
	rt.Get("/{id}", h.GetById())
	rt.Post("/", h.Save())
	rt.Patch("/{id}", h.Update())
	rt.Delete("/{id}", h.Delete())
	return rt
}
