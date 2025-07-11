package seller

import "ProyectoFinal/pkg/models"

type SellerService interface {
	Create(seller models.Seller) (models.Seller, error)
	GetAll() []models.Seller
	GetById(id int) (models.Seller, error)
	Delete(id int) error
	Update(id int, seller models.Seller) (models.Seller, error)
}
