package models

type Buyer struct {
	Id           int
	CardNumberId string
	FirstName    string
	LastName     string
}

type BuyerCreateDTO struct {
	CardNumberId string
	FirstName    string
	LastName     string
}

type BuyerUpdateDTO struct {
	CardNumberId *string `json:"card_number_id,omitempty" validate:"omitempty,min=1"`
	FirstName    *string `json:"first_name,omitempty" validate:"omitempty,min=1"`
	LastName     *string `json:"last_name,omitempty" validate:"omitempty,min=1"`
}

func DTOToBuyer(dto BuyerCreateDTO) Buyer {
	return Buyer{
		CardNumberId: dto.CardNumberId,
		FirstName:    dto.FirstName,
		LastName:     dto.LastName,
	}
}
