package buyer

const (
	GET_BUYER             = "SELECT  id, id_card_number, first_name, last_name FROM buyers WHERE id = ?"
	GET_ALL_BUYERS        = "SELECT id, id_card_number, first_name, last_name FROM buyers"
	CREATE_BUYER          = "INSERT INTO buyers (id_card_number, first_name, last_name) VALUES (?,?,?)"
	UPDATE_BUYER          = "UPDATE buyers SET id_card_number = ?, first_name = ?, last_name = ? WHERE id = ?"
	DELETE_BUYER          = "DELETE FROM buyers WHERE id = ?"
	EXISTS_BY_CARD_NUMBER = "SELECT 1 FROM buyers WHERE id_card_number = ? LIMIT 1"
)
