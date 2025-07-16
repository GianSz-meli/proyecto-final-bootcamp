package inbound_order

import (
	"ProyectoFinal/pkg/models"
)

type Service interface {
	Create(inboundOrder models.InboundOrder) (models.InboundOrder, error)
	GetEmployeeInboundOrdersReportByEmployeeId(employeeId int) (models.EmployeeInboundOrdersReport, error)
	GetEmployeeInboundOrdersReportAll() ([]models.EmployeeInboundOrdersReport, error)
}
