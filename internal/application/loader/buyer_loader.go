package loader

import (
	"ProyectoFinal/pkg/models"
	"encoding/json"
	"fmt"
	"os"
)

type BuyerLoader struct {
	path string
}

func (l *BuyerLoader) Load() (map[int]models.Buyer, error) {
	file, err := os.Open(l.path)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	var data []models.Buyer

	if err = json.NewDecoder(file).Decode(&data); err != nil {
		return nil, err
	}

	newMap := make(map[int]models.Buyer)

	for _, buyer := range data {
		newMap[buyer.Id] = buyer
	}

	return newMap, nil
}
