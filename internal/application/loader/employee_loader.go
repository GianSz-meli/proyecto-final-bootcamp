package loader

import (
	employeemodel "ProyectoFinal/pkg/models/employee"
	"encoding/json"
	"fmt"
	"os"
)

type EmployeeLoader struct {
	path string
}

func (e *EmployeeLoader) Load() (map[int]employeemodel.Employee, error) {
	file, err := os.Open(e.path)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	var employeesJSON []employeemodel.Employee

	if err = json.NewDecoder(file).Decode(&employeesJSON); err != nil {
		return nil, err
	}

	employeeMap := map[int]employeemodel.Employee{}

	for _, employee := range employeesJSON {
		employeeMap[employee.ID] = employee
	}

	return employeeMap, nil
}
