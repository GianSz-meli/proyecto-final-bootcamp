package di

import (
	"ProyectoFinal/internal/handler"
	employeeRepository "ProyectoFinal/internal/repository/employee"
	employeeService "ProyectoFinal/internal/service/employee"
	"ProyectoFinal/pkg/models"
	"database/sql"
)

// GetEmployeeHandler returns a handler configured with MySQL repository
func GetEmployeeHandler(db *sql.DB) *handler.EmployeeHandler {
	employeeRepo := employeeRepository.NewMySQLRepository(db)
	employeeSrv := employeeService.NewService(employeeRepo)
	employeeHdl := handler.NewEmployeeHandler(employeeSrv)
	return employeeHdl
}

// GetEmployeeHandlerWithMap returns a handler configured with map-based repository (JSON)
// This is kept for backward compatibility if needed
func GetEmployeeHandlerWithMap(db map[int]models.Employee) *handler.EmployeeHandler {
	employeeRepo := employeeRepository.NewRepository(db)
	employeeSrv := employeeService.NewService(employeeRepo)
	employeeHdl := handler.NewEmployeeHandler(employeeSrv)
	return employeeHdl
}
