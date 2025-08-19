package middlewares

const SQL_INSERT_LOGGER = "INSERT INTO http_logs (method, request_path, remote_address, duration_ms, response, status_code) VALUES (?, ?, ?, ?, ?,?)"
