package models

type ProductRecord struct {
	ID             int
	LastUpdateDate string
	PurchasePrice  float64
	SalePrice      float64
	ProductID      int
}

type ProductRecordDoc struct {
	ID             int     `json:"id"`
	LastUpdateDate string  `json:"last_update_date" validate:"required"`
	PurchasePrice  float64 `json:"purchase_price" validate:"required,gte=0"`
	SalePrice      float64 `json:"sale_price" validate:"required,gte=0"`
	ProductID      int     `json:"product_id" validate:"required,gt=0"`
}

func (p *ProductRecord) ModelToDoc() ProductRecordDoc {
	return ProductRecordDoc{
		ID:             p.ID,
		LastUpdateDate: p.LastUpdateDate,
		PurchasePrice:  p.PurchasePrice,
		SalePrice:      p.SalePrice,
		ProductID:      p.ProductID,
	}
}

func (p *ProductRecordDoc) DocToModel() ProductRecord {
	return ProductRecord{
		ID:             p.ID,
		LastUpdateDate: p.LastUpdateDate,
		PurchasePrice:  p.PurchasePrice,
		SalePrice:      p.SalePrice,
		ProductID:      p.ProductID,
	}
}

type ReportProductData struct {
	ProductID    int
	Description  string
	RecordsCount int
}
