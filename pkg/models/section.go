package models

type Section struct {
	ID int `json:"id"`
	SectionAttributes
}

type SectionAttributes struct {
	SectionNumber      string  `json:"section_number" validate:"required"`
	CurrentTemperature float64 `json:"current_temperature" validate:"gte=-50,lte=50,gtefield=MinimumTemperature"`
	MinimumTemperature float64 `json:"minimum_temperature" validate:"gte=-50,lte=50"`
	CurrentCapacity    int     `json:"current_capacity" validate:"gte=0,gtefield=MinimumCapacity,ltefield=MaximumCapacity"`
	MinimumCapacity    int     `json:"minimum_capacity" validate:"gte=0,ltefield=MaximumCapacity"`
	MaximumCapacity    int     `json:"maximum_capacity" validate:"required,gt=0"`
	WarehouseID        int     `json:"warehouse_id" validate:"required,gt=0"`
	ProductTypeID      int     `json:"product_type_id" validate:"required,gt=0"`
	ProductBatches     []int   `json:"product_batches,omitempty"`
}

func (s Section) ModelToDoc() SectionDoc {
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
	SectionNumber      string  `json:"section_number"`
	CurrentTemperature float64 `json:"current_temperature"`
	MinimumTemperature float64 `json:"minimum_temperature"`
	CurrentCapacity    int     `json:"current_capacity"`
	MinimumCapacity    int     `json:"minimum_capacity"`
	MaximumCapacity    int     `json:"maximum_capacity"`
	WarehouseID        int     `json:"warehouse_id"`
	ProductTypeID      int     `json:"product_type_id"`
}

func (s SectionDoc) DocToModel() Section {
	return Section{
		ID: s.ID,
		SectionAttributes: SectionAttributes{
			SectionNumber:      s.SectionNumber,
			CurrentTemperature: s.CurrentTemperature,
			MinimumTemperature: s.MinimumTemperature,
			CurrentCapacity:    s.CurrentCapacity,
			MinimumCapacity:    s.MinimumCapacity,
			MaximumCapacity:    s.MaximumCapacity,
			WarehouseID:        s.WarehouseID,
			ProductTypeID:      s.ProductTypeID,
		},
	}
}

type CreateSectionRequest struct {
	SectionNumber      string  `json:"section_number" validate:"required"`
	CurrentTemperature float64 `json:"current_temperature" validate:"gte=-50,lte=50,gtefield=MinimumTemperature"`
	MinimumTemperature float64 `json:"minimum_temperature" validate:"gte=-50,lte=50"`
	CurrentCapacity    int     `json:"current_capacity" validate:"gte=0,gtefield=MinimumCapacity,ltefield=MaximumCapacity"`
	MinimumCapacity    int     `json:"minimum_capacity" validate:"gte=0,ltefield=MaximumCapacity"`
	MaximumCapacity    int     `json:"maximum_capacity" validate:"required,gt=0"`
	WarehouseID        int     `json:"warehouse_id" validate:"required,gt=0"`
	ProductTypeID      int     `json:"product_type_id" validate:"required,gt=0"`
}

func (c CreateSectionRequest) DocToModel() Section {
	return Section{
		SectionAttributes: SectionAttributes{
			SectionNumber:      c.SectionNumber,
			CurrentTemperature: c.CurrentTemperature,
			MinimumTemperature: c.MinimumTemperature,
			CurrentCapacity:    c.CurrentCapacity,
			MinimumCapacity:    c.MinimumCapacity,
			MaximumCapacity:    c.MaximumCapacity,
			WarehouseID:        c.WarehouseID,
			ProductTypeID:      c.ProductTypeID,
		},
	}
}

type UpdateSectionRequest struct {
	SectionNumber      *string  `json:"section_number" validate:"omitempty"`
	CurrentTemperature *float64 `json:"current_temperature" validate:"omitempty,gte=-50,lte=50"`
	MinimumTemperature *float64 `json:"minimum_temperature" validate:"omitempty,gte=-50,lte=50"`
	CurrentCapacity    *int     `json:"current_capacity" validate:"omitempty,gte=0"`
	MinimumCapacity    *int     `json:"minimum_capacity" validate:"omitempty,gte=0"`
	MaximumCapacity    *int     `json:"maximum_capacity" validate:"omitempty,gt=0"`
	WarehouseID        *int     `json:"warehouse_id" validate:"omitempty,gt=0"`
	ProductTypeID      *int     `json:"product_type_id" validate:"omitempty,gt=0"`
}
