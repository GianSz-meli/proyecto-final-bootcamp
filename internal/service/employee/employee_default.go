package employee

import (
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
	employee, ok := s.repository.GetById(id)
	if !ok {
		newError := fmt.Errorf("%w : employee with id %d not found", errors.ErrNotFound, id)
		return models.Employee{}, newError
	}

	return employee, nil
}

func (s *service) Create(employee models.Employee) (models.Employee, error) {
	if s.repository.ExistsByCardNumberId(employee.CardNumberID) {
		newError := errors.WrapErrAlreadyExist("employee", "card_number_id", employee.CardNumberID)
		return models.Employee{}, newError
	}

	if err := s.repository.Create(&employee); err != nil {
		newError := fmt.Errorf("%w : failed to create employee with card_number_id %s", errors.ErrGeneral, employee.CardNumberID)
		return models.Employee{}, newError
	}

	return employee, nil
}

func (s *service) Update(id int, employee models.Employee) (models.Employee, error) {
	current, ok := s.repository.GetById(id)
	if !ok {
		newError := fmt.Errorf("%w : employee with id %d not found", errors.ErrNotFound, id)
		return models.Employee{}, newError
	}

	if current.CardNumberID != employee.CardNumberID {
		if s.repository.ExistsByCardNumberId(employee.CardNumberID) {
			newError := errors.WrapErrAlreadyExist("employee", "card_number_id", employee.CardNumberID)
			return models.Employee{}, newError
		}
	}

	if err := s.repository.Update(id, employee); err != nil {
		newError := fmt.Errorf("%w : failed to update employee with id %d", errors.ErrGeneral, id)
		return models.Employee{}, newError
	}

	return employee, nil
}

func (s *service) Delete(id int) error {
	_, ok := s.repository.GetById(id)
	if !ok {
		return fmt.Errorf("%w : employee with id %d not found", errors.ErrNotFound, id)
	}

	s.repository.Delete(id)
	return nil
}
