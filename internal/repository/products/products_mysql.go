package products

import (
	pkgErrors "ProyectoFinal/pkg/errors"
	"ProyectoFinal/pkg/models"
	"database/sql"
	"errors"
)

func NewProductSQL(db *sql.DB) *ProductSQL {
	return &ProductSQL{db: db}
}

type ProductSQL struct {
	db *sql.DB
}

func (r *ProductSQL) CreateProduct(newProd models.Product) (models.Product, error) {
	res, err := r.db.Exec(
		QueryCreateProduct,
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
		newProd.SellerID,
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
	rows, err := r.db.Query(QueryFindAllProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.ProductCode, &p.Description, &p.Width, &p.Height, &p.Length, &p.NetWeight, &p.ExpirationRate, &p.Temperature, &p.FreezingRate, &p.ProductTypeID, &p.SellerID); err != nil {
			return nil, err
		}
		products[p.ID] = p
	}
	return products, nil
}

func (r *ProductSQL) FindProductsById(id int) (models.Product, error) {
	var p models.Product

	err := r.db.QueryRow(QueryFindProductById, id).Scan(&p.ID, &p.ProductCode, &p.Description, &p.Width, &p.Height, &p.Length, &p.NetWeight, &p.ExpirationRate, &p.Temperature, &p.FreezingRate, &p.ProductTypeID, &p.SellerID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Product{}, pkgErrors.WrapErrNotFound("product", "id", id)
		}
		return models.Product{}, err
	}
	return p, nil
}

func (r *ProductSQL) UpdateProduct(id int, prod models.Product) (models.Product, error) {
	_, err := r.db.Exec(
		QueryUpdateProduct,
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
		prod.SellerID,
		id,
	)
	if err != nil {
		return models.Product{}, err
	}
	return prod, nil
}

func (r *ProductSQL) DeleteProduct(id int) error {
	_, err := r.db.Exec(QueryDeleteProduct, id)
	if err != nil {
		return err
	}
	return nil
}
