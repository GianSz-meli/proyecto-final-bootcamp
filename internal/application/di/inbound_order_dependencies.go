package di

import (
	"ProyectoFinal/internal/handler"
	employeeRepository "ProyectoFinal/internal/repository/employee"
	inboundOrderRepository "ProyectoFinal/internal/repository/inbound_order"
	inboundOrderService "ProyectoFinal/internal/service/inbound_order"
	"database/sql"
)

func GetInboundOrderHandler(db *sql.DB) *handler.InboundOrderHandler {
	inboundOrderRepo := inboundOrderRepository.NewMySQLRepository(db)
	employeeRepo := employeeRepository.NewMySQLRepository(db)
	inboundOrderSrv := inboundOrderService.NewService(inboundOrderRepo, employeeRepo)
	inboundOrderHdl := handler.NewInboundOrderHandler(inboundOrderSrv)
	return inboundOrderHdl
}
