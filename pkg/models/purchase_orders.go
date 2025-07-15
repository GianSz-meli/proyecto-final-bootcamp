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

type PurchaseOrderCreateDTO struct {
	OrderNumber   string `json:"order_number" validate:"required"`
	OrderDate     string `json:"order_date" validate:"required"`
	TrackingCode  string `json:"tracking_code" validate:"required"`
	BuyerId       int    `json:"buyer" validate:"required"`
	CarrierId     int    `json:"carrier" validate:"required"`
	OrderStatusId int    `json:"order_status" validate:"required"`
	WarehouseId   int    `json:"warehouse" validate:"required"`
}

func (po PurchaseOrderCreateDTO) CreateDtoToModel() *PurchaseOrder {
	return &PurchaseOrder{
		OrderNumber:   po.OrderNumber,
		OrderDate:     po.OrderDate,
		TrackingCode:  po.TrackingCode,
		BuyerId:       po.BuyerId,
		CarrierId:     po.CarrierId,
		OrderStatusId: po.OrderStatusId,
		WarehouseId:   po.WarehouseId,
	}
}

type PurchaseOrderDoc struct {
	Id            int    `json:"id"`
	OrderNumber   string `json:"order_number"`
	OrderDate     string `json:"order_date"`
	TrackingCode  string `json:"tracking_code"`
	BuyerId       int    `json:"buyer"`
	CarrierId     int    `json:"carrier"`
	OrderStatusId int    `json:"order_status"`
	WarehouseId   int    `json:"warehouse"`
}

func (po PurchaseOrder) ModelToDoc() PurchaseOrderDoc {
	return PurchaseOrderDoc{
		Id:            po.Id,
		OrderNumber:   po.OrderNumber,
		OrderDate:     po.OrderDate,
		TrackingCode:  po.TrackingCode,
		BuyerId:       po.BuyerId,
		CarrierId:     po.CarrierId,
		OrderStatusId: po.OrderStatusId,
		WarehouseId:   po.WarehouseId,
	}
}

type PurchaseOrderWithAllFieldsDoc struct {
	Id           int               `json:"id"`
	OrderNumber  string            `json:"order_number"`
	OrderDate    string            `json:"order_date"`
	TrackingCode string            `json:"tracking_code"`
	Buyer        BuyerDoc          `json:"buyer"`
	Carrier      CarrierTemp       `json:"carrier"`
	OrderStatus  OrderStatus       `json:"order_status"`
	Warehouse    WarehouseDocument `json:"warehouse"`
}

func (po PurchaseOrderWithAllFields) ModelToDoc() PurchaseOrderWithAllFieldsDoc {
	return PurchaseOrderWithAllFieldsDoc{
		Id:           po.Id,
		OrderNumber:  po.OrderNumber,
		OrderDate:    po.OrderDate,
		TrackingCode: po.TrackingCode,
		Buyer:        po.Buyer,
		Carrier:      po.Carrier,
		OrderStatus:  po.OrderStatus,
		Warehouse:    po.Warehouse,
	}
}
