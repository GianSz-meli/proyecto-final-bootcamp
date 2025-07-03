package di

import (
	"ProyectoFinal/internal/handler"
	buyerepo "ProyectoFinal/internal/repository/buyer"
	buyersvc "ProyectoFinal/internal/service/buyer"
	"ProyectoFinal/pkg/models"
)

func GetBuyerHandler(db map[int]models.Buyer) *handler.BuyerHandler {
	repository := buyerepo.NewBuyerRepository(db)
	service := buyersvc.NewBuyerService(repository)
	return handler.NewBuyerHandler(service)
}
