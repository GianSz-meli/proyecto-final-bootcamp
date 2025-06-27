package loader

import (
	"ProyectoFinal/pkg/models"
	"encoding/json"
	"fmt"
	"os"
)

type SellerJSONFile struct {
	path string
}

func (s *SellerJSONFile) Load() (map[int]models.Seller, error) {
	file, err := os.Open(s.path)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	var sellersJSON []models.Seller

	if err = json.NewDecoder(file).Decode(&sellersJSON); err != nil {
		return nil, err
	}

	sellerMap := map[int]models.Seller{}

	for _, seller := range sellersJSON {
		sellerMap[seller.Id] = seller
	}

	return sellerMap, nil
}
