package router

import (
	employeeHandler "ProyectoFinal/internal/handler/employee"

	"github.com/go-chi/chi/v5"
)

func EmployeeRoutes(ctr *employeeHandler.EmployeeHandler) chi.Router {
	r := chi.NewRouter()
	r.Get("/", ctr.GetAll())
	r.Get("/{id}", ctr.GetById())
	r.Post("/", ctr.Create())
	r.Patch("/{id}", ctr.Update())
	r.Delete("/{id}", ctr.Delete())
	return r
}
