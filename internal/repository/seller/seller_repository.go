package seller

import "ProyectoFinal/pkg/models"

type SellerRepository interface {
	Create(seller *models.Seller) error
	GetById(id int) (models.Seller, bool)
	ExistsByCid(cid int) bool
	GetAll() []models.Seller
	Delete(id int)
}
