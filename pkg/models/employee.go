package models

type Employee struct {
	ID           int
	CardNumberID string
	FirstName    string
	LastName     string
	WarehouseID  *int
}

type EmployeeRequest struct {
	CardNumberID string `json:"card_number_id" validate:"required"`
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	WarehouseID  *int   `json:"warehouse_id,omitempty" validate:"omitempty,gt=0"`
}

func (r EmployeeRequest) DocToModel() Employee {
	return Employee{
		CardNumberID: r.CardNumberID,
		FirstName:    r.FirstName,
		LastName:     r.LastName,
		WarehouseID:  r.WarehouseID,
	}
}

type EmployeeUpdateRequest struct {
	CardNumberID *string `json:"card_number_id,omitempty" validate:"omitempty,min=1"`
	FirstName    *string `json:"first_name,omitempty" validate:"omitempty,min=1"`
	LastName     *string `json:"last_name,omitempty" validate:"omitempty,min=1"`
	WarehouseID  *int    `json:"warehouse_id,omitempty" validate:"omitempty,gt=0"`
}

type EmployeeDoc struct {
	ID           int    `json:"id"`
	CardNumberID string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseID  *int   `json:"warehouse_id,omitempty"`
}

func (e Employee) ModelToDoc() EmployeeDoc {
	return EmployeeDoc{
		ID:           e.ID,
		CardNumberID: e.CardNumberID,
		FirstName:    e.FirstName,
		LastName:     e.LastName,
		WarehouseID:  e.WarehouseID,
	}
}

func (e EmployeeDoc) DocToModel() Employee {
	return Employee{
		ID:           e.ID,
		CardNumberID: e.CardNumberID,
		FirstName:    e.FirstName,
		LastName:     e.LastName,
		WarehouseID:  e.WarehouseID,
	}
}
