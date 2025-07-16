package buyer

const (
	// GetBuyer retrieves a single buyer by their ID.
	// Returns: id, id_card_number, first_name, last_name
	GetBuyer = "SELECT  id, id_card_number, first_name, last_name FROM buyers WHERE id = ?"

	// GetAllBuyers retrieves all buyers from the database.
	// Returns: id, id_card_number, first_name, last_name for each buyer
	GetAllBuyers = "SELECT id, id_card_number, first_name, last_name FROM buyers"

	// CreateBuyer inserts a new buyer into the database.
	// Parameters: id_card_number, first_name, last_name (ID is auto-generated)
	CreateBuyer = "INSERT INTO buyers (id_card_number, first_name, last_name) VALUES (?,?,?)"

	// UpdateBuyer modifies an existing buyer's information.
	// Parameters: id_card_number, first_name, last_name, id (WHERE condition)
	UpdateBuyer = "UPDATE buyers SET id_card_number = ?, first_name = ?, last_name = ? WHERE id = ?"

	// DeleteBuyer removes a buyer from the database by their ID.
	// Parameters: id (buyer to delete)
	DeleteBuyer = "DELETE FROM buyers WHERE id = ?"

	// GetBuyerWithPurchaseOrders retrieves a specific buyer with their total purchase orders count.
	// Uses LEFT JOIN to include buyer even if they have zero orders.
	// Parameters: buyer_id
	// Returns: id, id_card_number, first_name, last_name, total_purchase_orders
	GetBuyerWithPurchaseOrders = "SELECT \nb.id,\nb.id_card_number,\nb.first_name,\nb.last_name,\nCOUNT(po.id) as total_purchase_orders\nFROM buyers b\nLEFT JOIN purchase_orders po ON b.id = po.buyer_id\nWHERE b.id = ?\nGROUP BY b.id, b.id_card_number, b.first_name, b.last_name"

	// GetAllBuyersWithPurchaseOrders retrieves all buyers with their respective purchase orders count.
	// Uses LEFT JOIN to include all buyers, showing 0 count for buyers without orders.
	// Results are ordered by buyer ID for consistent output.
	// Returns: id, id_card_number, first_name, last_name, total_purchase_orders for each buyer
	GetAllBuyersWithPurchaseOrders = "SELECT \nb.id,\nb.id_card_number,\nb.first_name,\nb.last_name,\nCOUNT(po.id) as total_purchase_orders\nFROM buyers b\nLEFT JOIN purchase_orders po ON b.id = po.buyer_id\nGROUP BY b.id, b.id_card_number, b.first_name, b.last_name\nORDER BY b.id;"
)
