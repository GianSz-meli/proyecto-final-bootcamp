package seller

import "ProyectoFinal/pkg/models"

type SellerService interface {
	Create(seller models.Seller) (models.Seller, error)
	GetAll() []models.Seller
}
