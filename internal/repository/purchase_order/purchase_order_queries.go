package purchase_order

const (
	GetPurchaseOrdersByBuyerId = "SELECT \npo.id, po.order_number, po.order_date, po.tracking_code,\nb.id, b.id_card_number, b.first_name, b.last_name,\nc.id, c.cid, c.company_name, c.address, c.telephone, c.locality_id,\nos.id, os.description,\nw.id, w.warehouse_code, w.address, w.telephone, w.minimum_capacity, \nw.minimum_temperature, w.locality_id\nFROM purchase_orders po\nINNER JOIN buyers b ON po.buyer_id = b.id\nINNER JOIN carriers c ON po.carrier_id = c.id\nINNER JOIN order_status os ON po.order_status_id = os.id\nINNER JOIN warehouses w ON po.warehouse_id = w.id\nWHERE po.buyer_id = ?\nORDER BY po.order_date DESC"
	CreatePurchaseOrder        = "INSERT INTO purchase_orders \n(order_number, order_date, tracking_code, buyer_id, carrier_id, order_status_id, warehouse_id)\nVALUES (?, ?, ?, ?, ?, ?, ?)"
	CreateOrderDetail          = "INSERT INTO order_details \n(cleanliness_status, quantity, temperature, product_record_id, purchase_order_id)\nVALUES (?, ?, ?, ?, ?)"
)
