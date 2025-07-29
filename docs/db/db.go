package db

import (
	"ProyectoFinal/internal/application/loader"
	"ProyectoFinal/pkg/models"
)

type Db struct {
	Seller    map[int]models.Seller
	Warehouse map[int]models.Warehouse
	Buyer     map[int]models.Buyer
	Section   map[int]models.Section
}

func LoadDB(loaderFilePath map[string]string) Db {
	factory := loader.NewLoaderFactory(loaderFilePath)

	//Load warehouse
	warehouseDB, err := factory.NewWarehouseLoader().Load()
	if err != nil {
		panic(err)
	}

	//Load buyer
	buyerDB, err := factory.NewBuyerLoader().Load()
	if err != nil {
		panic(err)
	}

	// Load sections
	sectionDb, err := factory.NewSectionLoader().Load()
	if err != nil {
		panic(err)
	}

	db := Db{
		Warehouse: warehouseDB,
		Buyer:     buyerDB,
		Section:   sectionDb,
	}
	return db
}
