package inbound_order

import (
	"ProyectoFinal/pkg/models"
)

type Repository interface {
	Create(inboundOrder *models.InboundOrder) error
	GetEmployeeInboundOrdersReportByEmployeeId(employeeId int) (models.EmployeeInboundOrdersReport, error)
	GetEmployeeInboundOrdersReportAll() ([]models.EmployeeInboundOrdersReport, error)
}
