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
	WarehouseID  int    `json:"warehouse_id" validate:"required,gt=0"`
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
	CardNumberID *string `json:"card_number_id,omitempty"`
	FirstName    *string `json:"first_name,omitempty"`
	LastName     *string `json:"last_name,omitempty"`
	WarehouseID  *int    `json:"warehouse_id,omitempty"`
}

func (r EmployeeUpdateRequest) UpdateFields(current Employee) Employee {
	updated := current
	if r.CardNumberID != nil {
		updated.CardNumberID = *r.CardNumberID
	}
	if r.FirstName != nil {
		updated.FirstName = *r.FirstName
	}
	if r.LastName != nil {
		updated.LastName = *r.LastName
	}
	if r.WarehouseID != nil {
		updated.WarehouseID = *r.WarehouseID
	}
	return updated
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
