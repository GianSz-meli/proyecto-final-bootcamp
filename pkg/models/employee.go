package models

type Employee struct {
	ID           int
	CardNumberID string
	FirstName    string
	LastName     string
	WarehouseID  int
}

type EmployeeRequest struct {
	CardNumberID string `json:"card_number_id" validate:"required"`
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	WarehouseID  int    `json:"warehouse_id" validate:"required"`
}

type EmployeeUpdateRequest struct {
	CardNumberID *string `json:"card_number_id,omitempty"`
	FirstName    *string `json:"first_name,omitempty"`
	LastName     *string `json:"last_name,omitempty"`
	WarehouseID  *int    `json:"warehouse_id,omitempty"`
}

type EmployeeDoc struct {
	ID           int    `json:"id"`
	CardNumberID string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseID  int    `json:"warehouse_id"`
}

func (e Employee) ModelToDoc() EmployeeDoc {
	return EmployeeDoc(e)
}

func (e EmployeeDoc) DocToModel() Employee {
	return Employee(e)
}
