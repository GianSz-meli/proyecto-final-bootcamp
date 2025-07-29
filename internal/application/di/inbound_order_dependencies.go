package di

import (
	"ProyectoFinal/internal/handler/inbound_order"
	employeeRepository "ProyectoFinal/internal/repository/employee"
	inboundOrderRepository "ProyectoFinal/internal/repository/inbound_order"
	inboundOrderService "ProyectoFinal/internal/service/inbound_order"
	"database/sql"
)

func GetInboundOrderHandler(db *sql.DB) *inbound_order.InboundOrderHandler {
	inboundOrderRepo := inboundOrderRepository.NewMySQLRepository(db)
	employeeRepo := employeeRepository.NewMySQLRepository(db)
	inboundOrderSrv := inboundOrderService.NewService(inboundOrderRepo, employeeRepo)
	inboundOrderHdl := inbound_order.NewInboundOrderHandler(inboundOrderSrv)
	return inboundOrderHdl
}
