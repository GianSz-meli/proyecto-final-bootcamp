package locality

const (
	SQL_CREATE    = "INSERT INTO localities (locality_name, province_id) VALUES (?,(SELECT p.id FROM provinces p WHERE p.province_name = ? 	AND p.country_id = (SELECT c.id FROM countries c WHERE c.country_name = ?) ))"
	SQL_GET_BY_ID = "SELECT l.id, l.locality_name ,p.id, p.province_name, c.id,c.country_name from localities l JOIN provinces p ON p.id = l.province_id JOIN countries c ON c.id  = p.country_id WHERE l.id = ?"
)
