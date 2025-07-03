package loader

import (
	"ProyectoFinal/pkg/models"
	"encoding/json"
	"fmt"
	"os"
)

type SellerLoader struct {
	path string
}

func (s *SellerLoader) Load() (map[int]models.Seller, error) {
	file, err := os.Open(s.path)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	var sellersJSON []models.SellerDoc

	if err = json.NewDecoder(file).Decode(&sellersJSON); err != nil {
		return nil, err
	}

	sellerMap := map[int]models.Seller{}

	for _, seller := range sellersJSON {
		sellerMap[seller.Id] = seller.DocToModel()
	}

	return sellerMap, nil
}
