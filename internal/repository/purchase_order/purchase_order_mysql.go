package purchase_order

import (
	"ProyectoFinal/pkg/models"
	"database/sql"
)

type purchaseOrderMySql struct {
	db *sql.DB
}

// NewPurchaseOrderMySqlRepository creates and returns a new MySQL implementation of the purchase order repository.
func NewPurchaseOrderMySqlRepository(newDB *sql.DB) Repository {
	return &purchaseOrderMySql{
		db: newDB,
	}
}

// Create creates a new purchase order with its order details in a single transaction
func (r *purchaseOrderMySql) Create(purchaseOrder *models.PurchaseOrder) (*models.PurchaseOrder, error) {
	tx, txErr := r.beginTransaction()
	if txErr != nil {
		return nil, txErr
	}
	defer tx.Rollback()

	purchaseOrderId, orderErr := r.createPurchaseOrder(tx, purchaseOrder)
	if orderErr != nil {
		return nil, orderErr
	}
	purchaseOrder.Id = purchaseOrderId

	if detailsErr := r.createOrderDetails(tx, purchaseOrder, purchaseOrderId); detailsErr != nil {
		return nil, detailsErr
	}

	if commitErr := tx.Commit(); commitErr != nil {
		return nil, commitErr
	}

	return purchaseOrder, nil
}

// beginTransaction starts a new database transaction
func (r *purchaseOrderMySql) beginTransaction() (*sql.Tx, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// createPurchaseOrder inserts the main purchase order record
func (r *purchaseOrderMySql) createPurchaseOrder(tx *sql.Tx, purchaseOrder *models.PurchaseOrder) (int, error) {
	result, execErr := tx.Exec(
		CreatePurchaseOrder,
		purchaseOrder.OrderNumber,
		purchaseOrder.OrderDate,
		purchaseOrder.TrackingCode,
		purchaseOrder.BuyerId,
		purchaseOrder.CarrierId,
		purchaseOrder.OrderStatusId,
		purchaseOrder.WarehouseId,
	)
	if execErr != nil {
		return 0, execErr
	}

	lastInsertId, idErr := result.LastInsertId()
	if idErr != nil {
		return 0, idErr
	}

	return int(lastInsertId), nil
}

// createOrderDetails inserts all order details for the purchase order
func (r *purchaseOrderMySql) createOrderDetails(tx *sql.Tx, purchaseOrder *models.PurchaseOrder, purchaseOrderId int) error {
	if len(purchaseOrder.OrderDetails) == 0 {
		return nil
	}

	stmt, stmtErr := tx.Prepare(CreateOrderDetail)
	if stmtErr != nil {
		return stmtErr
	}
	defer stmt.Close()

	for i := range purchaseOrder.OrderDetails {
		detailResult, detailErr := stmt.Exec(
			purchaseOrder.OrderDetails[i].CleanlinessStatus,
			purchaseOrder.OrderDetails[i].Quantity,
			purchaseOrder.OrderDetails[i].Temperature,
			purchaseOrder.OrderDetails[i].ProductRecordId,
			purchaseOrderId,
		)
		if detailErr != nil {
			return detailErr
		}

		detailId, detailIdErr := detailResult.LastInsertId()
		if detailIdErr != nil {
			return detailIdErr
		}

		purchaseOrder.OrderDetails[i].Id = int(detailId)
		purchaseOrder.OrderDetails[i].PurchaseOrderId = purchaseOrderId
	}

	return nil
}
