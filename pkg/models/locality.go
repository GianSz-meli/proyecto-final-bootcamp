package models

type Country struct {
	Id          int
	CountryName string
}

type Province struct {
	Id           int
	ProvinceName string
	Country      Country
}

type Locality struct {
	Id           int
	LocalityName string
	Province     Province
}

type LocalityCreateRequest struct {
	Id           int    `json:"id" validate:"omitempty,gt=0"`
	LocalityName string `json:"locality_name" validate:"required,min=1"`
	ProvinceName string `json:"province_name" validate:"required,min=1"`
	CountryName  string `json:"country_name" validate:"required,min=1"`
}

func (m *LocalityCreateRequest) DocToModel() Locality {
	return Locality{
		Id:           m.Id,
		LocalityName: m.LocalityName,
		Province: Province{
			ProvinceName: m.ProvinceName,
			Country: Country{
				CountryName: m.CountryName,
			},
		},
	}
}

func (l *Locality) ModelToDoc() LocalityCreateRequest {
	return LocalityCreateRequest{
		Id:           l.Id,
		LocalityName: l.LocalityName,
		ProvinceName: l.Province.ProvinceName,
		CountryName:  l.Province.Country.CountryName,
	}
}

type SellersByLocalityReport struct {
	LocalityId   int    `json:"locality_id"`
	LocalityName string `json:"locality_name"`
	SellersCount int    `json:"sellers_count"`
}
