package section

import (
	"ProyectoFinal/pkg/models"

	"github.com/stretchr/testify/mock"
)

// MockSectionRepository - implementation of section repository interface
type MockSectionRepository struct {
	mock.Mock
}

func (m *MockSectionRepository) GetAll() ([]models.Section, error) {
	args := m.Called()
	return args.Get(0).([]models.Section), args.Error(1)
}

func (m *MockSectionRepository) GetById(id int) (models.Section, error) {
	args := m.Called(id)
	return args.Get(0).(models.Section), args.Error(1)
}

func (m *MockSectionRepository) Create(section models.Section) (models.Section, error) {
	args := m.Called(section)
	return args.Get(0).(models.Section), args.Error(1)
}

func (m *MockSectionRepository) Update(id int, section models.Section) (models.Section, error) {
	args := m.Called(id, section)
	return args.Get(0).(models.Section), args.Error(1)
}

func (m *MockSectionRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockSectionRepository) ExistBySectionNumber(sectionNumber string) bool {
	args := m.Called(sectionNumber)
	return args.Bool(0)
}
