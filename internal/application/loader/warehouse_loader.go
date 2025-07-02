package loader

import (
	"ProyectoFinal/pkg/models"
	"encoding/json"
	"os"
)

type WarehouseLoader struct {
	path string
}

func (l *WarehouseLoader) Load() (map[int]models.Warehouse, error) {
	file, err := os.Open(l.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var warehouseDocuments []models.WarehouseDocument
	if err = json.NewDecoder(file).Decode(&warehouseDocuments); err != nil {
		return nil, err
	}

	warehouseMap := map[int]models.Warehouse{}

	for _, warehouseDocument := range warehouseDocuments {
		warehouseMap[warehouseDocument.ID] = warehouseDocument.DocToModel()
	}

	return warehouseMap, nil
}
