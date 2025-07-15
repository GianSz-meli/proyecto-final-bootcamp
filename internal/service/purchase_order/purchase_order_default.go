package purchase_order

import (
	"ProyectoFinal/internal/repository/purchase_order"
	"ProyectoFinal/pkg/models"
)

type purchaseOrderService struct {
	repository purchase_order.Repository
}

func NewPurchaseOrderService(newRepository purchase_order.Repository) Service {
	return &purchaseOrderService{
		repository: newRepository,
	}
}

func (s *purchaseOrderService) Create(purchaseOrder *models.PurchaseOrder) (*models.PurchaseOrder, error) {
	return s.repository.Create(purchaseOrder)
}

func (s *purchaseOrderService) GetByBuyerId(buyerId int) ([]*models.PurchaseOrderWithAllFields, error) {
	return s.repository.GetByBuyerId(buyerId)
}
