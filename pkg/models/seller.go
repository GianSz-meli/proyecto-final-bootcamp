package models

type Seller struct {
	Id          int    `json:"id"`
	Cid         int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
}

func (m *Seller) ModelToDoc() SellerDoc {
	return SellerDoc{
		Id:          m.Id,
		Cid:         m.Cid,
		CompanyName: m.CompanyName,
		Address:     m.Address,
		Telephone:   m.Telephone,
	}
}

type SellerDoc struct {
	Id          int    `json:"id"`
	Cid         int    `json:"cid" validate:"required"`
	CompanyName string `json:"company_name" validate:"required"`
	Address     string `json:"address" validate:"required"`
	Telephone   string `json:"telephone" validate:"required"`
}

func (s *SellerDoc) DocToModel() Seller {
	return Seller{
		Id:          s.Id,
		Cid:         s.Cid,
		CompanyName: s.CompanyName,
		Address:     s.Address,
		Telephone:   s.Telephone,
	}
}
