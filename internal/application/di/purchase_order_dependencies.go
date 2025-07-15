package di

import (
	"ProyectoFinal/internal/handler"
	repository "ProyectoFinal/internal/repository/purchase_order"
	service "ProyectoFinal/internal/service/purchase_order"
	"database/sql"
)

func GetPurchaseOrderHandler(db *sql.DB) *handler.PurchaseOrderHandler {
	poRepository := repository.NewPurchaseOrderMySqlRepository(db)
	poService := service.NewPurchaseOrderService(poRepository)
	return handler.NewpPurchaseOrderHandler(poService)
}
