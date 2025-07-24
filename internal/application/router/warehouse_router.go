package router

import (
	handler "ProyectoFinal/internal/handler/warehouse"

	"github.com/go-chi/chi/v5"
)

func GetWarehouseRouter(handler *handler.WarehouseHandler) chi.Router {
	rt := chi.NewRouter()

	rt.Get("/", handler.GetAllWarehouses)
	rt.Get("/{id}", handler.GetWarehouseById)
	rt.Post("/", handler.CreateWarehouse)
	rt.Patch("/{id}", handler.UpdateWarehouse)
	rt.Delete("/{id}", handler.DeleteWarehouse)

	return rt
}
