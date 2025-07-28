package router

import (
	"ProyectoFinal/internal/handler"
	employeeHandler "ProyectoFinal/internal/handler/employee"

	"github.com/go-chi/chi/v5"
)

func EmployeeRoutes(ctr *employeeHandler.EmployeeHandler, inboundOrderHandler *handler.InboundOrderHandler) chi.Router {
	r := chi.NewRouter()
	r.Get("/", ctr.GetAll())
	r.Get("/{id}", ctr.GetById())
	r.Post("/", ctr.Create())
	r.Patch("/{id}", ctr.Update())
	r.Delete("/{id}", ctr.Delete())
	r.Get("/reportInboundOrders", inboundOrderHandler.GetEmployeeInboundOrdersReport())
	return r
}
