package models

import "time"

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

type SectionProductReport struct {
	SectionID     int    `json:"section_id"`
	SectionNumber string `json:"section_number"`
	ProductsCount int    `json:"products_count"`
}
