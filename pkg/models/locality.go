package models

type Locality struct {
	Id           int
	LocalityName string
	ProvinceName string
	CountryName  string
}

func (m *Locality) ModelToDoc() LocalityDoc {
	return LocalityDoc{
		Id:           m.Id,
		LocalityName: m.LocalityName,
		ProvinceName: m.ProvinceName,
		CountryName:  m.CountryName,
	}
}

type LocalityDoc struct {
	Id           int    `json:"id"`
	LocalityName string `json:"locality_name" validate:"required,min=1"`
	ProvinceName string `json:"province_name" validate:"required,min=1"`
	CountryName  string `json:"country_name" validate:"required,min=1"`
}

func (l *LocalityDoc) DocToModel() Locality {
	return Locality{
		Id:           l.Id,
		LocalityName: l.LocalityName,
		ProvinceName: l.ProvinceName,
		CountryName:  l.CountryName,
	}
}

type Country struct {
	id          int
	countryName string
}
type Province struct {
	id           int
	provinceName string
	idCountry    int
}
