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

	var sectionsDoc []models.SectionDoc
	err = json.NewDecoder(file).Decode(&sectionsDoc)
	if err != nil {
		return
	}

	s = make(map[int]models.Section)
	for _, secDoc := range sectionsDoc {
		section := models.Section{
			ID: secDoc.ID,
			SectionAttributes: models.SectionAttributes{
				SectionNumber:      secDoc.SectionNumber,
				CurrentTemperature: secDoc.CurrentTemperature,
				MinimumTemperature: secDoc.MinimumTemperature,
				CurrentCapacity:    secDoc.CurrentCapacity,
				MinimumCapacity:    secDoc.MinimumCapacity,
				MaximumCapacity:    secDoc.MaximumCapacity,
				WarehouseID:        secDoc.WarehouseID,
				ProductTypeID:      secDoc.ProductTypeID,
				ProductBatches:     nil,
			},
		}
		s[section.ID] = section
	}

	return
}
