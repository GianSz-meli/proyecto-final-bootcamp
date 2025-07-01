package buyer

import (
	"ProyectoFinal/internal/repository/buyer"
	"ProyectoFinal/pkg/models"
)

type buyerService struct {
	repository buyer.Repository
}

func NewBuyerService(newRepository buyer.Repository) Service {
	return &buyerService{
		repository: newRepository,
	}
}

func (b *buyerService) Save(buyerDto models.BuyerCreateDto) error {
	//TODO implement me
	panic("implement me")
}

func (b *buyerService) GetById(id int) (models.Buyer, error) {
	return b.repository.GetById(id)
}

func (b *buyerService) GetAll() []models.Buyer {
	return b.repository.GetAll()
}

func (b *buyerService) Update(buyerUpdateDto models.BuyerUpdateDto) error {
	//TODO implement me
	panic("implement me")
}

func (b *buyerService) Delete(id int) error {
	return b.repository.Delete(id)
}
