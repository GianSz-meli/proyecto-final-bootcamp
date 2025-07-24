package di

import (
	employeeHandler "ProyectoFinal/internal/handler/employee"
	employeeRepository "ProyectoFinal/internal/repository/employee"
	employeeService "ProyectoFinal/internal/service/employee"
	"database/sql"
)

func GetEmployeeHandler(db *sql.DB) *employeeHandler.EmployeeHandler {
	employeeRepo := employeeRepository.NewMySQLRepository(db)
	employeeSrv := employeeService.NewService(employeeRepo)
	employeeHdl := employeeHandler.NewEmployeeHandler(employeeSrv)
	return employeeHdl
}
