package di

import (
	"ProyectoFinal/internal/handler/product_record"
	prodRecordRepo "ProyectoFinal/internal/repository/product_record"
	prodRecordService "ProyectoFinal/internal/service/product_record"
	"database/sql"
)

func GetProductRecordHandler(sqlDB *sql.DB) *product_record.ProductRecordHandler {
	productRecordRepository := prodRecordRepo.NewProductRecordSQL(sqlDB)
	productRecordServiceImpl := prodRecordService.NewProductRecordDefault(productRecordRepository)
	return product_record.NewProductRecordHandler(productRecordServiceImpl)
}
