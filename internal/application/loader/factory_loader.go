package loader

import "ProyectoFinal/pkg/models"

const (
	Seller    string = "sellers"
	Product   string = "products"
	Warehouse string = "warehouse"
	Section   string = "sections"
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
