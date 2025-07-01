package employee

import (
	employeeRepository "ProyectoFinal/internal/repository/employee"
	"ProyectoFinal/pkg/errors"
	employeemodel "ProyectoFinal/pkg/models/employee"
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

func (s *service) GetAll() ([]employeemodel.Employee, error) {
	employees, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func (s *service) GetById(id int) (employeemodel.Employee, error) {
	employee, ok := s.repository.GetById(id)
	if !ok {
		return employeemodel.Employee{}, errors.ErrNotFound
	}

	return employee, nil
}

func (s *service) Create(employee employeemodel.Employee) (employeemodel.Employee, error) {
	if s.repository.ExistsByCardNumberId(employee.CardNumberID) {
		newError := fmt.Errorf("%w : employee with card_number_id %s already exists", errors.ErrAlreadyExists, employee.CardNumberID)
		return employeemodel.Employee{}, newError
	}

	if err := s.repository.Create(&employee); err != nil {
		return employeemodel.Employee{}, errors.ErrGeneral
	}

	return employee, nil
}

func (s *service) Update(id int, employee employeemodel.Employee) (employeemodel.Employee, error) {
	_, ok := s.repository.GetById(id)
	if !ok {
		return employeemodel.Employee{}, errors.ErrNotFound
	}

	if err := s.repository.Update(id, employee); err != nil {
		return employeemodel.Employee{}, errors.ErrGeneral
	}

	return employee, nil
}

func (s *service) Delete(id int) error {
	exists := s.repository.Delete(id)
	if !exists {
		return fmt.Errorf("%w : employee with id %d not found", errors.ErrNotFound, id)
	}

	return nil
}
