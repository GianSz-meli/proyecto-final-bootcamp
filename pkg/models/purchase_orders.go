package models

type PurchaseOrder struct {
	Id            int            `json:"id"`
	OrderNumber   string         `json:"order_number"`
	OrderDate     string         `json:"order_date"`
	TrackingCode  string         `json:"tracking_code"`
	BuyerId       int            `json:"buyer"`
	CarrierId     int            `json:"carrier"`
	OrderStatusId int            `json:"order_status"`
	WarehouseId   int            `json:"warehouse"`
	OrderDetails  []OrderDetails `json:"order_details"`
}
type PurchaseOrderCreateDTO struct {
	OrderNumber   string                  `json:"order_number" validate:"required"`
	OrderDate     string                  `json:"order_date" validate:"required"`
	TrackingCode  string                  `json:"tracking_code" validate:"required"`
	BuyerId       int                     `json:"buyer" validate:"required"`
	CarrierId     int                     `json:"carrier" validate:"required"`
	OrderStatusId int                     `json:"order_status" validate:"required"`
	WarehouseId   int                     `json:"warehouse" validate:"required"`
	OrderDetails  []OrderDetailsCreateDTO `json:"order_details" validate:"required,dive"`
}

func (po PurchaseOrderCreateDTO) CreateDtoToModel() *PurchaseOrder {
	var orderDetails []OrderDetails
	for _, detailDto := range po.OrderDetails {
		orderDetails = append(orderDetails, OrderDetails{
			CleanlinessStatus: detailDto.CleanlinessStatus,
			Quantity:          detailDto.Quantity,
			Temperature:       detailDto.Temperature,
			ProductRecordId:   detailDto.ProductRecordId,
		})
	}

	return &PurchaseOrder{
		OrderNumber:   po.OrderNumber,
		OrderDate:     po.OrderDate,
		TrackingCode:  po.TrackingCode,
		BuyerId:       po.BuyerId,
		CarrierId:     po.CarrierId,
		OrderStatusId: po.OrderStatusId,
		WarehouseId:   po.WarehouseId,
		OrderDetails:  orderDetails,
	}
}

type PurchaseOrderDoc struct {
	Id            int            `json:"id"`
	OrderNumber   string         `json:"order_number"`
	OrderDate     string         `json:"order_date"`
	TrackingCode  string         `json:"tracking_code"`
	BuyerId       int            `json:"buyer"`
	CarrierId     int            `json:"carrier"`
	OrderStatusId int            `json:"order_status"`
	WarehouseId   int            `json:"warehouse"`
	OrderDetails  []OrderDetails `json:"order_details"`
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
		OrderDetails:  po.OrderDetails,
	}
}
