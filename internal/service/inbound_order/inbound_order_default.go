package inbound_order

import (
	employeeRepository "ProyectoFinal/internal/repository/employee"
	inboundOrderRepository "ProyectoFinal/internal/repository/inbound_order"
	"ProyectoFinal/pkg/models"
)

type service struct {
	inboundOrderRepository inboundOrderRepository.Repository
	employeeRepository     employeeRepository.Repository
}

func NewService(inboundOrderRepository inboundOrderRepository.Repository, employeeRepository employeeRepository.Repository) Service {
	return &service{
		inboundOrderRepository: inboundOrderRepository,
		employeeRepository:     employeeRepository,
	}
}

func (s *service) Create(inboundOrder models.InboundOrder) (models.InboundOrder, error) {
	_, err := s.employeeRepository.GetById(inboundOrder.EmployeeID)
	if err != nil {
		return models.InboundOrder{}, err
	}

	if err := s.inboundOrderRepository.Create(&inboundOrder); err != nil {
		return models.InboundOrder{}, err
	}

	return inboundOrder, nil
}

func (s *service) GetEmployeeInboundOrdersReportByEmployeeId(employeeId int) (models.EmployeeInboundOrdersReport, error) {
	report, err := s.inboundOrderRepository.GetEmployeeInboundOrdersReportByEmployeeId(employeeId)
	if err != nil {
		return models.EmployeeInboundOrdersReport{}, err
	}

	return report, nil
}

func (s *service) GetEmployeeInboundOrdersReportAll() ([]models.EmployeeInboundOrdersReport, error) {
	reports, err := s.inboundOrderRepository.GetEmployeeInboundOrdersReportAll()
	if err != nil {
		return nil, err
	}

	return reports, nil
}
