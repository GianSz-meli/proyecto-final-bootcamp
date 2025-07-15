package models

import "time"

type InboundOrder struct {
	ID             int       `json:"id"`
	OrderDate      time.Time `json:"order_date"`
	OrderNumber    string    `json:"order_number"`
	EmployeeID     int       `json:"employee_id"`
	ProductBatchID int       `json:"product_batch_id"`
	WarehouseID    int       `json:"warehouse_id"`
}

type InboundOrderRequest struct {
	OrderDate      string `json:"order_date" validate:"required"`
	OrderNumber    string `json:"order_number" validate:"required"`
	EmployeeID     int    `json:"employee_id" validate:"required,gt=0"`
	ProductBatchID int    `json:"product_batch_id" validate:"required,gt=0"`
	WarehouseID    int    `json:"warehouse_id" validate:"required,gt=0"`
}

func (r InboundOrderRequest) DocToModel() InboundOrder {
	orderDate, err := time.Parse("2006-01-02", r.OrderDate)
	if err != nil {
		orderDate = time.Time{}
	}

	return InboundOrder{
		OrderDate:      orderDate,
		OrderNumber:    r.OrderNumber,
		EmployeeID:     r.EmployeeID,
		ProductBatchID: r.ProductBatchID,
		WarehouseID:    r.WarehouseID,
	}
}

type InboundOrderDoc struct {
	ID             int       `json:"id"`
	OrderDate      time.Time `json:"order_date"`
	OrderNumber    string    `json:"order_number"`
	EmployeeID     int       `json:"employee_id"`
	ProductBatchID int       `json:"product_batch_id"`
	WarehouseID    int       `json:"warehouse_id"`
}

func (i InboundOrder) ModelToDoc() InboundOrderDoc {
	return InboundOrderDoc{
		ID:             i.ID,
		OrderDate:      i.OrderDate,
		OrderNumber:    i.OrderNumber,
		EmployeeID:     i.EmployeeID,
		ProductBatchID: i.ProductBatchID,
		WarehouseID:    i.WarehouseID,
	}
}
