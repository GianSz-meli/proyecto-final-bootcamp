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
