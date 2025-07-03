package loader

import (
	"ProyectoFinal/pkg/models"
	"encoding/json"
	"fmt"
	"os"
)

type ProductLoader struct {
	path string
}

func (s *ProductLoader) Load() (map[int]models.Product, error) {
	file, err := os.Open(s.path)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	var productsJSON []models.ProductDoc

	if err = json.NewDecoder(file).Decode(&productsJSON); err != nil {
		return nil, err
	}

	productsMap := map[int]models.Product{}

	for _, product := range productsJSON{
		productsMap[product.ID] = product.DocToModel()
	}

	return productsMap, nil
}