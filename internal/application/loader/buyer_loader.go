package loader

import (
	"ProyectoFinal/pkg/models"
	"encoding/json"
	"os"
)

type BuyerLoader struct {
	path string
}

func (l *BuyerLoader) Load() (map[int]models.Buyer, error) {
	file, err := os.Open(l.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []models.BuyerDoc
	if err = json.NewDecoder(file).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to decode buyers from %s: %w", l.path, err)
	}

	newMap := make(map[int]models.Buyer)

	for _, buyerDoc := range data {
		buyer := buyerDoc.DocToModel()
		newMap[buyer.Id] = buyer
	}

	return newMap, nil
}
