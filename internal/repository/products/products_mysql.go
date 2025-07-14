package repository

import (
	"ProyectoFinal/pkg/models"
	"database/sql"
)

func NewProductSQL(db *sql.DB) *ProductSQL {
	return &ProductSQL{db: db}
}

type ProductSQL struct {
	db *sql.DB
}

func (r *ProductSQL) ExistsProdCode(prodCode string) bool {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM products WHERE product_code = ?)"
	err := r.db.QueryRow(query, prodCode).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func (r *ProductSQL) CreateProduct(newProd models.Product) (models.Product, error) {
	println("hola")
	query := `
        INSERT INTO products 
            (product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `

	res, err := r.db.Exec(
		query,
		newProd.ProductCode,
		newProd.Description,
		newProd.Width,
		newProd.Height,
		newProd.Length,
		newProd.NetWeight,
		newProd.ExpirationRate,
		newProd.Temperature,
		newProd.FreezingRate,
		newProd.ProductTypeID,
	)

	if err != nil {
		return newProd, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return newProd, err
	}
	newProd.ID = int(id)
	return newProd, nil
}

func (r *ProductSQL) FindAllProducts() (map[int]models.Product, error) {
	products := make(map[int]models.Product)
	rows, err := r.db.Query("SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.ProductCode, &p.Description, &p.Width, &p.Height, &p.Length, &p.NetWeight, &p.ExpirationRate, &p.Temperature, &p.FreezingRate, &p.ProductTypeID); err != nil {
			return nil, err
		}
		products[p.ID] = p
	}
	return products, nil
}

func (r *ProductSQL) FindProductsById(id int) (models.Product, bool) {
	var p models.Product
	query := "SELECT id, product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id FROM products WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.ProductCode, &p.Description, &p.Width, &p.Height, &p.Length, &p.NetWeight, &p.ExpirationRate, &p.Temperature, &p.FreezingRate, &p.ProductTypeID)
	if err != nil {
		if err == sql.ErrNoRows {
			return p, false
		}
		return p, false
	}
	return p, true
}

func (r *ProductSQL) UpdateProduct(id int, prod models.Product) (models.Product, error) {
	query := `
        UPDATE products SET
            product_code = ?,
            description = ?,
            width = ?,
            height = ?,
            length = ?,
            net_weight = ?,
            expiration_rate = ?,
            recommended_freezing_temperature = ?,
            freezing_rate = ?,
            product_type_id = ?
        WHERE id = ?
    `
	_, err := r.db.Exec(
		query,
		prod.ProductCode,
		prod.Description,
		prod.Width,
		prod.Height,
		prod.Length,
		prod.NetWeight,
		prod.ExpirationRate,
		prod.Temperature,
		prod.FreezingRate,
		prod.ProductTypeID,
		id,
	)
	if err != nil {
		return models.Product{}, err
	}
	return prod, nil
}

func (r *ProductSQL) DeleteProduct(id int) {
	r.db.Exec("DELETE FROM products WHERE id = ?", id)
}
