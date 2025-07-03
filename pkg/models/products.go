package models

type ProductType struct {
	ID          int
	Description string
}
type Product struct {
	ID             int
	ProductCode    string
	Description    string
	Width          float64
	Height         float64
	Length         float64
	NetWeight      float64
	ExpirationRate int
	Temperature    float32
	FreezingRate   float64
	ProductType    *ProductType
}
type ProductTypeDoc struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}
type ProductDoc struct {
	ID             int          `json:"id"`
	ProductCode    *string      `json:"product_code" validate:"required,min=1"`
	Description    *string      `json:"description"`
	Width          *float64     `json:"width"`
	Height         *float64     `json:"height"`
	Length         *float64     `json:"length"`
	NetWeight      *float64     `json:"net_weight"`
	ExpirationRate *int         `json:"expiration_rate"`
	Temperature    *float32     `json:"recommended_freezing_temperature"`
	FreezingRate   *float64     `json:"freezing_rate"`
	ProductType    *ProductType `json:"product_type_id"`
}

type ProductDocUpdate struct {
	ProductCode    *string      `json:"product_code" validate:"required,min=1"`
	Description    *string      `json:"description"`
	Width          *float64     `json:"width"`
	Height         *float64     `json:"height"`
	Length         *float64     `json:"length"`
	NetWeight      *float64     `json:"net_weight"`
	ExpirationRate *int         `json:"expiration_rate"`
	Temperature    *float32     `json:"recommended_freezing_temperature"`
	FreezingRate   *float64     `json:"freezing_rate"`
	ProductType    *ProductType `json:"product_type_id"`
}

func (p *Product) ModelToDoc() ProductDoc {
	return ProductDoc{
		ID:             p.ID,
		ProductCode:    &p.ProductCode,
		Description:    &p.Description,
		Width:          &p.Width,
		Height:         &p.Height,
		Length:         &p.Length,
		NetWeight:      &p.NetWeight,
		ExpirationRate: &p.ExpirationRate,
		Temperature:    &p.Temperature,
		FreezingRate:   &p.FreezingRate,
		ProductType:    p.ProductType,
	}
}

func (p *ProductDoc) DocToModel() Product {
	return Product{
		ID:             p.ID,
		ProductCode:    *p.ProductCode,
		Description:    *p.Description,
		Width:          *p.Width,
		Height:         *p.Height,
		Length:         *p.Length,
		NetWeight:      *p.NetWeight,
		ExpirationRate: *p.ExpirationRate,
		Temperature:    *p.Temperature,
		FreezingRate:   *p.FreezingRate,
		ProductType:    p.ProductType,
	}
}

func (p *ProductTypeDoc) DocToModel() ProductType {
	return ProductType{
		ID:          p.ID,
		Description: p.Description,
	}
}
