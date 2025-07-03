package loader

import (
	"ProyectoFinal/pkg/models"
)

const (
	Seller    string = "sellers"
	Product   string = "products"
	Warehouse string = "warehouse"
	Section   string = "sections"
	Employee  string = "employee"
	Buyer     string = "buyers"
)

type FactoryLoader struct {
	paths map[string]string
}

func NewLoaderFactory(paths map[string]string) *FactoryLoader {
	return &FactoryLoader{paths: paths}
}

func (f *FactoryLoader) NewSellerLoader() Loader[models.Seller] {
	return &SellerLoader{path: f.paths[Seller]}
}

func (f *FactoryLoader) NewSectionLoader() Loader[models.Section] {
	return &SectionJSONFile{path: f.paths[Section]}
}

func (f *FactoryLoader) NewEmployeeLoader() Loader[models.Employee] {
	return &EmployeeLoader{path: f.paths[Employee]}
}

func (f *FactoryLoader) NewWarehouseLoader() Loader[models.Warehouse] {
	return &WarehouseLoader{path: f.paths[Warehouse]}
}

func (f *FactoryLoader) NewBuyerLoader() Loader[models.Buyer] {
	return &BuyerLoader{path: f.paths[Buyer]}
}
