package buyer

import (
	"ProyectoFinal/pkg/models"
)

type Repository interface {
	GetById(id int) (*models.Buyer, error)
	GetAll() ([]*models.Buyer, error)
	Create(buyer *models.Buyer) (*models.Buyer, error)
	Update(buyer *models.Buyer) (*models.Buyer, error)
	Delete(id int) error
	ExistsByCardNumberId(id string) (bool, error)
}

type RepositoryMap interface {
	Create(buyer models.Buyer) models.Buyer
	GetById(id int) (models.Buyer, bool)
	GetAll() []models.Buyer
	Update(buyer models.Buyer) models.Buyer
	Delete(id int)
	ExistsByCardNumberId(id string) bool
}
