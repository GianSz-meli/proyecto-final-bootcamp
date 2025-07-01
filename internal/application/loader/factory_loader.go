package loader

import (
	"ProyectoFinal/pkg/models"
	employeemodel "ProyectoFinal/pkg/models/employee"
)

const (
	Seller    string = "sellers"
	Product   string = "products"
	Warehouse string = "warehouse"
	Employee  string = "employee"
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

func (f *FactoryLoader) NewEmployeeLoader() Loader[employeemodel.Employee] {
	return &EmployeeLoader{path: f.paths[Employee]}
}
func (f *FactoryLoader) NewWarehouseLoader() Loader[models.Warehouse] {
	return &WarehouseLoader{path: f.paths[Warehouse]}
}
