package purchase_order

import "ProyectoFinal/pkg/models"

// Service defines the business logic contract for purchase order operations.
type Service interface {
	Create(purchaseOrder *models.PurchaseOrder) (*models.PurchaseOrder, error)
}
