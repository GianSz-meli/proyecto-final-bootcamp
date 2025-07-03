package models

type Section struct {
	ID int `json:"id"`
	SectionAttributes
}

type SectionAttributes struct {
	SectionNumber      int            `json:"section_number" validate:"required,gt=0"`
	CurrentTemperature float64        `json:"current_temperature" validate:"gte=-50,lte=50,gtefield=MinimumTemperature"`
	MinimumTemperature float64        `json:"minimum_temperature" validate:"gte=-50,lte=50"`
	CurrentCapacity    int            `json:"current_capacity" validate:"gte=0,gtefield=MinimumCapacity,ltefield=MaximumCapacity"`
	MinimumCapacity    int            `json:"minimum_capacity" validate:"gte=0,ltefield=MaximumCapacity"`
	MaximumCapacity    int            `json:"maximum_capacity" validate:"required,gt=0"`
	WarehouseID        int            `json:"warehouse_id" validate:"required,gt=0"`
	ProductTypeID      int            `json:"product_type_id" validate:"required,gt=0"`
	ProductBatches     []ProductBatch `json:"product_batches,omitempty"`
}

func (s Section) ToModelDoc() SectionDoc {
	return SectionDoc{
		ID:                 s.ID,
		SectionNumber:      s.SectionNumber,
		CurrentTemperature: s.CurrentTemperature,
		MinimumTemperature: s.MinimumTemperature,
		CurrentCapacity:    s.CurrentCapacity,
		MinimumCapacity:    s.MinimumCapacity,
		MaximumCapacity:    s.MaximumCapacity,
		WarehouseID:        s.WarehouseID,
		ProductTypeID:      s.ProductTypeID,
	}
}

type SectionDoc struct {
	ID                 int     `json:"id"`
	SectionNumber      int     `json:"section_number"`
	CurrentTemperature float64 `json:"current_temperature"`
	MinimumTemperature float64 `json:"minimum_temperature"`
	CurrentCapacity    int     `json:"current_capacity"`
	MinimumCapacity    int     `json:"minimum_capacity"`
	MaximumCapacity    int     `json:"maximum_capacity"`
	WarehouseID        int     `json:"warehouse_id"`
	ProductTypeID      int     `json:"product_type_id"`
}

type ProductBatch struct {
	ID         int   `json:"id"`
	ProductsID []int `json:"products_id"`
}
