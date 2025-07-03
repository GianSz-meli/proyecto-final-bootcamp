package seller

import "ProyectoFinal/pkg/models"

type SellerService interface {
	Create(seller models.Seller) (models.Seller, error)
	GetById(id int) (models.Seller, error)
	Update(id int, seller models.Seller) (models.Seller, error)
}
