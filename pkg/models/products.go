package models

type Product struct {
	ID             int
	ProductCode    string
	Description    string
	Width          float64
	Height         float64
	Length         float64
	NetWeight      float64
	ExpirationRate float64
	Temperature    float32
	FreezingRate   float64
	ProductTypeID  int
	SellerID       *int
}

type ProductDoc struct {
	ID             int     `json:"id"`
	ProductCode    string  `json:"product_code" validate:"required"`
	Description    string  `json:"description" validate:"required,min=1"`
	Width          float64 `json:"width" validate:"required,gte=0"`
	Height         float64 `json:"height" validate:"required,gte=0"`
	Length         float64 `json:"length" validate:"required,gte=0"`
	NetWeight      float64 `json:"net_weight" validate:"required,gte=0"`
	ExpirationRate float64 `json:"expiration_rate" validate:"required,gte=0"`
	Temperature    float32 `json:"recommended_freezing_temperature" validate:"required"`
	FreezingRate   float64 `json:"freezing_rate" validate:"required,gte=0"`
	ProductTypeID  int     `json:"product_type_id"`
	SellerID       *int    `json:"seller_id"`
}

type ProductDocUpdate struct {
	ProductCode    *string  `json:"product_code" validate:"required"`
	Description    *string  `json:"description"`
	Width          *float64 `json:"width"`
	Height         *float64 `json:"height"`
	Length         *float64 `json:"length"`
	NetWeight      *float64 `json:"net_weight"`
	ExpirationRate *float64 `json:"expiration_rate"`
	Temperature    *float32 `json:"recommended_freezing_temperature"`
	FreezingRate   *float64 `json:"freezing_rate"`
	ProductTypeID  *int     `json:"product_type_id"`
}

func (p *Product) ModelToDoc() ProductDoc {
	return ProductDoc{
		ID:             p.ID,
		ProductCode:    p.ProductCode,
		Description:    p.Description,
		Width:          p.Width,
		Height:         p.Height,
		Length:         p.Length,
		NetWeight:      p.NetWeight,
		ExpirationRate: p.ExpirationRate,
		Temperature:    p.Temperature,
		FreezingRate:   p.FreezingRate,
		ProductTypeID:  p.ProductTypeID,
	}
}

func (p *ProductDoc) DocToModel() Product {
	return Product{
		ID:             p.ID,
		ProductCode:    p.ProductCode,
		Description:    p.Description,
		Width:          p.Width,
		Height:         p.Height,
		Length:         p.Length,
		NetWeight:      p.NetWeight,
		ExpirationRate: p.ExpirationRate,
		Temperature:    p.Temperature,
		FreezingRate:   p.FreezingRate,
		ProductTypeID:  p.ProductTypeID,
	}
}
