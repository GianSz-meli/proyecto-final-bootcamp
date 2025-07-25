package di

import (
	"ProyectoFinal/internal/handler/purchase_order"
	repository "ProyectoFinal/internal/repository/purchase_order"
	service "ProyectoFinal/internal/service/purchase_order"
	"database/sql"
)

// GetPurchaseOrderHandler creates and configures a complete purchase order handler with all its dependencies.
func GetPurchaseOrderHandler(db *sql.DB) *purchase_order.PurchaseOrderHandler {
	poRepository := repository.NewPurchaseOrderMySqlRepository(db)
	poService := service.NewPurchaseOrderService(poRepository)
	return purchase_order.NewpPurchaseOrderHandler(poService)
}
