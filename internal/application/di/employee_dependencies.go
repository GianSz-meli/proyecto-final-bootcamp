package di

import (
	"ProyectoFinal/internal/handler"
	employeeRepository "ProyectoFinal/internal/repository/employee"
	employeeService "ProyectoFinal/internal/service/employee"
	"ProyectoFinal/pkg/models"
	"database/sql"
)

func GetEmployeeHandler(db *sql.DB) *handler.EmployeeHandler {
	employeeRepo := employeeRepository.NewMySQLRepository(db)
	employeeSrv := employeeService.NewService(employeeRepo)
	employeeHdl := handler.NewEmployeeHandler(employeeSrv)
	return employeeHdl
}

func GetEmployeeHandlerWithMap(db map[int]models.Employee) *handler.EmployeeHandler {
	employeeRepo := employeeRepository.NewRepository(db)
	employeeSrv := employeeService.NewService(employeeRepo)
	employeeHdl := handler.NewEmployeeHandler(employeeSrv)
	return employeeHdl
}
