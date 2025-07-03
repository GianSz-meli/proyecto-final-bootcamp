package models

type Buyer struct {
	Id           int    `json:"id"`
	CardNumberId string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}

type BuyerCreateDTO struct {
	CardNumberId string `json:"card_number_id" validate:"required"`
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
}

type BuyerUpdateDTO struct {
	CardNumberId *string `json:"card_number_id,omitempty" validate:"omitempty,min=1"`
	FirstName    *string `json:"first_name,omitempty" validate:"omitempty,min=1"`
	LastName     *string `json:"last_name,omitempty" validate:"omitempty,min=1"`
}

func (b BuyerCreateDTO) DocToModel() Buyer {
	return Buyer{
		CardNumberId: b.CardNumberId,
		FirstName:    b.FirstName,
		LastName:     b.LastName,
	}
}
