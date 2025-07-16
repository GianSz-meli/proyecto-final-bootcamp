package buyer

const (
	GetBuyer                       = "SELECT  id, id_card_number, first_name, last_name FROM buyers WHERE id = ?"
	GetAllBuyers                   = "SELECT id, id_card_number, first_name, last_name FROM buyers"
	CreateBuyer                    = "INSERT INTO buyers (id_card_number, first_name, last_name) VALUES (?,?,?)"
	UpdateBuyer                    = "UPDATE buyers SET id_card_number = ?, first_name = ?, last_name = ? WHERE id = ?"
	DeleteBuyer                    = "DELETE FROM buyers WHERE id = ?"
	GetBuyerWithPurchaseOrders     = "SELECT \nb.id,\nb.id_card_number,\nb.first_name,\nb.last_name,\nCOUNT(po.id) as total_purchase_orders\nFROM buyers b\nLEFT JOIN purchase_orders po ON b.id = po.buyer_id\nWHERE b.id = ?\nGROUP BY b.id, b.id_card_number, b.first_name, b.last_name"
	GetAllBuyersWithPurchaseOrders = "SELECT \nb.id,\nb.id_card_number,\nb.first_name,\nb.last_name,\nCOUNT(po.id) as total_purchase_orders\nFROM buyers b\nLEFT JOIN purchase_orders po ON b.id = po.buyer_id\nGROUP BY b.id, b.id_card_number, b.first_name, b.last_name\nORDER BY b.id;"
)
