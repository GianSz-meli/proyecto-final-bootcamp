package loader

import "ProyectoFinal/pkg/models"

const (
	Seller    string = "sellers"
	Product   string = "products"
	Warehouse string = "warehouse"
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
