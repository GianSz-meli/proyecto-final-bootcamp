package inbound_order

const (
	QueryCreateInboundOrder = `
		INSERT INTO inbound_orders (order_date, order_number, employee_id, product_batch_id, warehouse_id) 
		VALUES (?, ?, ?, ?, ?)
	`

	QueryExistsByOrderNumber = `
		SELECT EXISTS(SELECT 1 FROM inbound_orders WHERE order_number = ?)
	`

	QueryGetEmployeeInboundOrdersReportByEmployeeId = `
		SELECT 
			e.id,
			e.card_number_id,
			e.first_name,
			e.last_name,
			e.warehouse_id,
			COUNT(io.id) as inbound_orders_count
		FROM employees e
		LEFT JOIN inbound_orders io ON e.id = io.employee_id
		WHERE e.id = ?
		GROUP BY e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id
	`

	QueryGetEmployeeInboundOrdersReportAll = `
		SELECT 
			e.id,
			e.card_number_id,
			e.first_name,
			e.last_name,
			e.warehouse_id,
			COUNT(io.id) as inbound_orders_count
		FROM employees e
		LEFT JOIN inbound_orders io ON e.id = io.employee_id
		GROUP BY e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id
		ORDER BY e.id
	`
)
