package seller

import "ProyectoFinal/pkg/models"

type SellerRepository interface {
	Create(seller models.Seller) (models.Seller, error)
	GetById(id int) (*models.Seller, error)
	Update(seller *models.Seller) (models.Seller, error)
	GetAll() ([]models.Seller, error)
	Delete(id int) error
}

type SellerRepositoryMap interface {
	Create(seller *models.Seller)
	GetById(id int) (models.Seller, bool)
	ExistsByCid(cid string) bool
	Update(seller *models.Seller)
	GetAll() []models.Seller
	Delete(id int)
}
