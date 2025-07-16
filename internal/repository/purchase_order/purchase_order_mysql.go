package purchase_order

import (
	pkgErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"database/sql"
)

type purchaseOrderMySql struct {
	db *sql.DB
}

func NewPurchaseOrderMySqlRepository(newDB *sql.DB) Repository {
	return &purchaseOrderMySql{
		db: newDB,
	}
}

func (r *purchaseOrderMySql) GetByBuyerId(buyerId int) ([]*models.PurchaseOrderWithAllFields, error) {
	rows, rowsErr := r.db.Query(GetPurchaseOrdersByBuyerId, buyerId)
	if rowsErr != nil {
		return nil, rowsErr
	}
	defer rows.Close()

	var purchaseOrders []*models.PurchaseOrderWithAllFields

	for rows.Next() {
		po := &models.PurchaseOrderWithAllFields{}
		var localityID sql.NullInt32

		scanErr := rows.Scan(
			&po.Id, &po.OrderNumber, &po.OrderDate, &po.TrackingCode,
			&po.Buyer.Id, &po.Buyer.CardNumberId, &po.Buyer.FirstName, &po.Buyer.LastName,
			&po.Carrier.Id, &po.Carrier.CID, &po.Carrier.CompanyName, &po.Carrier.Address, &po.Carrier.Telephone, &po.Carrier.LocalityID,
			&po.OrderStatus.Id, &po.OrderStatus.Description,
			&po.Warehouse.ID, &po.Warehouse.WarehouseCode, &po.Warehouse.Address, &po.Warehouse.Telephone, &po.Warehouse.MinimumCapacity, &po.Warehouse.MinimumTemperature, &localityID,
		)

		if scanErr != nil {
			return nil, scanErr
		}

		if localityID.Valid {
			localityValue := int(localityID.Int32)
			po.Warehouse.LocalityId = &localityValue
		} else {
			po.Warehouse.LocalityId = nil
		}

		purchaseOrders = append(purchaseOrders, po)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(purchaseOrders) == 0 {
		return nil, pkgErrors.WrapErrNotFound("buyer", "id", buyerId)
	}

	return purchaseOrders, nil
}

func (r *purchaseOrderMySql) Create(purchaseOrder *models.PurchaseOrder) (*models.PurchaseOrder, error) {
	tx, txErr := r.db.Begin()
	if txErr != nil {
		return nil, txErr
	}
	defer tx.Rollback()

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
		return nil, execErr
	}

	lastInsertId, idErr := result.LastInsertId()
	if idErr != nil {
		return nil, idErr
	}
	purchaseOrderId := int(lastInsertId)
	purchaseOrder.Id = purchaseOrderId

	stmt, stmtErr := tx.Prepare(CreateOrderDetail)
	if stmtErr != nil {
		return nil, stmtErr
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
			return nil, detailErr
		}

		detailId, detailIdErr := detailResult.LastInsertId()
		if detailIdErr != nil {
			return nil, detailIdErr
		}

		purchaseOrder.OrderDetails[i].Id = int(detailId)
		purchaseOrder.OrderDetails[i].PurchaseOrderId = purchaseOrderId
	}

	if commitErr := tx.Commit(); commitErr != nil {
		return nil, commitErr
	}

	return purchaseOrder, nil
}
