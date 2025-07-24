package db

import (
	"ProyectoFinal/internal/application/loader"
	"ProyectoFinal/pkg/models"
)

type Db struct {
	Seller    map[int]models.Seller
	Warehouse map[int]models.Warehouse
	Product   map[int]models.Product
	Buyer     map[int]models.Buyer
	Section   map[int]models.Section
}

func LoadDB(loaderFilePath map[string]string) Db {
	factory := loader.NewLoaderFactory(loaderFilePath)

	// Load sellers
	sellerDB, err := factory.NewSellerLoader().Load()
	if err != nil {
		panic(err)
	}

	//Load warehouse
	warehouseDB, err := factory.NewWarehouseLoader().Load()
	if err != nil {
		panic(err)
	}

	//Load Product
	productDB, err := factory.NewProductLoader().Load()
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
		Seller:    sellerDB,
		Warehouse: warehouseDB,
		Product:   productDB,
		Buyer:     buyerDB,
		Section:   sectionDb,
	}
	return db
}
