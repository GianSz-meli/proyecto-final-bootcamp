package seller

const (
	SQL_CREATE       = "INSERT INTO sellers (cid, company_name, address, telephone, locality_id) VALUES (?,?,?,?,?)"
	SQL_GET_BY_ID    = "SELECT  id, cid, company_name, address, telephone, locality_id FROM sellers WHERE id = ?"
	SQL_EXIST_BY_CID = "SELECT EXISTS(SELECT 1 FROM sellers WHERE cid = ?)"
	SQL_GET_ALL      = "SELECT id, cid, company_name, address, telephone, locality_id FROM sellers"
	SQL_DELETE       = "DELETE FROM sellers WHERE id = ?"
	SQL_UPDATE       = "UPDATE sellers SET cid = ?, company_name = ?, address = ?, telephone = ?, locality_id = ? WHERE id = ?"
)
