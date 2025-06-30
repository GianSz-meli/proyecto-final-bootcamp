package models

type Seller struct {
	Id          int    `json:"id" validate:"required"`
	Cid         int    `json:"cid" validate:"required"`
	CompanyName string `json:"company_name" validate:"required"`
	Address     string `json:"address" validate:"required"`
	Telephone   string `json:"telephone" validate:"required"`
}
