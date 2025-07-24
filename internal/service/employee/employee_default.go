package employee

import (
	"ProyectoFinal/internal/handler/utils"
	employeeRepository "ProyectoFinal/internal/repository/employee"
	"ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"fmt"
)

type service struct {
	repository employeeRepository.Repository
}

func NewService(repository employeeRepository.Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetAll() ([]models.Employee, error) {
	employees, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (s *service) GetById(id int) (models.Employee, error) {
	employee, err := s.repository.GetById(id)
	if err != nil {
		return models.Employee{}, err
	}

	return employee, nil
}

func (s *service) Create(employee models.Employee) (models.Employee, error) {
	if err := s.repository.Create(&employee); err != nil {
		return models.Employee{}, err
	}

	return employee, nil
}

func (s *service) PatchUpdate(id int, updateRequest *models.EmployeeUpdateRequest) (models.Employee, error) {
	// Get existing employee
	employeeToUpdate, err := s.repository.GetById(id)
	if err != nil {
		return models.Employee{}, err
	}

	// Apply partial updates
	if updated := utils.UpdateFields(&employeeToUpdate, updateRequest); !updated {
		return models.Employee{}, fmt.Errorf("%w : no fields provided for update", errors.ErrUnprocessableEntity)
	}

	// Persist changes
	if err := s.repository.Update(id, employeeToUpdate); err != nil {
		return models.Employee{}, err
	}

	return employeeToUpdate, nil
}

func (s *service) Delete(id int) error {
	if err := s.repository.Delete(id); err != nil {
		return err
	}
	return nil
}
