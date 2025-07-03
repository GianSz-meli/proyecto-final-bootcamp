package seller

import "ProyectoFinal/pkg/models"

type SellerRepository interface {
	Create(seller *models.Seller)
	GetById(id int) (models.Seller, bool)
	ExistsByCid(cid int) bool
	Update(seller *models.Seller)
}
