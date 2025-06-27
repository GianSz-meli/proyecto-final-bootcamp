package loader

import (
	"ProyectoFinal/pkg/models"
	"encoding/json"
	"os"
)

func NewSectionJSONFile(path string) *SectionJSONFile {
	return &SectionJSONFile{
		path: path,
	}
}

type SectionJSONFile struct {
	path string
}

func (l *SectionJSONFile) Load() (s map[int]models.Section, err error) {
	file, err := os.Open(l.path)
	if err != nil {
		return
	}
	defer file.Close()

	var sectionsJSON []models.Section
	err = json.NewDecoder(file).Decode(&sectionsJSON)
	if err != nil {
		return
	}

	s = make(map[int]models.Section)
	for _, sec := range sectionsJSON {
		s[sec.ID] = models.Section{
			ID: sec.ID,
			SectionAttributes: models.SectionAttributes{
				SectionNumber:      sec.SectionNumber,
				CurrentTemperature: sec.CurrentTemperature,
				MinimumTemperature: sec.MinimumTemperature,
				CurrentCapacity:    sec.CurrentCapacity,
				MinimumCapacity:    sec.MinimumCapacity,
				MaximumCapacity:    sec.MaximumCapacity,
				WarehouseID:        sec.WarehouseID,
				ProductTypeID:      sec.ProductTypeID,
				ProductBatches:     sec.ProductBatches,
			},
		}
	}

	return
}
