package di

import (
	"ProyectoFinal/internal/handler"
 	prodRecordRepo "ProyectoFinal/internal/repository/product_record"
	prodRecordService "ProyectoFinal/internal/service/product_record"
	"database/sql"
)

func GetProductRecordHandler(sqlDB *sql.DB) *handler.ProductRecordHandler {
	productRecordRepository := prodRecordRepo.NewProductRecordSQL(sqlDB)
	productRecordServiceImpl := prodRecordService.NewProductRecordDefault(productRecordRepository)
	return handler.NewProductRecordHandler(productRecordServiceImpl)
}