package di

import (
	"ProyectoFinal/internal/handler"
	productbatchrepo "ProyectoFinal/internal/repository/product_batch"
	productbatchsvc "ProyectoFinal/internal/service/product_batch"
	"database/sql"
)

func GetProductBatchHandler(db *sql.DB) *handler.ProductBatchHandler {
	productBatchRepository := productbatchrepo.NewProductBatchMySQL(db)
	productBatchService := productbatchsvc.NewProductBatchService(productBatchRepository)
	productBatchHandler := handler.NewProductBatchHandler(productBatchService)
	return productBatchHandler
}
