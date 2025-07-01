package di

import (
	"ProyectoFinal/internal/handler"
	employeeRepository "ProyectoFinal/internal/repository/employee"
	employeeService "ProyectoFinal/internal/service/employee"
	employeemodel "ProyectoFinal/pkg/models/employee"
)

func GetEmployeeHandler(db map[int]employeemodel.Employee) *handler.EmployeeHandler {
	employeeRepo := employeeRepository.NewRepository(db)
	employeeSrv := employeeService.NewService(employeeRepo)
	employeeHdl := handler.NewEmployeeHandler(employeeSrv)
	return employeeHdl
}
