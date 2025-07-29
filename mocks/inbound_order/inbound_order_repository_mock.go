package inbound_order

import (
	"ProyectoFinal/pkg/models"

	"github.com/stretchr/testify/mock"
)

type MockInboundOrderRepository struct {
	mock.Mock
}

func (m *MockInboundOrderRepository) Create(inboundOrder *models.InboundOrder) error {
	args := m.Called(inboundOrder)
	return args.Error(0)
}

func (m *MockInboundOrderRepository) GetEmployeeInboundOrdersReportByEmployeeId(employeeId int) (models.EmployeeInboundOrdersReport, error) {
	args := m.Called(employeeId)
	return args.Get(0).(models.EmployeeInboundOrdersReport), args.Error(1)
}

func (m *MockInboundOrderRepository) GetEmployeeInboundOrdersReportAll() ([]models.EmployeeInboundOrdersReport, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.EmployeeInboundOrdersReport), args.Error(1)
}
