package models

type Seller struct {
	Id          int    `json:"id"`
	Cid         string `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
	LocalityId  int    `json:"locality_id"`
}

func (m *Seller) ModelToDoc() SellerDoc {
	return SellerDoc{
		Id:          m.Id,
		Cid:         m.Cid,
		CompanyName: m.CompanyName,
		Address:     m.Address,
		Telephone:   m.Telephone,
		LocalityId:  m.LocalityId,
	}
}

type SellerDoc struct {
	Id          int    `json:"id"`
	Cid         string `json:"cid" validate:"required"`
	CompanyName string `json:"company_name" validate:"required"`
	Address     string `json:"address" validate:"required"`
	Telephone   string `json:"telephone" validate:"required"`
	LocalityId  int    `json:"locality_id" validate:"required"`
}

func (s *SellerDoc) DocToModel() Seller {
	return Seller{
		Id:          s.Id,
		Cid:         s.Cid,
		CompanyName: s.CompanyName,
		Address:     s.Address,
		Telephone:   s.Telephone,
		LocalityId:  s.LocalityId,
	}
}

type CreateSellerRequest struct {
	Cid         string `json:"cid" validate:"required"`
	CompanyName string `json:"company_name" validate:"required"`
	Address     string `json:"address" validate:"required"`
	Telephone   string `json:"telephone" validate:"required"`
	LocalityId  int    `json:"locality_id" validate:"required,gt=0"`
}

func (s *CreateSellerRequest) DocToModel() Seller {
	return Seller{
		Cid:         s.Cid,
		CompanyName: s.CompanyName,
		Address:     s.Address,
		Telephone:   s.Telephone,
		LocalityId:  s.LocalityId,
	}
}

type UpdateSellerRequest struct {
	Cid         *string `json:"cid" validate:"omitempty,min=1"`
	CompanyName *string `json:"company_name" validate:"omitempty,min=1"`
	Address     *string `json:"address" validate:"omitempty,min=1"`
	Telephone   *string `json:"telephone" validate:"omitempty,min=1"`
	LocalityId  *int    `json:"locality_id" validate:"omitempty,gt=0"`
}
