package purchase_order

import (
	"ProyectoFinal/pkg/models"
	"database/sql"
	"fmt"
)

type purchaseOrderMySql struct {
	db *sql.DB
}

func NewPurchaseOrderMySqlRepository(newDB *sql.DB) Repository {
	return &purchaseOrderMySql{
		db: newDB,
	}
}

func (r *purchaseOrderMySql) Create(purchaseOrder *models.PurchaseOrder) (*models.PurchaseOrder, error) {
	//TODO implement me
	panic("implement me")
}

func (r *purchaseOrderMySql) GetByBuyerId(buyerId int) ([]*models.PurchaseOrderWithAllFields, error) {
	rows, err := r.db.Query(GetByBuyerId, buyerId)
	if err != nil {
		return nil, fmt.Errorf("error querying purchase orders by buyer ID %d: %w", buyerId, err)
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
			return nil, fmt.Errorf("error scanning purchase order row: %w", scanErr)
		}

		if localityID.Valid {
			localityValue := int(localityID.Int32)
			po.Warehouse.LocalityId = &localityValue
		} else {
			po.Warehouse.LocalityId = nil
		}

		purchaseOrders = append(purchaseOrders, po)
	}

	if rowsErr := rows.Err(); rowsErr != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", rowsErr)
	}

	return purchaseOrders, nil
}
