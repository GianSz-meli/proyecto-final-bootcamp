package models

type OrderDetails struct {
	Id                int     `json:"id"`
	CleanlinessStatus string  `json:"cleanliness_status"`
	Quantity          int     `json:"quantity"`
	Temperature       float64 `json:"temperature"`
	ProductRecordId   int     `json:"product_record_id"`
	PurchaseOrderId   int     `json:"purchase_order_id"`
}

type OrderDetailsCreateDTO struct {
	CleanlinessStatus string  `json:"cleanliness_status" validate:"required"`
	Quantity          int     `json:"quantity" validate:"required,min=1"`
	Temperature       float64 `json:"temperature" validate:"required"`
	ProductRecordId   int     `json:"product_record_id" validate:"required"`
}
