package seller

import "ProyectoFinal/pkg/models"

type SellerRepository interface {
	Create(seller models.Seller) (models.Seller, error)
	GetById(id int) (*models.Seller, error)
	ExistsByCid(cid string) (bool, error)
	Update(seller *models.Seller) (models.Seller, error)
	GetAll() ([]models.Seller, error)
	Delete(id int) error
}
