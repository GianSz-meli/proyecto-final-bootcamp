package loader

import (
	"ProyectoFinal/pkg/models"
	"encoding/json"
	"fmt"
	"os"
)

type EmployeeLoader struct {
	path string
}

func (e *EmployeeLoader) Load() (map[int]models.Employee, error) {
	file, err := os.Open(e.path)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	var employeesJSON []models.EmployeeDoc

	if err = json.NewDecoder(file).Decode(&employeesJSON); err != nil {
		return nil, err
	}

	employeeMap := map[int]models.Employee{}

	for _, employeeDoc := range employeesJSON {
		employee := employeeDoc.DocToModel()
		employeeMap[employee.ID] = employee
	}

	return employeeMap, nil
}

