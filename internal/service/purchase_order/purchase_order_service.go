package purchase_order

import "ProyectoFinal/pkg/models"

type Service interface {
	Create(purchaseOrder *models.PurchaseOrder) (*models.PurchaseOrder, error)
	GetByBuyerId(buyerId int) ([]*models.PurchaseOrderWithAllFields, error)
}
