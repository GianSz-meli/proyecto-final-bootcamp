package models

type PurchaseOrder struct {
	Id            int    `json:"id"`
	OrderNumber   string `json:"order_number"`
	OrderDate     string `json:"order_date"`
	TrackingCode  string `json:"tracking_code"`
	BuyerId       int    `json:"buyer"`
	CarrierId     int    `json:"carrier"`
	OrderStatusId int    `json:"order_status"`
	WarehouseId   int    `json:"warehouse"`
}

type PurchaseOrderWithAllFields struct {
	Id           int               `json:"id"`
	OrderNumber  string            `json:"order_number"`
	OrderDate    string            `json:"order_date"`
	TrackingCode string            `json:"tracking_code"`
	Buyer        BuyerDoc          `json:"buyer"`
	Carrier      CarrierTemp       `json:"carrier"`
	OrderStatus  OrderStatus       `json:"order_status"`
	Warehouse    WarehouseDocument `json:"warehouse"`
}

type OrderStatus struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}

type CarrierTemp struct {
	Id          int    `json:"id"`
	CID         string `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
	LocalityID  int    `json:"locality_id"`
}
