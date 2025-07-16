package locality

const (
	SQL_CREATE                 = "INSERT INTO localities (id,locality_name, province_id) VALUES (?,?,(SELECT p.id FROM provinces p WHERE p.province_name = ? 	AND p.country_id = (SELECT c.id FROM countries c WHERE c.country_name = ?) ))"
	SQL_GET_BY_ID              = "SELECT l.id, l.locality_name ,p.id, p.province_name, c.id,c.country_name from localities l JOIN provinces p ON p.id = l.province_id JOIN countries c ON c.id  = p.country_id WHERE l.id = ?"
	SQL_SELLERS_BY_ID_LOCALITY = "SELECT COUNT(*) as sellers_count, l.id, l.locality_name FROM sellers s JOIN localities l ON s.locality_id = l.id WHERE l.id = ? GROUP BY l.id"
	SQL_SELLERS_BY_LOCALITY    = "SELECT COUNT(*) as sellers_count, l.id, l.locality_name FROM sellers s JOIN localities l ON s.locality_id = l.id GROUP BY l.id"
)
