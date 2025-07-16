package purchase_order

import (
	"ProyectoFinal/internal/repository/purchase_order"
	"ProyectoFinal/pkg/models"
)

type purchaseOrderService struct {
	repository purchase_order.Repository
}

// NewPurchaseOrderService creates and returns a new instance of the purchase order service.
func NewPurchaseOrderService(newRepository purchase_order.Repository) Service {
	return &purchaseOrderService{
		repository: newRepository,
	}
}

// Create processes a new purchase order creation by delegating to the repository layer.
func (s *purchaseOrderService) Create(purchaseOrder *models.PurchaseOrder) (*models.PurchaseOrder, error) {
	return s.repository.Create(purchaseOrder)
}
