package models

// Warehouse represents a warehouse/fulfillment center
type Warehouse struct {
	ID                 int
	WarehouseCode      string
	Address            string
	Telephone          string
	MinimunCapacity    int
	MinimunTemperature float64
}

type WarehouseDocument struct {
	ID                 int    `json:"id"`
	WarehouseCode      string `json:"warehouse_code"`
	Address            string `json:"address"`
	Telephone          string `json:"telephone"`
	MinimunCapacity    int    `json:"minimun_capacity"`
	MinimunTemperature float64 `json:"minimun_temperature"`
}
