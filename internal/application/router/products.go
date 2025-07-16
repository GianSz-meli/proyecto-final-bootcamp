package router

import (
	"ProyectoFinal/internal/handler"

	"github.com/go-chi/chi/v5"
)

func ProductRoutes(hd *handler.ProductHandler, getProdRecords *handler.ProductRecordHandler) chi.Router {
	r := chi.NewRouter()
	r.Post("/", hd.CreateProduct)
	r.Get("/", hd.FindAllProducts)
	r.Get("/{id}", hd.FindProductsById)
	r.Patch("/{id}", hd.UpdateProduct)
	r.Delete("/{id}", hd.DeleteProduct)

	r.Get("/reportRecords", getProdRecords.GetProductRecordsCount)

	return r
}
