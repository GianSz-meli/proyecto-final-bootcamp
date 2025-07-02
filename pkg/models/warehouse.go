package models

// Warehouse represents a warehouse/fulfillment center
type Warehouse struct {
	ID                 int
	WarehouseCode      *string
	Address            *string
	Telephone          *string
	MinimumCapacity    *int
	MinimumTemperature *float64
}

func (w Warehouse) ModelToDoc() WarehouseDocument {
	return WarehouseDocument {
		ID:                 w.ID,
		WarehouseCode:      *w.WarehouseCode,
		Address:            *w.Address,
		Telephone:          *w.Telephone,
		MinimumCapacity:    *w.MinimumCapacity,
		MinimumTemperature: *w.MinimumTemperature,
	}
}

type WarehouseDocument struct {
	ID                 int    `json:"id"`
	WarehouseCode      string `json:"warehouse_code"`
	Address            string `json:"address"`
	Telephone          string `json:"telephone"`
	MinimumCapacity    int    `json:"minimum_capacity"`
	MinimumTemperature float64 `json:"minimum_temperature"`
}

func (w WarehouseDocument) DocToModel() Warehouse {
	return Warehouse{
		ID:                 w.ID,
		WarehouseCode:      &w.WarehouseCode,
		Address:            &w.Address,
		Telephone:          &w.Telephone,
		MinimumCapacity:    &w.MinimumCapacity,
		MinimumTemperature: &w.MinimumTemperature,
	}
}

type CreateWarehouseRequest struct {
	WarehouseCode      string  `json:"warehouse_code" validate:"required"`
	Address            string  `json:"address" validate:"required"`
	Telephone          string  `json:"telephone" validate:"required"`
	MinimumCapacity    int     `json:"minimum_capacity" validate:"required,min=1"`
	MinimumTemperature float64 `json:"minimum_temperature" validate:"required"`
}

func (c CreateWarehouseRequest) DocToModel() Warehouse {
	return Warehouse{
		WarehouseCode:      &c.WarehouseCode,
		Address:            &c.Address,
		Telephone:          &c.Telephone,
		MinimumCapacity:    &c.MinimumCapacity,
		MinimumTemperature: &c.MinimumTemperature,
	}
}