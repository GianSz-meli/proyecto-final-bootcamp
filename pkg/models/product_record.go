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
	LastUpdateDate string  `json:"last_update_date"`
	PurchasePrice  float64 `json:"purchase_price"`
	SalePrice      float64 `json:"sale_price"`
	ProductID      int     `json:"product_id"`
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

// type ReportProductDataDoc struct {
// 	ProductID    int    `json:"product_id"`
// 	Description  string `json:"description"`
// 	RecordsCount int    `json:"records_count"`
// }

// func (p *ReportProductData) ModelToDoc() ReportProductDataDoc {
// 	return ReportProductDataDoc{
// 		ProductID:    p.ProductID,
// 		Description:  p.Description,
// 		RecordsCount: p.RecordsCount,
// 	}
// }

// func (p *ReportProductDataDoc) DocToModel() ReportProductData {
// 	return ReportProductData{
// 		ProductID:    p.ProductID,
// 		Description:  p.Description,
// 		RecordsCount: p.RecordsCount,
// 	}
// }

