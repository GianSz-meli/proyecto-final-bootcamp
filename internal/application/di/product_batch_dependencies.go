package di

import (
	"ProyectoFinal/internal/handler/product_batch"
	productbatchrepo "ProyectoFinal/internal/repository/product_batch"
	productbatchsvc "ProyectoFinal/internal/service/product_batch"
	"database/sql"
)

func GetProductBatchHandler(db *sql.DB) *product_batch.ProductBatchHandler {
	productBatchRepository := productbatchrepo.NewProductBatchMySQL(db)
	productBatchService := productbatchsvc.NewProductBatchService(productBatchRepository)
	productBatchHandler := product_batch.NewProductBatchHandler(productBatchService)
	return productBatchHandler
}
