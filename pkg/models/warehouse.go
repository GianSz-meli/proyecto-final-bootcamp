package models

type Warehouse struct {
	ID                 int
	WarehouseCode      string
	Address            string
	Telephone          string
	MinimumCapacity    int
	MinimumTemperature float64
	LocalityId         *int
}

func (w Warehouse) ModelToDoc() WarehouseDocument {
	return WarehouseDocument{
		ID:                 w.ID,
		WarehouseCode:      w.WarehouseCode,
		Address:            w.Address,
		Telephone:          w.Telephone,
		MinimumCapacity:    w.MinimumCapacity,
		MinimumTemperature: w.MinimumTemperature,
		LocalityId:         w.LocalityId,
	}
}

type WarehouseDocument struct {
	ID                 int     `json:"id"`
	WarehouseCode      string  `json:"warehouse_code"`
	Address            string  `json:"address"`
	Telephone          string  `json:"telephone"`
	MinimumCapacity    int     `json:"minimum_capacity"`
	MinimumTemperature float64 `json:"minimum_temperature"`
	LocalityId         *int     `json:"locality_id"`
}

func (w WarehouseDocument) DocToModel() Warehouse {
	return Warehouse{
		ID:                 w.ID,
		WarehouseCode:      w.WarehouseCode,
		Address:            w.Address,
		Telephone:          w.Telephone,
		MinimumCapacity:    w.MinimumCapacity,
		MinimumTemperature: w.MinimumTemperature,
		LocalityId:         w.LocalityId,
	}
}

type CreateWarehouseRequest struct {
	WarehouseCode      string   `json:"warehouse_code" validate:"required,min=1"`
	Address            string   `json:"address" validate:"required,min=1"`
	Telephone          string   `json:"telephone" validate:"required,numeric,min=7"`
	MinimumCapacity    *int     `json:"minimum_capacity" validate:"required,min=0"`
	MinimumTemperature *float64 `json:"minimum_temperature" validate:"required"`
	LocalityId         *int     `json:"locality_id"`
}

func (c CreateWarehouseRequest) DocToModel() Warehouse {
	return Warehouse{
		WarehouseCode:      c.WarehouseCode,
		Address:            c.Address,
		Telephone:          c.Telephone,
		MinimumCapacity:    *c.MinimumCapacity,
		MinimumTemperature: *c.MinimumTemperature,
		LocalityId:         c.LocalityId,
	}
}

type UpdateWarehouseRequest struct {
	WarehouseCode      *string  `json:"warehouse_code" validate:"omitempty,min=1"`
	Address            *string  `json:"address" validate:"omitempty,min=1"`
	Telephone          *string  `json:"telephone" validate:"omitempty,numeric,min=7"`
	MinimumCapacity    *int     `json:"minimum_capacity" validate:"omitempty,min=0"`
	MinimumTemperature *float64 `json:"minimum_temperature"`
	LocalityId         *int     `json:"locality_id"`
}

func (u UpdateWarehouseRequest) DocToModel() Warehouse {
	return Warehouse{
		WarehouseCode:      *u.WarehouseCode,
		Address:            *u.Address,
		Telephone:          *u.Telephone,
		MinimumCapacity:    *u.MinimumCapacity,
		MinimumTemperature: *u.MinimumTemperature,
		LocalityId:         u.LocalityId,
	}
}
