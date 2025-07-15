package models

type Carrier struct {
	Id          int       `json:"id"`
	Cid         string    `json:"cid"`
	CompanyName string    `json:"company_name"`
	Address     string    `json:"address"`
	Telephone   string    `json:"telephone"`
	LocalityId  int       `json:"locality_id"`
	Locality    *Locality `json:"locality"`
}

type CarrierCreateDTO struct {
	Cid         string `json:"cid" validate:"required"`
	CompanyName string `json:"company_name" validate:"required"`
	Address     string `json:"address" validate:"required"`
	Telephone   string `json:"telephone" validate:"required"`
	LocalityId  int    `json:"locality_id" validate:"required"`
}

type CarrierDoc struct {
	Id          int       `json:"id"`
	Cid         string    `json:"cid"`
	CompanyName string    `json:"company_name"`
	Address     string    `json:"address"`
	Telephone   string    `json:"telephone"`
	LocalityId  int       `json:"locality_id"`
	Locality    *Locality `json:"locality"`
}

func (c CarrierCreateDTO) CreateDtoToModel() *Carrier {
	return &Carrier{
		Cid:         c.Cid,
		CompanyName: c.CompanyName,
		Address:     c.Address,
		Telephone:   c.Telephone,
		LocalityId:  c.LocalityId,
		Locality:    &Locality{Id: c.LocalityId},
	}
}

func (c Carrier) ModelToDoc() CarrierDoc {
	return CarrierDoc{
		Id:          c.Id,
		Cid:         c.Cid,
		CompanyName: c.CompanyName,
		Address:     c.Address,
		Telephone:   c.Telephone,
		LocalityId:  c.LocalityId,
		Locality:    c.Locality,
	}
}

func (c CarrierDoc) DocToModel() Carrier {
	return Carrier{
		Id:          c.Id,
		Cid:         c.Cid,
		CompanyName: c.CompanyName,
		Address:     c.Address,
		Telephone:   c.Telephone,
		LocalityId:  c.LocalityId,
		Locality:    c.Locality,
	}
}
