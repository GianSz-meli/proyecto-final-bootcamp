package models

import "time"

type Section struct {
	ID int `json:"id"`
	SectionAttributes
}

type SectionAttributes struct {
	SectionNumber      string         `json:"section_number" validate:"required"`
	CurrentTemperature float64        `json:"current_temperature" validate:"gte=-50,lte=50,gtefield=MinimumTemperature"`
	MinimumTemperature float64        `json:"minimum_temperature" validate:"gte=-50,lte=50"`
	CurrentCapacity    int            `json:"current_capacity" validate:"gte=0,gtefield=MinimumCapacity,ltefield=MaximumCapacity"`
	MinimumCapacity    int            `json:"minimum_capacity" validate:"gte=0,ltefield=MaximumCapacity"`
	MaximumCapacity    int            `json:"maximum_capacity" validate:"required,gt=0"`
	WarehouseID        int            `json:"warehouse_id" validate:"required,gt=0"`
	ProductTypeID      int            `json:"product_type_id" validate:"required,gt=0"`
	ProductBatches     []ProductBatch `json:"product_batches,omitempty"`
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

type ProductBatch struct {
	ID                 int       `json:"id"`
	BatchNumber        string    `json:"batch_number" validate:"required"`
	CurrentQuantity    int       `json:"current_quantity" validate:"required,min=0"`
	CurrentTemperature float64   `json:"current_temperature" validate:"required"`
	DueDate            time.Time `json:"due_date" validate:"required"`
	InitialQuantity    int       `json:"initial_quantity" validate:"required,min=0"`
	ManufacturingDate  time.Time `json:"manufacturing_date" validate:"required"`
	ManufacturingHour  int       `json:"manufacturing_hour" validate:"required,min=0,max=23"`
	MinimumTemperature float64   `json:"minimum_temperature" validate:"required"`
	ProductID          int       `json:"product_id" validate:"required"`
	SectionID          int       `json:"section_id" validate:"required"`
}

type ProductBatchCreateRequest struct {
	BatchNumber        string    `json:"batch_number" validate:"required"`
	CurrentQuantity    int       `json:"current_quantity" validate:"required,min=0"`
	CurrentTemperature float64   `json:"current_temperature" validate:"required"`
	DueDate            time.Time `json:"due_date" validate:"required"`
	InitialQuantity    int       `json:"initial_quantity" validate:"required,min=0"`
	ManufacturingDate  time.Time `json:"manufacturing_date" validate:"required"`
	ManufacturingHour  int       `json:"manufacturing_hour" validate:"required,min=0,max=23"`
	MinimumTemperature float64   `json:"minimum_temperature" validate:"required"`
	ProductID          int       `json:"product_id" validate:"required"`
	SectionID          int       `json:"section_id" validate:"required"`
}

type ProductBatchDoc struct {
	ID                 int       `json:"id"`
	BatchNumber        string    `json:"batch_number"`
	CurrentQuantity    int       `json:"current_quantity"`
	CurrentTemperature float64   `json:"current_temperature"`
	DueDate            time.Time `json:"due_date"`
	InitialQuantity    int       `json:"initial_quantity"`
	ManufacturingDate  time.Time `json:"manufacturing_date"`
	ManufacturingHour  int       `json:"manufacturing_hour"`
	MinimumTemperature float64   `json:"minimum_temperature"`
	ProductID          int       `json:"product_id"`
	SectionID          int       `json:"section_id"`
}

type SectionProductReport struct {
	SectionID     int    `json:"section_id"`
	SectionNumber string `json:"section_number"`
	ProductsCount int    `json:"products_count"`
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

func (pb ProductBatch) ModelToDoc() ProductBatchDoc {
	return ProductBatchDoc{
		ID:                 pb.ID,
		BatchNumber:        pb.BatchNumber,
		CurrentQuantity:    pb.CurrentQuantity,
		CurrentTemperature: pb.CurrentTemperature,
		DueDate:            pb.DueDate,
		InitialQuantity:    pb.InitialQuantity,
		ManufacturingDate:  pb.ManufacturingDate,
		ManufacturingHour:  pb.ManufacturingHour,
		MinimumTemperature: pb.MinimumTemperature,
		ProductID:          pb.ProductID,
		SectionID:          pb.SectionID,
	}
}

func (pb ProductBatchDoc) DocToModel() ProductBatch {
	return ProductBatch{
		ID:                 pb.ID,
		BatchNumber:        pb.BatchNumber,
		CurrentQuantity:    pb.CurrentQuantity,
		CurrentTemperature: pb.CurrentTemperature,
		DueDate:            pb.DueDate,
		InitialQuantity:    pb.InitialQuantity,
		ManufacturingDate:  pb.ManufacturingDate,
		ManufacturingHour:  pb.ManufacturingHour,
		MinimumTemperature: pb.MinimumTemperature,
		ProductID:          pb.ProductID,
		SectionID:          pb.SectionID,
	}
}

func (pbr ProductBatchCreateRequest) CreateRequestToModel() ProductBatch {
	return ProductBatch{
		BatchNumber:        pbr.BatchNumber,
		CurrentQuantity:    pbr.CurrentQuantity,
		CurrentTemperature: pbr.CurrentTemperature,
		DueDate:            pbr.DueDate,
		InitialQuantity:    pbr.InitialQuantity,
		ManufacturingDate:  pbr.ManufacturingDate,
		ManufacturingHour:  pbr.ManufacturingHour,
		MinimumTemperature: pbr.MinimumTemperature,
		ProductID:          pbr.ProductID,
		SectionID:          pbr.SectionID,
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
