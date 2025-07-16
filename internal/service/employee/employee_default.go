package employee

import (
	employeeRepository "ProyectoFinal/internal/repository/employee"
	"ProyectoFinal/pkg/models"
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

func (s *service) Update(id int, employee models.Employee) (models.Employee, error) {
	_, err := s.repository.GetById(id)
	if err != nil {
		return models.Employee{}, err
	}

	if err := s.repository.Update(id, employee); err != nil {
		return models.Employee{}, err
	}

	return employee, nil
}

func (s *service) Delete(id int) error {
	if err := s.repository.Delete(id); err != nil {
		return err
	}
	return nil
}
