package models

type Buyer struct {
	Id           int
	CardNumberId int
	FirstName    string
	LastName     string
}

type BuyerCreateDto struct {
	CardNumberId int
	FirstName    string
	LastName     string
}

type BuyerUpdateDto struct {
	CardNumberId *int    `json:"card_number_id,omitempty"`
	FirstName    *string `json:"first_name,omitempty"`
	LastName     *string `json:"last_name,omitempty"`
}
