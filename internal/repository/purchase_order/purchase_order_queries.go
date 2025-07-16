package purchase_order

const (
	// CreatePurchaseOrder inserts a new purchase order record into the purchase_orders table.
	// Parameters: order_number, order_date, tracking_code, buyer_id, carrier_id, order_status_id, warehouse_id (ID is auto-generated)
	CreatePurchaseOrder = "INSERT INTO purchase_orders \n(order_number, order_date, tracking_code, buyer_id, carrier_id, order_status_id, warehouse_id)\nVALUES (?, ?, ?, ?, ?, ?, ?)"

	// CreateOrderDetail inserts a new order detail record into the order_details table.
	// Parameters: cleanliness_status, quantity, temperature, product_record_id, purchase_order_id (ID is auto-generated)
	CreateOrderDetail = "INSERT INTO order_details \n(cleanliness_status, quantity, temperature, product_record_id, purchase_order_id)\nVALUES (?, ?, ?, ?, ?)"
)
