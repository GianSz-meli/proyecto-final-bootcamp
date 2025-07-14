package models

// Country model according to database schema
type Country struct {
	Id          *int
	CountryName *string
}

// Province model according to database schema
type Province struct {
	Id           *int
	ProvinceName *string
	Country      *Country
}

// Locality model according to database schema
type Locality struct {
	Id           *int
	LocalityName *string
	Province     *Province
}
