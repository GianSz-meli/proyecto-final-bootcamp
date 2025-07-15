package inbound_order

import (
	"ProyectoFinal/pkg/models"
)

type Repository interface {
	Create(inboundOrder *models.InboundOrder) error
	ExistsByOrderNumber(orderNumber string) (bool, error)
	GetEmployeeInboundOrdersReportByEmployeeId(employeeId int) (models.EmployeeInboundOrdersReport, error)
	GetEmployeeInboundOrdersReportAll() ([]models.EmployeeInboundOrdersReport, error)
}
