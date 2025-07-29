package inbound_order

import (
	"ProyectoFinal/pkg/models"

	"github.com/stretchr/testify/mock"
)

type MockInboundOrderService struct {
	mock.Mock
}

func (m *MockInboundOrderService) Create(inboundOrder models.InboundOrder) (models.InboundOrder, error) {
	args := m.Called(inboundOrder)
	return args.Get(0).(models.InboundOrder), args.Error(1)
}

func (m *MockInboundOrderService) GetEmployeeInboundOrdersReportByEmployeeId(employeeId int) (models.EmployeeInboundOrdersReport, error) {
	args := m.Called(employeeId)
	return args.Get(0).(models.EmployeeInboundOrdersReport), args.Error(1)
}

func (m *MockInboundOrderService) GetEmployeeInboundOrdersReportAll() ([]models.EmployeeInboundOrdersReport, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.EmployeeInboundOrdersReport), args.Error(1)
}
