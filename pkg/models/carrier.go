package models

// Domain Model
type Carrier struct {
	Id          int
	Cid         string
	CompanyName string
	Address     string
	Telephone   string
	LocalityId  int
}

type CarrierReport struct {
	LocalityId    int
	LocalityName  string
	CarriersCount int
}

// Request DTO
type CarrierCreateDTO struct {
	Cid         string `json:"cid" validate:"required"`
	CompanyName string `json:"company_name" validate:"required"`
	Address     string `json:"address" validate:"required"`
	Telephone   string `json:"telephone" validate:"required,numeric,min=7"`
	LocalityId  int    `json:"locality_id" validate:"required"`
}

// Response DTO
type CarrierCreateDoc struct {
	Id          int       `json:"id"`
	Cid         string    `json:"cid"`
	CompanyName string    `json:"company_name"`
	Address     string    `json:"address"`
	Telephone   string    `json:"telephone"`
	LocalityId  int       `json:"locality_id"`
}

// Request DTO to Domain Model
func (c CarrierCreateDTO) CreateDtoToModel() *Carrier {
	return &Carrier{
		Cid:         c.Cid,
		CompanyName: c.CompanyName,
		Address:     c.Address,
		Telephone:   c.Telephone,
		LocalityId:  c.LocalityId,
	}
}

func (c Carrier) ModelToDoc() CarrierCreateDoc {
	return CarrierCreateDoc(c)
}
