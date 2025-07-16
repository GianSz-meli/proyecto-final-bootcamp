package purchase_order

import "ProyectoFinal/pkg/models"

// Repository defines the contract for purchase order data access operations.
type Repository interface {
	Create(purchaseOrder *models.PurchaseOrder) (*models.PurchaseOrder, error)
}
