package di

import (
	"ProyectoFinal/internal/handler"
	employeeRepository "ProyectoFinal/internal/repository/employee"
	employeeService "ProyectoFinal/internal/service/employee"
	"ProyectoFinal/pkg/models"
)

func GetEmployeeHandler(db map[int]models.Employee) *handler.EmployeeHandler {
	employeeRepo := employeeRepository.NewRepository(db)
	employeeSrv := employeeService.NewService(employeeRepo)
	employeeHdl := handler.NewEmployeeHandler(employeeSrv)
	return employeeHdl
}
