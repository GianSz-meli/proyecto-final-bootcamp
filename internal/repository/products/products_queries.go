package products

const (
	QueryCreateProduct = `
		INSERT INTO products (product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)	
	`
	QueryFindAllProducts = `
		SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id
		FROM products
	`
	QueryFindProductById = `
		SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id, seller_id
		FROM products
		WHERE id = ?
	`
	QueryUpdateProduct = `
		UPDATE products
		SET product_code = ?, description = ?, width = ?, height = ?, length = ?, net_weight = ?, expiration_rate = ?, recommended_freezing_temperature = ?, freezing_rate = ?, product_type_id = ?, seller_id = ?	
		WHERE id = ?
	`
	QueryDeleteProduct = `
		DELETE FROM products
		WHERE id = ?
	`
	QueryExistsProdCode = `
		SELECT EXISTS(SELECT 1 FROM products WHERE product_code = ?)
	`
)
